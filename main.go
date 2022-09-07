package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:            "smart-tool",
		Width:            1024,
		Height:           768,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			Theme: windows.SystemDefault,
			CustomTheme: &windows.ThemeSettings{
				DarkModeTitleBar:         windows.RGB(24, 24, 28),
				DarkModeTitleBarInactive: windows.RGB(24, 24, 28),
				DarkModeTitleText:        windows.RGB(200, 200, 200),
				DarkModeBorder:           windows.RGB(24, 24, 28),
				DarkModeBorderInactive:   windows.RGB(24, 24, 28),
				LightModeTitleBar:        windows.RGB(255, 255, 255),
				LightModeTitleText:       windows.RGB(20, 20, 20),
				LightModeBorder:          windows.RGB(255, 255, 255),
			},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
