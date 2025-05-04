package page

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/mozillazg/go-pinyin"
)

func DevScreen() fyne.CanvasObject {

	header := canvas.NewText("Debug", theme.Color(theme.ColorNamePrimary))
	header.TextSize = 20
	header.Alignment = fyne.TextAlignCenter

	input := widget.NewEntry()
	input.MultiLine = true
	input.Wrapping = fyne.TextWrapWord

	input.SetPlaceHolder("Input Text Or Read from Clipboard")

	output := widget.NewEntry()
	output.MultiLine = true
	output.Wrapping = fyne.TextWrapBreak
	output.SetPlaceHolder("Output Result")

	encode := widget.NewButtonWithIcon("Encode", theme.MediaSkipNextIcon(), func() {
		if input.Text == "" {
			// input.Text = w.Clipboard().Content()
			// input.Refresh()
		}
		pin := fmt.Sprint(pinyin.LazyConvert(input.Text, nil))
		ss := strings.ReplaceAll(pin, "[", "")
		ss = strings.ReplaceAll(ss, "]", "")

		// 去掉汉子拼音中的空格
		ss2 := strings.ReplaceAll(ss, " ", "")
		output.Text = fmt.Sprintf("%v\n\n%v", ss, ss2)
		output.Refresh()
	})
	encode.Importance = widget.HighImportance
	clear := widget.NewButtonWithIcon("clear", theme.ContentClearIcon(), func() {
		output.Text = ""
		output.Refresh()
		input.Text = ""
		input.Refresh()
	})
	clear.Importance = widget.MediumImportance

	// return container.NewGridWithRows(0, form)
	return container.NewBorder(header, nil, nil, nil, container.NewGridWithRows(
		2,
		container.NewBorder(nil, container.NewGridWithColumns(2, encode, clear), nil, nil, input), output,
	))

}
