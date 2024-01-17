package main

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/layout"
	"github.com/mathiasmantai/query-generator/src"
)

func main() {
	myApp := app.New()

	myWindow := myApp.NewWindow("Query Generator")
	size := fyne.NewSize(850, 600)
	myWindow.Resize(size)
	myWindow.SetFixedSize(true)

	//multiline input for data
	entry := widget.NewEntry()
	entry.MultiLine = true

	//input for query
	query := widget.NewEntry()

	openFileDialogButton := widget.NewButton("Load Data from File", func() {
		openFileDialog(myWindow, entry)
	})

	//insertBefore input
	insertBeforeInput := widget.NewEntry()

    //insertAfter input
    insertAfterInput := widget.NewEntry()

	//ignore last element
	ignoreLastElement := widget.NewCheck("Ignore Last Element", func(value bool) {

    })

	//use data for IN clause
	useForInClause := widget.NewCheck("Use Data for IN clause", func(value bool) {

    })


	output := widget.NewEntry()
	output.MultiLine = true


	//submit button to generate queries
	submit := widget.NewButton("Generate Query", func() {
		text := entry.Text
		toInsertBefore := insertBeforeInput.Text
		toInsertAfter := insertAfterInput.Text
		excludeLastElement := ignoreLastElement.Checked

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
			fileDialog.Resize(fyne.NewSize(850,600))
			fileDialog.Show()
		}
	})

	//process button to generate queries and display in output input field
	processButton := widget.NewButton("Process Query", func() {
		text := entry.Text
		toInsertBefore := insertBeforeInput.Text
		toInsertAfter := insertAfterInput.Text
		excludeLastElement := ignoreLastElement.Checked
		content := src.Process(text, toInsertBefore, toInsertAfter, excludeLastElement)
		output.SetText(content)
	})


	grid := container.New(layout.NewFormLayout(),
		widget.NewLabel("Query:"),
		query,
		widget.NewLabel("Data:"),
        entry,
		widget.NewLabel(""),
		openFileDialogButton,
		widget.NewLabel("Options:"),
		container.New(
			layout.NewGridLayout(2),
			container.New(
				layout.NewFormLayout(),
				widget.NewLabel("Insert Before every Element:"),
				insertBeforeInput,
			),
			container.New(
				layout.NewFormLayout(),
				widget.NewLabel("Insert After every Element:"),
				insertAfterInput,
				widget.NewLabel("Ignore Last Element:"),
				ignoreLastElement,
			),
			container.New(
				layout.NewFormLayout(),
				widget.NewLabel("Use For IN Clause:"),
				useForInClause,
			),
		),
		widget.NewLabel(""),
        submit,
		widget.NewLabel(""),
		processButton,
		widget.NewLabel(""),
		output,
	)

	myWindow.SetContent(grid)
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
	fileDialog.Resize(fyne.NewSize(850,600))
	fileDialog.Show()
}