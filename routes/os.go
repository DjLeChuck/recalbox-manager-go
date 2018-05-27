package routes

import (
	"os"
	"os/exec"
	"syscall"

	"github.com/djlechuck/recalbox-manager/utils/errors"
	"github.com/kataras/iris"
)

// GetRebootOsHandler handles GET requests on /os/reboot.
func GetRebootOsHandler(ctx iris.Context) {
	binary, err := exec.LookPath("reboot")
	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}
	args := []string{"reboot"}
	env := os.Environ()

	err = syscall.Exec(binary, args, env)
	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	ctx.Redirect("/wait")
}

// GetShutdownOsHandler handles GET requests on /os/shutdown.
func GetShutdownOsHandler(ctx iris.Context) {
	binary, err := exec.LookPath("shutdown")
	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}
	args := []string{"shutdown", "-h", "now"}
	env := os.Environ()

	err = syscall.Exec(binary, args, env)
	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	ctx.Redirect("/wait")
}
