package actions

import (
	"context"
	"dumbapp/scheduler"
	"runtime"
)

func Reboot(appContext context.Context, timer int) {
	rebootCommand := GetRebootCommand()
	scheduler.Schedule(appContext, rebootCommand, timer)
}

func GetRebootCommand() []string {
	switch runtime.GOOS {
	case "windows":
		return []string{"cmd", "shutdown", "/r", "/t", "0"}
	case "darwin": // macOS
		return []string{"osascript", "-e", "tell application \"System Events\" to restart"}
	default: // Linux/Unix
		return []string{"shutdown", "-r", "now"}
	}
}
