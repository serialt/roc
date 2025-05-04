package page

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/serialt/roc/utils/mfa"
)

func Mfa(w fyne.Window) fyne.CanvasObject {

	header := canvas.NewText("2fa", theme.Color(theme.ColorNamePrimary))
	header.TextSize = 20
	header.Alignment = fyne.TextAlignCenter

	keychain := mfa.ReadKeychain(filepath.Join(os.Getenv("HOME"), ".2fa"))
	mfaList := keychain.List()
	mfaWL := GetListData(&mfaList)

	output := widget.NewMultiLineEntry()
	output.MultiLine = true
	output.Wrapping = fyne.TextWrapBreak
	output.SetPlaceHolder("Output Result")

	mfaWL.OnSelected = func(id widget.ListItemID) {
		output.Text = keychain.Code(mfaList[id])
		output.Refresh()
	}

	listWG := widget.NewButtonWithIcon("List all", theme.ListIcon(), func() {
		kcMap := keychain.ShowAll()
		kcSlice := []string{}
		for name, code := range kcMap {
			kcSlice = append(kcSlice, fmt.Sprintf("%-*s\t%s", 6, code, name))
		}

		output.Text = fmt.Sprint(strings.Join(kcSlice, "\n"))
		output.Refresh()
	})
	clear := widget.NewButtonWithIcon("Clear", theme.ContentClearIcon(), func() {
		output.Text = ""
		output.Refresh()
	})
	clear.Importance = widget.MediumImportance
	copy := widget.NewButtonWithIcon("Copy", theme.ContentCutIcon(), func() {
		clipboard := fyne.CurrentApp().Clipboard()
		clipboard.SetContent(output.Text)
	})

	eData := container.NewBorder(nil, nil, nil, nil, mfaWL)
	return container.NewBorder(header, nil, nil, nil, container.NewGridWithRows(
		2,
		container.NewBorder(nil, container.NewGridWithColumns(5, listWG, copy, clear), nil, nil, eData), output,
	))

}

func GetListData(nameData *[]string) *widget.List {
	// list binding, bind pod list data to data
	data := binding.BindStringList(
		nameData,
	)

	list := widget.NewListWithData(data,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})

	return list
}
