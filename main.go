package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
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
	mainWin := new(MainWin)
	globalWin = mainWin
	mainWin.app = app.NewWithID((AppID))

	mainWin.win = mainWin.app.NewWindow("roc 工具集")
	mainWin.win.Resize(fyne.NewSize(WindowWidth, WindowHeight))
	mainWin.win.SetPadded(false)
	mainWin.win.SetMaster()      //退出窗体则退出程序
	mainWin.win.CenterOnScreen() //屏幕中央
	mainWin.list = NewListL()
	mainWin.tabs = container.NewDocTabs(
		container.NewTabItem("首页信息", page.WelcomeScreen()),
	)

	// mainWin.list.OnSelected = func(id widget.ListItemID) {
	// 	// if id == 0 {
	// 	globalWin.tabs = container.NewDocTabs(
	// 		// container.NewTabItem("首页信息", page.WelcomeScreen()),
	// 		container.NewTabItem("转换", page.TransScreen(mainWin.win, mainWin.app)),
	// 		container.NewTabItem("MFA", page.Mfa(mainWin.win)),
	// 		container.NewTabItem("中文转拼音", page.PinyinScreen()),
	// 		container.NewTabItem("加解密", page.HashScreen(mainWin.win)),
	// 		container.NewTabItem("IP&DNS", page.IpDNSScreen(mainWin.win)),
	// 		container.NewTabItem("随机密码", page.PwScreen()),
	// 	)

	// 	Refresh(globalWin)

	// 	// globalWin.win.SetContent(globalWin.tabs)

	// 	// }
	// }
	mainWin.app.Settings().SetTheme(rocTheme.LightTheme{})
	//tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))

	// tabs.SetTabLocation(container.TabLocationLeading)

	leftBtnBox := container.NewHBox(widget.NewLabel("List"), layout.NewSpacer())
	leftCard := container.NewBorder(leftBtnBox, nil, nil, nil, mainWin.list) //边框
	leftPanel := widget.NewCard("", "", leftCard)

	// CONTENT
	content := container.NewHSplit(leftPanel, mainWin.tabs)
	content.SetOffset(0.15)
	home := container.NewBorder(nil, nil, nil, nil, content)
	mainWin.win.SetContent(home)

	mainWin.win.ShowAndRun()
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

type MainWin struct {
	app      fyne.App
	win      fyne.Window
	tabs     *container.DocTabs
	tree     *widget.Tree
	list     *widget.List
	mainMenu *fyne.MainMenu
}

var (
	WindowWidth  float32 = 900
	WindowHeight float32 = 600
	AppID                = "io.local.roc"
	globalWin    *MainWin
)

func Refresh(mainWin *MainWin) {
	leftBtnBox := container.NewHBox(widget.NewLabel("List"), layout.NewSpacer())
	leftCard := container.NewBorder(leftBtnBox, nil, nil, nil, mainWin.list) //边框
	leftPanel := widget.NewCard("", "", leftCard)

	// CONTENT
	content := container.NewHSplit(leftPanel, mainWin.tabs)
	content.SetOffset(0.15)
	mainWin.win.Content().Refresh()
	mainWin.tabs.Refresh()
	home := container.NewBorder(nil, nil, nil, nil, content)
	mainWin.win.SetContent(home)

}
