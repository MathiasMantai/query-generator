package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	root := app.New()
	window := root.NewWindow("Root")

	queryTemplate := widget.NewEntry()
	queryTemplate.SetPlaceHolder("Enter Query Template")

	content := container.NewVBox(queryTemplate)

	window.SetContent(content)
	window.ShowAndRun()
}