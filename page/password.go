package page

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/spf13/cast"
)

func PwScreen() fyne.CanvasObject {
	header := canvas.NewText("随机密码生成", theme.Color(theme.ColorNamePrimary))
	header.TextSize = 20
	header.Alignment = fyne.TextAlignCenter

	output := widget.NewEntry()
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
	pwLen := widget.NewEntry()
	pwLen.SetPlaceHolder("密码的长度")
	pwCount := widget.NewEntry()
	pwCount.SetPlaceHolder("密码位数")

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
		nLen := cast.ToInt(pwLen.Text)
		if nLen == 0 {
			nLen = 18
		}
		nCount := cast.ToInt(pwCount.Text)
		if nCount == 0 {
			nCount = 10
		}
		output.Text = GenPW(pw, nLen, nCount)
		output.Refresh()
	})
	encode.Importance = widget.HighImportance
	clear := widget.NewButtonWithIcon("clear", theme.ContentClearIcon(), func() {
		output.Text = ""
		output.Refresh()
	})
	check1 := container.NewVBox(standard, num, lowercase, uppercase)
	check2 := container.NewVBox(mixcase, symbols, ambiguous)
	check3 := container.NewVBox(pwLen, pwCount)
	checkB := container.NewBorder(nil, container.NewGridWithColumns(3, encode, clear), nil, nil, container.NewGridWithColumns(4, check1, check2, container.NewGridWithColumns(1, check3)))

	cb := container.NewGridWithRows(2, checkB, output)

	return container.NewBorder(header, nil, nil, nil, cb)

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
