package scheduler

import (
	"context"
	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
	"os/exec"
	"time"
)

func Schedule(appContext context.Context, command []string, timer int) {
	time.AfterFunc(time.Duration(timer)*time.Second, func() {
		err := exec.Command(command[0], command[1:]...).Run()
		if err != nil {
			wailsRuntime.MessageDialog(appContext, wailsRuntime.MessageDialogOptions{
				Type:          wailsRuntime.ErrorDialog,
				Title:         "Error",
				Message:       "Failed To Execute Command",
				DefaultButton: "OK",
			})
			return
		}
	})
}
