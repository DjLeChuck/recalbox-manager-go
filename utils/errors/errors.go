package errors

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/kataras/iris/context"
)

func todayFilename() string {
	today := time.Now().Format("2006-01-02")

	return today + ".txt"
}

func getRequestLogs(ctx context.Context) string {
	status := ctx.GetStatusCode()
	path := ctx.Path()
	method := ctx.Method()

	return fmt.Sprintf("%d %s %s", status, path, method)
}

// NewLogFile opens a file named with the current date and returns it.
func NewLogFile() *os.File {
	filename := todayFilename()
	// open an output file, this will append to the today's file if server restarted.
	f, err := os.OpenFile("logs/"+filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		panic(err)
	}

	return f
}

// FormatErrorForLog gets the stacktrace of an error and format it to be log.
// It also take into account the current context of the application.
func FormatErrorForLog(ctx context.Context, err error) string {
	stacktrace := ""

	for i := 1; ; i++ {
		_, f, l, got := runtime.Caller(i)
		if !got {
			break

		}

		stacktrace += fmt.Sprintf("%s:%d\n", f, l)
	}

	logMessage := fmt.Sprintf("Throw from a route's Handler('%s')\n", ctx.HandlerName())
	logMessage += fmt.Sprintf("At Request: %s\n", getRequestLogs(ctx))
	logMessage += fmt.Sprintf("Trace: %s\n", err)
	logMessage += fmt.Sprintf("\n%s", stacktrace)

	return logMessage
}
