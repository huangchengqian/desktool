package ui

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)


var pic *canvas.Image
var direct *canvas.Image

func MainWindow() {


    // 创建应用程序对象
	mainApp := app.New()
    window := mainApp.NewWindow("tool")
    // 创建窗口
	tabs := container.NewAppTabs(
		container.NewTabItem("翻译", translateWin()),
		container.NewTabItem("Tab 2", widget.NewLabel("World!")),
	)
	tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))

	tabs.SetTabLocation(container.TabLocationLeading)

    window.SetContent(tabs)
    window.Resize(fyne.NewSize(1200, 800))
    window.ShowAndRun()
}
