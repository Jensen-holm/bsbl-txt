package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	win := a.NewWindow("Hello World")

	win.SetContent(widget.NewLabel("Hello World"))
	win.ShowAndRun()
}
