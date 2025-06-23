package page

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/spf13/cast"
)

// 自定义布局器：两个子元素，第一个占 1/5 高度，第二个占 4/5 高度
type ProportionalTwoRowLayout struct{}

func (l *ProportionalTwoRowLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	if len(objects) != 2 {
		return
	}

	// 第一个区域占 1/5 高度
	firstHeight := size.Height * 4 / 10
	objects[0].Resize(fyne.NewSize(size.Width, firstHeight))
	objects[0].Move(fyne.NewPos(0, 0))

	// 第二个区域占 4/5 高度
	secondHeight := size.Height * 6 / 10
	objects[1].Resize(fyne.NewSize(size.Width, secondHeight))
	objects[1].Move(fyne.NewPos(0, firstHeight))
}

func (l *ProportionalTwoRowLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(100, 200)
}

func PwScreen() fyne.CanvasObject {
	header := canvas.NewText("随机密码生成", theme.Color(theme.ColorNamePrimary))
	header.TextSize = 20
	header.Alignment = fyne.TextAlignCenter

	output := widget.NewMultiLineEntry()
	output.MultiLine = true
	output.Wrapping = fyne.TextWrapBreak
	output.SetPlaceHolder("Output Result")

	widget.NewCheck("Check", func(on bool) {})

	standard := widget.NewCheck("标准", func(value bool) {
	})
	standard.Checked = true

	num := widget.NewCheck("数字", func(value bool) {})
	uppercase := widget.NewCheck("大写", func(value bool) {})
	lowercase := widget.NewCheck("小写", func(value bool) {})
	mixcase := widget.NewCheck("大小写", func(value bool) {})
	ambiguous := widget.NewCheck("排除歧义字符", func(value bool) {})
	symbols := widget.NewCheck("特殊字符", func(value bool) {})
	pwLen := widget.NewMultiLineEntry()
	pwLen.SetPlaceHolder("密码的长度")
	pwCount := widget.NewMultiLineEntry()
	pwCount.SetPlaceHolder("密码位数")

	length := 10
	lengthBind := binding.NewFloat()
	_ = lengthBind.Set(float64(length))

	count := 10
	countBind := binding.NewFloat()
	_ = countBind.Set(float64(count))

	slide1 := widget.NewSliderWithData(0, 64, lengthBind)
	slide1.Step = 1
	slide2 := widget.NewSliderWithData(0, 64, countBind)
	slide2.Step = 1
	lengthText := widget.NewLabelWithData(binding.FloatToStringWithFormat(lengthBind, "密码长度：%0.0f"))
	countText := widget.NewLabelWithData(binding.FloatToStringWithFormat(countBind, "密码位数：%0.0f"))

	buttons := container.NewGridWithColumns(8,
		widget.NewButton("8", func() {
			_ = lengthBind.Set(8)
		}),
		widget.NewButton("16", func() {
			_ = lengthBind.Set(16)
		}),
		widget.NewButton("32", func() {
			_ = lengthBind.Set(32)
		}),
		widget.NewButton("64", func() {
			_ = lengthBind.Set(64)
		}))

	lengthLabel := container.NewGridWithColumns(3, container.New(layout.NewFormLayout(), lengthText, slide1), buttons)

	countLabel := container.NewGridWithColumns(3, container.New(layout.NewFormLayout(), countText, slide2))

	encode := widget.NewButtonWithIcon("计算", theme.MediaSkipNextIcon(), func() {
		pw := &PW{}
		if standard.Checked {
			pw.Num = true
			pw.Uppercase = true
			pw.Lowercase = true
			pw.Ambiguous = true
		}
		if uppercase.Checked {
			pw.Uppercase = true
		}
		if lowercase.Checked {
			pw.Lowercase = true
		}
		if num.Checked {
			pw.Num = true
		}
		if ambiguous.Checked {
			pw.Ambiguous = true
		}
		if symbols.Checked {
			pw.Symbols = true
		}
		if mixcase.Checked {
			pw.MixedAlpha = true
		}

		mLen, _ := lengthBind.Get()
		mCount, _ := countBind.Get()
		output.Text = GenPW(pw, cast.ToInt(mLen), cast.ToInt(mCount))
		output.Refresh()
	})
	encode.Importance = widget.HighImportance
	clear := widget.NewButtonWithIcon("clear", theme.ContentClearIcon(), func() {
		output.Text = ""
		output.Refresh()
	})

	check1 := container.NewHBox(standard, num, lowercase, uppercase)
	check2 := container.NewHBox(mixcase, symbols, ambiguous)
	// check3 := container.NewVBox(pwLen, pwCount)
	// check4 := container.NewHBox(lengthLabel)
	checkB := container.NewBorder(nil, container.NewGridWithColumns(10, encode, clear), nil, nil, container.NewVBox(check1, check2, lengthLabel, countLabel))

	aas := container.New(&ProportionalTwoRowLayout{}, checkB, output)
	// cb := container.NewGridWithRows(3, checkB, output)

	return container.NewBorder(header, nil, nil, nil, aas)

}

type PW struct {
	Num        bool
	Symbols    bool
	Uppercase  bool
	Lowercase  bool
	MixedAlpha bool
	Ambiguous  bool
}

func GenPW(pw *PW, pwLen int, count int) (text string) {
	var num = "0123456789"
	var capitalAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var lowercaseAlpha = "abcdefghijklmnopqrstuvwxyz"
	var symbols = "!@#$%^&*()-_=+[]{}|;:,.<>?/`~"
	ambiguous := []string{"0", "1", "I", "O", "l", "o"}

	var srcText string
	if pw.Uppercase {
		srcText = srcText + capitalAlpha
	}
	if pw.Lowercase {
		srcText = srcText + lowercaseAlpha
	}
	if pw.MixedAlpha {
		srcText = capitalAlpha + lowercaseAlpha
	}
	if pw.Num {
		srcText = srcText + num
	}
	if pw.Symbols {
		srcText = srcText + symbols
	}
	if srcText == "" {
		srcText = num + lowercaseAlpha
	}
	if pw.Ambiguous {
		for _, ab := range ambiguous {
			srcText = strings.ReplaceAll(srcText, ab, "")
		}

	}
	var pwList []string
	for c := 0; c < count; c++ {
		txtTmp, _ := generateRandomDigits(pwLen, srcText)
		pwList = append(pwList, txtTmp)
	}
	fmt.Println(pwList)
	text = fmt.Sprint(strings.Join(pwList, "\n"))

	return
}

func generateRandomDigits(length int, digits string) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("length must be greater than 0")
	}

	// 生成随机字符串
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", fmt.Errorf("error generating random number: %v", err)
		}
		result[i] = digits[index.Int64()]
	}

	return string(result), nil
}
