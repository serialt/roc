package page

import (
	"bytes"
	"encoding/json"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/BurntSushi/toml"
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
		"TOML Fomated",
	}

	selectOP := widget.NewSelect(option, func(value string) {})
	var transData string

	encode := widget.NewButtonWithIcon("Trans", theme.MediaSkipNextIcon(), func() {

		if input.Text == "" {
			input.Text = app.Clipboard().Content()
			input.Refresh()
		}

		transData = TransData(selectOP.Selected, input.Text)
		winSub := app.NewWindow("Trans Data")
		podYamlScroll := container.NewScroll(widget.NewLabel(transData))

		bottomBox := container.NewVBox(
			widget.NewButtonWithIcon("Copy data", theme.ContentCopyIcon(), func() {
				app.Clipboard().SetContent(transData)
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
	var err error
	switch op {
	case "YAML to JSON":
		tData, _ = yaml.YAMLToJSON(srcB)
	case "YAML Fomated":
		t1, _ := yaml.YAMLToJSON(srcB)
		tData, _ = yaml.JSONToYAML(t1)
	case "YAML to TOML":
		tData, err = YamlToToml(srcB)
		if err != nil {
			log.Printf("YAML to TOML failed, err: %v", err)
		}
	case "JSON to YAML":
		tData, _ = yaml.JSONToYAML(srcB)
	case "JSON Fomated":
		var prettyJSON bytes.Buffer
		_ = json.Indent(&prettyJSON, srcB, "", "  ")
		tData = prettyJSON.Bytes()
	case "JSON to TOML":
		tData, err = JsonToToml(srcB)
		if err != nil {
			log.Printf("JSON to TOML failed, err: %v", err)
		}
	case "TOML Fomated":
		tData, err = TomlFormat(srcB)
		if err != nil {
			log.Printf("TOML Fomated failed, err: %v", err)
		}
	}
	text = string(tData)
	return
}

func JsonToToml(in []byte) (out []byte, err error) {
	// 解析 JSON 到 map[string]interface{}
	var data map[string]interface{}
	if err = json.Unmarshal(in, &data); err != nil {
		return
	}
	// 使用 toml.Encoder 输出 TOML 字符串
	var tomlResult bytes.Buffer
	encoder := toml.NewEncoder(&tomlResult)
	if err = encoder.Encode(data); err != nil {
		return
	}
	out = tomlResult.Bytes()
	return
}

func YamlToToml(in []byte) (out []byte, err error) {
	// 使用 map[string]interface{} 存储解析后的数据
	var data map[string]interface{}
	err = yaml.Unmarshal(in, &data)
	if err != nil {
		return
	}
	var tomlResult any = data
	out, err = toml.Marshal(tomlResult)
	if err != nil {
		return
	}
	return
}

func TomlFormat(in []byte) (out []byte, err error) {
	// Step 1: 解析 TOML 字符串到 map
	var parsed map[string]interface{}
	_, err = toml.Decode(string(in), &parsed)
	if err != nil {
		return
	}
	// Step 2: 使用 Encoder 重新生成格式化后的 TOML
	var formatted strings.Builder
	encoder := toml.NewEncoder(&formatted)
	err = encoder.Encode(parsed)
	if err != nil {
		return
	}

	return []byte(formatted.String()), nil
}
