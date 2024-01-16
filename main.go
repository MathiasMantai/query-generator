package main

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"

	"github.com/mathiasmantai/query-generator/src"
)

func main() {
	myApp := app.New()

	myWindow := myApp.NewWindow("Query Generator")
	size := fyne.NewSize(800, 600)
	myWindow.Resize(size)
	myWindow.SetFixedSize(true)

	entry := widget.NewEntry()
	entry.MultiLine = true

	query := widget.NewEntry()

	openFileDialogButton := widget.NewButton("Load File", func() {
		openFileDialog(myWindow, entry)
	})

	submit := widget.NewButton("Generate Query", func() {
		text := entry.Text
		toInsertBefore := "\""
		toInsertAfter := "\","
		excludeLastElement := true
		if strings.TrimSpace(text) != "" {
			fileDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
				if err != nil {
					return
				}

				if writer == nil {
					// Dialog was canceled
					return
				}
				defer writer.Close()

				content := src.Process(text, toInsertBefore, toInsertAfter, excludeLastElement)

				_, err2 := writer.Write([]byte(content))
				if err2 != nil {
					return
				}
			}, myWindow)
			fileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt", ".csv"}))
			fileDialog.Resize(fyne.NewSize(800,600))
			fileDialog.Show()
		}
	})

	test := widget.NewEntry()



	content := container.NewVBox(
		query,
		openFileDialogButton,
		entry,
		container.NewHBox(test, submit),
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func createLayout(window *fyne.Window) {

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
	fileDialog.Resize(fyne.NewSize(800,600))
	fileDialog.Show()
}