package page

import (
	"bytes"
	"encoding/json"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"sigs.k8s.io/yaml"
)

func TransScreen(w fyne.Window, app fyne.App) fyne.CanvasObject {
	header := canvas.NewText("转换", theme.Color(theme.ColorNamePrimary))
	header.TextSize = 20
	header.Alignment = fyne.TextAlignCenter

	input := widget.NewMultiLineEntry()
	input.Wrapping = fyne.TextWrapWord
	input.SetPlaceHolder("Input Text Or Read from Clipboard")

	option := []string{
		"YAML to JSON",
		"YAML Fomated",
		"YAML to TOML",
		"JSON to YAML",
		"JSON Fomated",
		"JSON to TOML",
		"JSON to PostgreSQL",
	}

	selectOP := widget.NewSelect(option, func(value string) {})
	var transData string

	encode := widget.NewButtonWithIcon("Trans", theme.MediaSkipNextIcon(), func() {

		if input.Text == "" {
			input.Text = w.Clipboard().Content()
			input.Refresh()
		}

		transData = TransData(selectOP.Selected, input.Text)
		winSub := app.NewWindow("Trans Data")
		podYamlScroll := container.NewScroll(widget.NewLabel(transData))

		bottomBox := container.NewVBox(
			widget.NewButtonWithIcon("Copy data", theme.ContentCopyIcon(), func() {
				winSub.Clipboard().SetContent(transData)
			}),
		)
		content := container.NewBorder(nil, bottomBox, nil, nil, podYamlScroll)

		winSub.SetContent(content)
		winSub.Resize(fyne.NewSize(500, 600))
		winSub.Show()
	})
	encode.Importance = widget.HighImportance
	clear := widget.NewButtonWithIcon("Clear", theme.ContentClearIcon(), func() {
		input.Text = ""
		input.Refresh()
	})
	clear.Importance = widget.MediumImportance
	bt := container.NewVBox(selectOP, encode, clear)
	return container.NewBorder(header, nil, nil, bt, input)

}

func TransData(op string, srcData string) (text string) {
	srcB := []byte(srcData)
	var tData []byte
	switch op {
	case "YAML to JSON":
		tData, _ = yaml.YAMLToJSON(srcB)
	case "YAML Fomated":
		t1, _ := yaml.YAMLToJSON(srcB)
		tData, _ = yaml.JSONToYAML(t1)
	case "YAML to TOML":
	case "JSON to YAML":
		tData, _ = yaml.JSONToYAML(srcB)
	case "JSON Fomated":
		var prettyJSON bytes.Buffer
		_ = json.Indent(&prettyJSON, srcB, "", "  ")
		tData = prettyJSON.Bytes()

	case "JSON to GO Struct":
		// // 格式化json
		// var prettyJSON bytes.Buffer
		// _ = json.Indent(&prettyJSON, srcB, "", "  ")
		// t2 := prettyJSON.Bytes()

		// var jsonData map[string]interface{}
		// _ = json.Unmarshal(t2, &jsonData)
		// generatedStruct := gengo.GenerateGoStruct(jsonData, "Data")
		// tData = []byte(generatedStruct)
	case "JSON to TOML":

	case "JSON to PostgreSQL":

	}
	text = string(tData)
	return
}
