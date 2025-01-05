package page

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/serialt/roc/utils/mfa"
)

func Mfa(w fyne.Window) fyne.CanvasObject {

	header := canvas.NewText("2fa", theme.Color(theme.ColorNamePrimary))
	header.TextSize = 20
	header.Alignment = fyne.TextAlignCenter

	// mfaName := widget.NewEntry()
	// mfaCode := widget.NewEntry()
	// mfaCode.SetPlaceHolder("Input Text Or Read from Clipboard")
	keychain := mfa.ReadKeychain(filepath.Join(os.Getenv("HOME"), ".2fa"))
	mfaList := keychain.List()
	mfaWL := widget.NewList(
		func() int {
			return len(mfaList)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Code")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(mfaList[i])
		})
	selectW := widget.NewSelect(mfaList, func(value string) {
	})
	selectW.PlaceHolder = "请选择"
	selectW.SetSelectedIndex(0)
	// selectAll := widget.NewButton("List all", tapped func())

	output := widget.NewEntry()
	output.MultiLine = true
	output.Wrapping = fyne.TextWrapBreak
	output.SetPlaceHolder("Output Result")

	listWG := widget.NewButtonWithIcon("List all", theme.ListIcon(), func() {
		kcMap := keychain.ShowAll()
		kcSlice := []string{}
		for name, code := range kcMap {
			kcSlice = append(kcSlice, fmt.Sprintf("%-*s\t%s", 6, code, name))
		}

		output.Text = fmt.Sprint(strings.Join(kcSlice, "\n"))
		output.Refresh()
		selectW.Refresh()
		mfaWL.Refresh()
	})

	encode := widget.NewButtonWithIcon("Get Code", theme.MediaSkipNextIcon(), func() {
		// if mfaCode.Text == "" {
		// input.Text = w.Clipboard().Content()
		// input.Refresh()
		// }

		output.Text = keychain.Code(selectW.Selected)
		output.Refresh()
		selectW.Refresh()
		mfaWL.Refresh()
	})
	encode.Importance = widget.HighImportance
	clear := widget.NewButtonWithIcon("Clear", theme.ContentClearIcon(), func() {
		output.Text = ""
		output.Refresh()
	})
	clear.Importance = widget.MediumImportance
	copy := widget.NewButtonWithIcon("Copy", theme.ContentCutIcon(), func() {
		clipboard := w.Clipboard()
		clipboard.SetContent(output.Text)
	})

	selectw := container.NewHBox(selectW)

	eData := container.NewBorder(nil, nil, nil, container.NewCenter(selectw), mfaWL)

	// return container.NewGridWithRows(0, form)
	return container.NewBorder(header, nil, nil, nil, container.NewGridWithRows(
		2,
		container.NewBorder(nil, container.NewGridWithColumns(5, encode, listWG, copy, clear), nil, nil, eData), output,
	))

}
