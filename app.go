package main

import (
	"context"
	"dumbapp/actions"
	"strconv"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Schedule(timer string, action string) {
	if timer == "" || action == "" {
		_, _ = runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type:          runtime.ErrorDialog,
			Title:         "Error",
			Message:       "Please select a timer and action button",
			DefaultButton: "OK",
		})
		return
	}

	timerInt, err := strconv.Atoi(timer)

	if err != nil {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type:          runtime.ErrorDialog,
			Title:         "Error",
			Message:       "Invalid timer value, please enter the value in seconds",
			DefaultButton: "OK",
		})
		return
	}

	switch action {
	case "Shutdown":
		actions.Shutdown(a.ctx, timerInt)
	case "Restart":
		actions.Reboot(a.ctx, timerInt)
	case "Sleep":
		actions.Sleep(a.ctx, timerInt)
	default:
		actions.Arbitrary(a.ctx, timerInt, action)
	}

	runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:          runtime.InfoDialog,
		Title:         "Scheduled",
		Message:       "Your Task Has Been Scheduled",
		DefaultButton: "OK",
	})

}
