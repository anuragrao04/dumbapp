package actions

import (
	"context"
	"dumbapp/scheduler"
	"runtime"
)

func Shutdown(appContext context.Context, timer int) {
	shutdownCommand := GetShutdownCommand()
	scheduler.Schedule(appContext, shutdownCommand, timer)
}

func GetShutdownCommand() []string {
	switch runtime.GOOS {
	case "windows":
		return []string{"cmd", "shutdown", "/s", "/t", "0"}
	case "darwin": // macOS
		return []string{"osascript", "-e", "tell application \"System Events\" to shut down"}
	default: // Linux/Unix
		return []string{"shutdown", "-h", "now"}
	}
}
