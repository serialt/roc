package page

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/serialt/crab"
)

// switch st {
// 	case "md5":
// 		out = cryptor.Md5String(in)
// 	case "sha1":
// 		out = cryptor.Sha1(in)
// 	case "sha256":
// 		out = cryptor.Sha256(in)
// 	case "sha512":
// 		out = cryptor.Sha512(in)
// 	case "base64 decode":
// 		out = cryptor.Base64StdDecode(in)
// 	case "base64 encode":
// 		out = cryptor.Base64StdEncode(in)
// 	case "url encode":
// 	case "url decode":
// 	// case "JSON 转义":
// 	// case "JSON 转义":

type Data struct {
	Code   string
	Action []string
}

var dataList = []*Data{
	{Code: "Base64", Action: []string{"decode", "encode"}},
	{Code: "URL", Action: []string{"decode", "encode"}},
	{Code: "JSON 转义", Action: []string{"decode", "encode"}},
	{Code: "MD5", Action: []string{"encode"}},
	{Code: "SHA1", Action: []string{"encode"}},
	{Code: "SHA256", Action: []string{"encode"}},
	{Code: "SHA512", Action: []string{"encode"}},
}

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
			input.Text = fyne.CurrentApp().Clipboard().Content()
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
		clipboard := fyne.CurrentApp().Clipboard()
		clipboard.SetContent(output.Text)
	})

	pd2 := container.NewGridWithColumns(3, encode, copy, clear)
	pd := container.NewGridWithColumns(3, selectW, pd2, layout.NewSpacer())

	return container.NewBorder(header, nil, nil, nil, container.NewGridWithRows(
		2,
		container.NewBorder(nil, pd, nil, nil, input), output,
	))

}

func Hash(st string, in string) (out string) {
	switch st {
	case "md5":
		out = crab.Md5String(in)
	case "sha1":
		out = crab.Sha1(in)
	case "sha256":
		out = crab.Sha256(in)
	case "sha512":
		out = crab.Sha512(in)
	case "base64 decode":
		out = crab.Base64StdDecode(in)
	case "base64 encode":
		out = crab.Base64StdEncode(in)
	case "url encode":
	case "url decode":
	// case "JSON 转义":
	// case "JSON 转义":

	default:
		out = "无法计算此 txt 的值，请检查"
	}
	return
}
