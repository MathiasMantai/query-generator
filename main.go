package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()

	myWindow := myApp.NewWindow("File Dialog Example")
	size := fyne.NewSize(800, 600)
	myWindow.Resize(size)

	entry := widget.NewEntry()
	entry.MultiLine = true

	openFileDialogButton := widget.NewButton("Open File Dialog", func() {
		openFileDialog(myWindow, entry)
	})



	content := container.NewVBox(
		openFileDialogButton,
		entry,
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func openFileDialog(window fyne.Window, entry *widget.Entry) {
	fileDialog := dialog.NewFileOpen(func(file fyne.URIReadCloser, err error) {
		if err != nil {
			// Handle error
			return
		}

		if file == nil {
			// Dialog was canceled
			return
		}
		defer file.Close()

		data := make([]byte, 1024)
		n, err := file.Read(data)
		if err != nil {
            // Handle error
            return
        }

		entry.SetText(string(data[:n]))
	}, window)

	fileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt", ".csv"})) // You can set filters if needed

	fileDialog.Show()
}