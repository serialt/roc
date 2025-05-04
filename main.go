package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"github.com/serialt/roc/page"
	rocTheme "github.com/serialt/roc/theme"
)

// 方式一  设置环境变量   通过go-findfont 寻找simkai.ttf 字体
func init() {
	// fontPaths := findfont.List()
	// for _, path := range fontPaths {
	// 	//fmt.Println(path)
	// 	//楷体:simkai.ttf
	// 	//黑体:simhei.ttf
	// 	//思源黑体:
	// 	//  SourceHanSansSC-Bold.ttf
	// 	//  SourceHanSansSC-ExtraLight.ttf
	// 	//  SourceHanSansSC-Heavy.ttf
	// 	//  SourceHanSansSC-Light.ttf
	// 	//  SourceHanSansSC-Medium.ttf
	// 	//  SourceHanSansSC-Normal.ttf
	// 	//  SourceHanSansSC-Regular.ttf
	// 	//
	// 	if strings.Contains(path, "SourceHanSansSC-Medium.ttf") {
	// 		// fmt.Println(path)
	// 		os.Setenv("FYNE_FONT", path) // 设置环境变量  // 取消环境变量 os.Unsetenv("FYNE_FONT")
	// 		break
	// 	}
	// }
	// fmt.Println("=============")
	_ = os.Setenv("FYNE_SCALE", "0.85")
}

func main() {
	myApp := app.NewWithID("io.local.roc")
	myApp.SetIcon(theme.FyneLogo())
	win := myApp.NewWindow("roc 工具集")

	tabs := container.NewDocTabs(
		container.NewTabItem("首页信息", page.WelcomeScreen()),
		container.NewTabItem("转换", page.TransScreen(win, myApp)),
		container.NewTabItem("MFA", page.Mfa(win)),
		container.NewTabItem("中文转拼音", page.PinyinScreen()),
		container.NewTabItem("加解密", page.HashScreen(win)),
		container.NewTabItem("IP&DNS", page.IpDNSScreen(win)),
		container.NewTabItem("随机密码", page.PwScreen()),
		// container.NewTabItem("table", page.TableScreen()),

	)

	myApp.Settings().SetTheme(rocTheme.LightTheme{})
	//tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))

	// tabs.SetTabLocation(container.TabLocationTrailing)

	win.SetContent(tabs)
	win.Resize(fyne.NewSize(1200, 800))
	win.ShowAndRun()
}

// func makeMenu(fyneApp fyne.App, window fyne.Window) *fyne.MainMenu {
// 	newItem := fyne.NewMenuItem("New", nil)
// 	checkedItem := fyne.NewMainMenu("Checked", nil)
// 	disabledItem := fyne.NewMenuItem("Disabled", nil)
// 	disabledItem.Disabled = true
// 	otherItem := fyne.NewMenuItem("Other", nil)

// }

// func makeNav(setTutorial func(tutorial tutorials.Tutorial), loadPrevious bool) fyne.CanvasObject {
// 	myApp := fyne.CurrentApp()
// 	tree := &widget.Tree{}

// }
