package page

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var data = [][]string{[]string{"top left", "top right"},
	[]string{"bottom left", "bottom right"}}

func TableScreen() fyne.CanvasObject {

	list := widget.NewTable(
		func() (int, int) {
			return 2, 2
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		})

	return list
}
