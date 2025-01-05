package page

import (
	"fmt"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func parseURL(urlStr string) *url.URL {
	link, err := url.Parse(urlStr)
	if err != nil {
		fmt.Println("Could not parse URL", err)
	}

	return link
}

func WelcomeScreen() fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle(
			"Welcom to the Roc app",
			fyne.TextAlignCenter,
			fyne.TextStyle{Bold: true}),
		container.NewCenter(widget.NewHyperlink("github", parseURL("https://github.com/serialt"))),
	))

}
