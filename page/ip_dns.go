package page

import (
	"net"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/serialt/gopkg"
	"github.com/spf13/cast"
)

func IpDNSScreen(w fyne.Window) fyne.CanvasObject {
	app := fyne.CurrentApp()
	header := canvas.NewText("IP&DNS", theme.Color(theme.ColorNamePrimary))
	header.TextSize = 20
	header.Alignment = fyne.TextAlignCenter

	input := widget.NewMultiLineEntry()
	// input.MultiLine = true
	input.Wrapping = fyne.TextWrapWord

	input.SetPlaceHolder("Input Text Or Read from Clipboard")

	output := widget.NewMultiLineEntry()
	output.MultiLine = true
	output.Wrapping = fyne.TextWrapBreak
	output.SetPlaceHolder("Output Result")

	selectW := widget.NewSelect([]string{
		"查询本地IP",
		"子网掩码换算成数字",
		"数字换算成子网掩码",
		"判断是否是公网IP",
		"查询本地的公网IP",
	}, func(value string) {
	})
	selectW.SetSelected("查询本地IP")
	selectW.PlaceHolder = "请选择"

	encode := widget.NewButtonWithIcon("计算", theme.MediaSkipNextIcon(), func() {
		if input.Text == "" {
			// input.Text = w.Clipboard().Content()
			// input.Refresh()
		}

		output.Text = Trans(selectW.Selected, input.Text)
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
		clipboard := app.Clipboard()
		clipboard.SetContent(output.Text)
	})

	return container.NewBorder(header, nil, nil, nil, container.NewGridWithRows(
		2,
		container.NewBorder(nil, container.NewGridWithColumns(4, selectW, encode, copy, clear), nil, nil, input), output,
	))

}

func Trans(st, in string) (out string) {
	switch st {
	case "查询本地IP":
	case "子网掩码换算成数字":
		resultTmp, _ := gopkg.SubNetMaskToLen(in)
		out = cast.ToString(resultTmp)
	case "数字换算成子网掩码":
		intLen, _ := strconv.Atoi(in)
		out = cast.ToString(gopkg.LenToSubNetMask(intLen))
	case "判断是否是公网ip":
		out = cast.ToString(gopkg.IsPublicIPv4(net.ParseIP(in)))
	case "查询本地的公网ip":
		out, _ = gopkg.GetPubIP()
	default:
		out = "输入或者选择有误，请检查后再提交"
	}
	return
}
