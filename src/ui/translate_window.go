package ui

import (

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)


func translateWin() fyne.CanvasObject {
    // 创建应用程序对象
	pic.SetMinSize(fyne.NewSize(60, 60))
	icon := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), pic, layout.NewSpacer())
	
	from := widget.NewEntry()
	from.SetPlaceHolder("please input")

	to := widget.NewEntry()
	to.SetPlaceHolder("")

	body1 := container.New(layout.NewCenterLayout(),  from)
	
	dir := container.New(layout.NewCenterLayout(), direct)

	body2 := container.New(layout.NewCenterLayout(),  to)

	btn := widget.NewButtonWithIcon("translate", theme.HomeIcon(), func() {
		
	})
	submit := container.New(layout.NewCenterLayout(), btn)

	content := container.New(layout.NewVBoxLayout(), icon, body1, dir, body2, submit)

	return content
}