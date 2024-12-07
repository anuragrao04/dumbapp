package actions

import (
	"context"
	"dumbapp/scheduler"
	"runtime"
)

func Sleep(appContext context.Context, timer int) {
	sleepCommand := GetSleepCommand()
	scheduler.Schedule(appContext, sleepCommand, timer)
}

func GetSleepCommand() []string {
	switch runtime.GOOS {
	case "windows":
		return []string{"rundll32.exe", "powrprof.dll,SetSuspendState", "Sleep"}
	case "darwin": // macOS
		return []string{"pmset", "sleepnow"}
	default: // Linux/Unix
		return []string{"systemctl", "suspend"}
	}
}
