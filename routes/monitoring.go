package routes

import (
	"github.com/kataras/iris"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"

	"github.com/djlechuck/recalbox-manager/utils/errors"
)

// GetMonitoringHandler handles GET requests on /monitoring.
func GetMonitoringHandler(ctx iris.Context) {
	vm, err := mem.VirtualMemory()
	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	c, err := cpu.Percent(0, true)
	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	d, err := disk.Partitions(false)
	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	var usage []*disk.UsageStat

	for _, part := range d {
		u, err := disk.Usage(part.Mountpoint)
		if err != nil {
			ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
			ctx.StatusCode(500)

			return
		}

		usage = append(usage, u)
	}

	ctx.ViewData("PageTitle", ctx.Translate("Monitoring"))
	ctx.ViewData("Cpu", c)
	ctx.ViewData("Memory", vm)
	ctx.ViewData("Disk", usage)

	ctx.View("views/monitoring.pug")
}

// fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", vm.Total, vm.Free, vm.UsedPercent)

// fmt.Println(u.Path + "\t" + strconv.FormatFloat(u.UsedPercent, 'f', 2, 64) + "% full.")
// fmt.Println("Total: " + strconv.FormatUint(u.Total/1024/1024/1024, 10) + " GiB")
// fmt.Println("Free:  " + strconv.FormatUint(u.Free/1024/1024/1024, 10) + " GiB")
// fmt.Println("Used:  " + strconv.FormatUint(u.Used/1024/1024/1024, 10) + " GiB")
