package page

import (
	"log"
	"net"
	"slices"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/serialt/crab"
	"github.com/spf13/cast"
)

func IpDNSScreen(w fyne.Window) fyne.CanvasObject {
	app := fyne.CurrentApp()
	header := canvas.NewText("IP&DNS", theme.Color(theme.ColorNamePrimary))
	header.TextSize = 20
	header.Alignment = fyne.TextAlignCenter

	input := widget.NewMultiLineEntry()
	// input.MultiLine = true
	input.Wrapping = fyne.TextWrapWord

	input.SetPlaceHolder("Input Text Or Read from Clipboard")

	output := widget.NewMultiLineEntry()
	output.MultiLine = true
	output.Wrapping = fyne.TextWrapBreak
	output.SetPlaceHolder("Output Result")

	selectW := widget.NewSelect([]string{
		"查询本地IP",
		"子网掩码换算成数字",
		"数字换算成子网掩码",
		"判断是否是公网IP",
		"查询本地的公网IP",
	}, func(value string) {
	})
	selectW.SetSelected("查询本地IP")
	selectW.PlaceHolder = "请选择"

	encode := widget.NewButtonWithIcon("计算", theme.MediaSkipNextIcon(), func() {
		if input.Text == "" {
			input.Text = fyne.CurrentApp().Clipboard().Content()
			input.Refresh()
		}

		output.Text = Trans(selectW.Selected, input.Text)
		output.Refresh()
	})
	encode.Importance = widget.HighImportance
	clear := widget.NewButtonWithIcon("Clear", theme.ContentClearIcon(), func() {
		output.Text = ""
		output.Refresh()
		input.Text = ""
		input.Refresh()
	})
	clear.Importance = widget.MediumImportance
	copy := widget.NewButtonWithIcon("Copy", theme.ContentCutIcon(), func() {
		clipboard := app.Clipboard()
		clipboard.SetContent(output.Text)
	})

	return container.NewBorder(header, nil, nil, nil, container.NewGridWithRows(
		2,
		container.NewBorder(nil, container.NewGridWithColumns(4, selectW, encode, copy, clear), nil, nil, input), output,
	))

}

func Trans(st, in string) (out string) {
	switch st {
	case "查询本地IP":
		ips, err := getLocalIPs()
		if err != nil {
			log.Printf("get local ip failed, err: %v", err)
			return
		}
		out = strings.Join(ips, "\n")
		ipv6s, err := getLocalIPv6s()
		if err != nil {
			log.Printf("get local ip failed, err: %v", err)
			return
		}
		outv6 := strings.Join(ipv6s, "\n")
		out = out + "\n\n" + outv6

	case "子网掩码换算成数字":
		resultTmp, _ := crab.SubNetMaskToLen(in)
		out = cast.ToString(resultTmp)
	case "数字换算成子网掩码":
		intLen, _ := strconv.Atoi(in)
		out = cast.ToString(crab.LenToSubNetMask(intLen))
	case "判断是否是公网IP":
		out = cast.ToString(crab.IsPublicIPv4(net.ParseIP(in)))
	case "查询本地的公网IP":
		urlList := []string{
			"http://cip.cc",
		}

		crab.SetPubIpUrl(urlList...)
		out, _ = crab.GetPubIP()

	default:
		out = "输入或者选择有误，请检查后再提交"
	}
	return
}

// 获取所有非回环 IPv4 地址
func getLocalIPs() ([]string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	var ips []string
	for _, addr := range addrs {
		// 检查是否为 ip 地址
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil { // 只取 IPv4
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	slices.Sort(ips)
	return ips, nil
}

// getLocalIPv6s 函数用于获取所有非回环的 IPv6 地址
func getLocalIPv6s() ([]string, error) {
	var ips []string
	// 获取所有网络接口
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, iface := range interfaces {
		// 获取指定接口的所有地址
		addrs, err := iface.Addrs()
		if err != nil {
			continue // 如果无法获取该接口的地址，则跳过
		}

		for _, addr := range addrs {
			var ip net.IP

			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			default:
				continue
			}

			// 检查是否为 IPv6 并且不是回环地址
			if ip.To4() == nil && !ip.IsLoopback() {
				ips = append(ips, ip.String())
			}
		}
	}
	slices.Sort(ips)
	return ips, nil
}
