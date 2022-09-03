package main

import (
	"context"
	"os"
	"os/exec"
	"syscall"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

type LoadReturn struct {
	Err  string `json:"err"`
	Data string `json:"data"`
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

// open some uri
func (a *App) Open(value string) string {
	defer func() {
		_ = recover()
	}()
	cmd := exec.Command("cmd", "/C", "start", value)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := cmd.Run()
	if err != nil {
		return err.Error()
	}
	return ""
}

func (a *App) Save(value string) string {
	str, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Title: "选择一个位置保存",
	})
	if err != nil {
		return err.Error()
	}
	err = os.WriteFile(str, []byte(value), 0777)
	if err != nil {
		return err.Error()
	}
	return ""
}

func (a *App) Load() LoadReturn {
	str, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "选择已导出的文件",
	})
	if err != nil {
		return LoadReturn{
			Err: err.Error(),
		}
	}
	bytes, err := os.ReadFile(str)
	if err != nil {
		return LoadReturn{
			Err: err.Error(),
		}
	}
	return LoadReturn{
		Data: string(bytes),
	}
}
