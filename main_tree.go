package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/serialt/roc/page"
)

func NewTree() (tree *widget.Tree) {

	tree = widget.NewTree(
		func(id widget.TreeNodeID) []widget.TreeNodeID {
			switch id {
			case "":
				return []widget.TreeNodeID{"a", "b", "c"}
			case "a":
				return []widget.TreeNodeID{"a1", "a2"}
			}
			return []string{}
		},
		func(id widget.TreeNodeID) bool {
			return id == "" || id == "a"
		},
		func(branch bool) fyne.CanvasObject {
			if branch {
				return widget.NewLabel("Branch template")
			}
			return widget.NewLabel("Leaf template")
		},
		func(id widget.TreeNodeID, branch bool, o fyne.CanvasObject) {
			text := id
			if branch {
				text += " (branch)"
			}
			o.(*widget.Label).SetText(text)
		})
	return
}

func NewListL() (list *widget.List) {
	data := []string{
		"工具箱",
		"中文转拼音",
		"随机密码",
	}
	list = widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		})
	list.OnSelected = func(id widget.ListItemID) {
		if data[id] == "工具箱" {
			globalWin.tabs = container.NewDocTabs(
				// container.NewTabItem("首页信息", page.WelcomeScreen()),
				container.NewTabItem("转换", page.TransScreen(globalWin.win, globalWin.app)),
				container.NewTabItem("MFA", page.Mfa(globalWin.win)),
				container.NewTabItem("中文转拼音", page.PinyinScreen()),
				container.NewTabItem("加解密", page.HashScreen(globalWin.win)),
				container.NewTabItem("IP&DNS", page.IpDNSScreen(globalWin.win)),
				container.NewTabItem("随机密码", page.PwScreen()),
			)
		} else {
			globalWin.tabs = container.NewDocTabs()
		}

		Refresh(globalWin)

		// globalWin.win.SetContent(globalWin.tabs)

		// }
	}
	return
}
