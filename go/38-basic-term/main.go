package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	fyneApp := app.New()
	windowsApp := fyneApp.NewWindow("germ")

	termUi := widget.NewTextGrid()
	termUi.SetText("I'm on a Terminal")

	windowsApp.SetContent(
		container.New(
			layout.NewGridWrapLayout(fyne.NewSize(420, 200)),
			termUi,
		),
	)

	windowsApp.ShowAndRun()
}
