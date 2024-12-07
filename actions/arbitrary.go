package actions

import (
	"context"
	"dumbapp/scheduler"
	"runtime"
)

func Arbitrary(appContext context.Context, timer int, command string) {
	commandSlice := GetCommandExecutionSlice(command)
	scheduler.Schedule(appContext, commandSlice, timer)
}

func GetCommandExecutionSlice(command string) []string {
	switch runtime.GOOS {
	case "windows":
		// Executes the command using cmd.exe
		return []string{"cmd", "/C", command}
	default:
		// Executes the command using bash
		return []string{"bash", "-c", command}
	}
}
