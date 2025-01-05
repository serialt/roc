package page

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/serialt/gopkg"
)

func FileLockScreen() fyne.CanvasObject {
	selectEntry := widget.NewSelectEntry([]string{"aes", "3des", "rsa", "ecc", "sha384", "sha512"})
	selectEntry.PlaceHolder = "请选择加密的算法"
	entry := widget.NewMultiLineEntry()
	textArea := widget.NewMultiLineEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "txt", Widget: entry}},
		OnSubmit: func() { // optional, handle form submission
			log.Println("Form submitted:", entry.Text)
			log.Println("multiline:", textArea.Text)
			// var result string
			var hashStr string
			switch selectEntry.Text {
			case "md5":
				hashStr = gopkg.Md5String(entry.Text)
			case "sha1":
				hashStr = gopkg.HashSha1(entry.Text)
			case "sha256":
				hashStr = gopkg.HashSha256(entry.Text)
			case "sha384":
				hashStr = gopkg.HashSha384(entry.Text)
			case "sha512":
				hashStr = gopkg.HashSha512(entry.Text)
			default:
				hashStr = "无法计算此 txt 的 hash 值，请检查后再次进行计算"
			}
			log.Println("选择的算法:", selectEntry.Text)
			log.Println("输入的txt:", entry.Text)
			log.Println("计算的hash值:", hashStr)
			textArea.SetText(hashStr)

		},
	}
	form.Append("hash算法", selectEntry)
	form.Append("hash值", textArea)

	return container.NewVBox(form)
}
