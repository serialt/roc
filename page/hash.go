package page

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/duke-git/lancet/v2/cryptor"
)

func HashScreen(w fyne.Window) fyne.CanvasObject {
	header := canvas.NewText("加解密计算", theme.Color(theme.ColorNamePrimary))
	header.TextSize = 20
	header.Alignment = fyne.TextAlignCenter

	input := widget.NewMultiLineEntry()
	input.MultiLine = true
	input.Wrapping = fyne.TextWrapWord

	input.SetPlaceHolder("Input Text Or Read from Clipboard")

	output := widget.NewMultiLineEntry()
	output.MultiLine = true
	output.Wrapping = fyne.TextWrapBreak
	output.SetPlaceHolder("Output Result")

	selectW := widget.NewSelect([]string{"base64 encode", "base64 decode", "md5", "sha1", "sha256", "sha512"}, func(value string) {
	})
	selectW.PlaceHolder = "请选择算法"
	selectW.SetSelected("base64 decode")

	encode := widget.NewButtonWithIcon("计算", theme.MediaSkipNextIcon(), func() {
		if input.Text == "" {
			input.Text = w.Clipboard().Content()
			input.Refresh()
		}
		output.Text = Hash(selectW.Selected, input.Text)
		output.Refresh()
	})
	encode.Importance = widget.HighImportance
	clear := widget.NewButtonWithIcon("Clear", theme.ContentClearIcon(), func() {
		output.Text = ""
		output.Refresh()
		input.Text = ""
		input.Refresh()
	})
	clear.Importance = widget.MediumImportance

	copy := widget.NewButtonWithIcon("Copy", theme.ContentCutIcon(), func() {
		clipboard := w.Clipboard()
		clipboard.SetContent(output.Text)
	})

	return container.NewBorder(header, nil, nil, nil, container.NewGridWithRows(
		2,
		container.NewBorder(nil, container.NewGridWithColumns(5, selectW, encode, copy, clear), nil, nil, input), output,
	))

}

func Hash(st, in string) (out string) {
	switch st {
	case "md5":
		out = cryptor.Md5String(in)
	case "sha1":
		out = cryptor.Sha1(in)
	case "sha256":
		out = cryptor.Sha256(in)
	case "sha512":
		out = cryptor.Sha512(in)
	case "base64 decode":
		out = cryptor.Base64StdDecode(in)
	case "base64 encode":
		out = cryptor.Base64StdEncode(in)
	default:
		out = "无法计算此 txt 的值，请检查"
	}
	return
}
