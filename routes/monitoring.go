package routes

import (
	"fmt"
	"math"
	"os/exec"
	"strconv"
	"strings"

	"github.com/kataras/iris"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"

	"github.com/djlechuck/recalbox-manager/structs"
	"github.com/djlechuck/recalbox-manager/utils/errors"
)

// GetMonitoringHandler handles GET requests on /monitoring.
func GetMonitoringHandler(ctx iris.Context) {
	// Memory usage
	vm, err := mem.VirtualMemory()
	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	// Convert to MB (X / 1024^2)
	vm.Available = vm.Available / 1048576
	vm.Total = vm.Total / 1048576

	// CPU percent usage
	cpuTmp, err := cpu.Percent(0, true)
	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	var cpu []*structs.CPU
	for i, c := range cpuTmp {
		cpu = append(cpu, &structs.CPU{
			Number: i,
			Value:  fmt.Sprintf("%.2f", c),
		})
	}

	// Mounted disks usages
	d, err := disk.Partitions(false)
	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	var usage []*structs.Disk

	for _, part := range d {
		u, err := disk.Usage(part.Mountpoint)
		if err != nil {
			ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
			ctx.StatusCode(500)

			return
		}

		s := disk.PartitionStat(part)
		disk := &structs.Disk{
			Device:      s.Device,
			Path:        u.Path,
			UsedPercent: strconv.FormatFloat(u.UsedPercent, 'f', 2, 64),
			Used:        strconv.FormatUint(u.Used/1024/1024/1024, 10),
			Free:        strconv.FormatUint(u.Free/1024/1024/1024, 10),
			Total:       strconv.FormatUint(u.Total/1024/1024/1024, 10),
		}

		usage = append(usage, disk)
	}

	if ctx.IsAjax() {
		ctx.JSON(iris.Map{
			"cpu":         cpu,
			"memory":      vm,
			"temperature": getTemperature(),
		})
	} else {
		ctx.ViewData("PageTitle", ctx.Translate("Monitoring"))
		ctx.ViewData("Cpu", cpu)
		ctx.ViewData("Memory", vm)
		ctx.ViewData("Disks", usage)
		ctx.ViewData("Temperature", getTemperature())

		ctx.View("views/monitoring.pug")
	}
}

func getTemperature() *structs.Temperature {
	out, err := exec.Command("cat", "/sys/class/thermal/thermal_zone0/temp").CombinedOutput()
	if err != nil {
		return nil
	}

	t, err := strconv.ParseFloat(strings.Trim(string(out), "\n"), 64)
	if err != nil {
		return nil
	}

	t = t / 1000
	percent := math.Floor(t)
	var color string

	if percent > 70 {
		color = "orange"
	} else if percent < 30 {
		color = "green"
	}

	return &structs.Temperature{
		Current:        math.Round(t),
		CurrentPercent: percent,
		Max:            100,
		Color:          color,
	}
}
