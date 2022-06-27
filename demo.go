package main

import (
	"bufio"
	"fmt"
	"github.com/gek64/displayController"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"syscall"
)

var (
	LANGZH = []string{
		"找不到任何物理显示器",
		"找到 %d 个显示器\n",
		"显示器 handle 编号: %d\n",
		"显示器描述: %s\n",
		"显示器当前亮度: %d, 最大亮度: %d\n",
		"输入你需要控制的显示器 handle 编号,直接回车默认选择所有显示器",
		"选择显示器错误,请重新输入你需要控制的显示器",
		"输入设置的新的亮度",
	}
	LANGEN = []string{
		"can't find any physical display",
		"found %d display monitors\n",
		"display monitor handle No: %d\n",
		"display monitor description: %s\n",
		"display monitor current brightness value: %d, maximum brightness value: %d\n",
		"enter the display handle No. you need to control, or select all the display by press ENTER",
		"select the display error, please re-enter the display you need to control",
		"enter the new brightness of the display",
	}
	LANG []string
)

func runDemo() (err error) {
	if lang, _ := getLocale(); lang == "zh" {
		LANG = LANGZH
	} else {
		LANG = LANGEN
	}

	err = showBasicInfo()
	if err != nil {
		return err
	}

	useAll, handle, err := selectMonitor()
	if err != nil {
		return err
	}

	err = setMonitor(useAll, handle)
	if err != nil {
		return err
	}

	return nil
}

func showBasicInfo() (err error) {
	if len(monitors) == 0 {
		return fmt.Errorf(LANG[0])
	}
	fmt.Printf(LANG[1], len(monitors))

	for _, monitor := range monitors {
		fmt.Printf(LANG[2], monitor.PhysicalInfo.Handle)
		fmt.Printf(LANG[3], monitor.PhysicalInfo.Description)

		current, max, err := displayController.GetVCPFeatureAndVCPFeatureReply(monitor.PhysicalInfo.Handle, displayController.Brightness)
		if err != nil {
			return err
		}

		fmt.Printf(LANG[4], current, max)
		fmt.Println()
	}
	return nil
}

func selectMonitor() (useAll bool, handle int, err error) {
	buf := bufio.NewReader(os.Stdin)
	fmt.Println(LANG[5])
	fmt.Printf("> ")

	line, _, err := buf.ReadLine()
	if err != nil {
		return false, -1, err
	}

	if string(line) == "" {
		return true, -1, nil
	}

	handle, err = strconv.Atoi(string(line))
	if err != nil {
		return false, -1, fmt.Errorf(LANG[6])
	}
	return false, handle, nil
}

func setMonitor(useAll bool, handle int) (err error) {
	buf := bufio.NewReader(os.Stdin)
	fmt.Println(LANG[7])
	fmt.Printf("> ")

	line, _, err := buf.ReadLine()
	if err != nil {
		return err
	}

	newValue, err := strconv.Atoi(string(line))
	if err != nil {
		return err
	}

	if useAll {
		for _, monitor := range monitors {
			err := displayController.SetVCPFeature(monitor.PhysicalInfo.Handle, displayController.Brightness, newValue)
			if err != nil {
				return err
			}
		}
	}

	if handle != -1 {
		err := displayController.SetVCPFeature(syscall.Handle(handle), displayController.Brightness, newValue)
		if err != nil {
			return err
		}
	}

	return nil
}

// get system language and location
// https://stackoverflow.com/a/64560642
func getLocale() (string, string) {
	osHost := runtime.GOOS
	defaultLang := "en"
	defaultLoc := "US"
	switch osHost {
	case "windows":
		// Exec powershell Get-Culture on Windows.
		cmd := exec.Command("powershell", "Get-Culture | select -exp Name")
		output, err := cmd.Output()
		if err == nil {
			langLocRaw := strings.TrimSpace(string(output))
			langLoc := strings.Split(langLocRaw, "-")
			lang := langLoc[0]
			loc := langLoc[1]
			return lang, loc
		}
	case "darwin":
		// Exec powershell Get-Culture on macOS.
		cmd := exec.Command("sh", "osascript -e 'user locale of (get system info)'")
		output, err := cmd.Output()
		if err == nil {
			langLocRaw := strings.TrimSpace(string(output))
			langLoc := strings.Split(langLocRaw, "_")
			lang := langLoc[0]
			loc := langLoc[1]
			return lang, loc
		}
	case "linux":
		envLang, ok := os.LookupEnv("LANG")
		if ok {
			langLocRaw := strings.TrimSpace(envLang)
			langLocRaw = strings.Split(envLang, ".")[0]
			langLoc := strings.Split(langLocRaw, "_")
			lang := langLoc[0]
			loc := langLoc[1]
			return lang, loc
		}
	}
	return defaultLang, defaultLoc
}
