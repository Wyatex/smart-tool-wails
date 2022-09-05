package main

import (
	"context"
	"os"
	"os/exec"
	"strings"
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
	args := append([]string{"/C", "start"}, strings.Split(value, " ")...)
	cmd := exec.Command("cmd", args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := cmd.Run()
	if err != nil {
		return err.Error()
	}
	return ""
}

func (a *App) Save(value string) string {
	str, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Title:           "选择一个位置保存",
		DefaultFilename: "backup.json",
	})
	if err != nil {
		return err.Error()
	}
	if str == "" {
		return "已取消导出"
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
		Filters: []runtime.FileFilter{
			{
				DisplayName: "备份数据",
				Pattern:     "*.json",
			},
		},
	})
	if err != nil {
		return LoadReturn{
			Err: err.Error(),
		}
	}
	if str == "" {
		return LoadReturn{
			Err: "已取消导入",
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
