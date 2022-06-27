package main

import (
	"bufio"
	"fmt"
	"github.com/gek64/displayController"
	"os"
	"strconv"
	"syscall"
)

func runDemo() (err error) {
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
		return fmt.Errorf("找不到任何物理显示器")
	}
	fmt.Printf("找到 %d 个显示器\n", len(monitors))

	for _, monitor := range monitors {
		fmt.Printf("显示器 handle 编号: %d\n", monitor.PhysicalInfo.Handle)
		fmt.Printf("显示器描述: %s\n", monitor.PhysicalInfo.Description)

		current, max, err := displayController.GetVCPFeatureAndVCPFeatureReply(monitor.PhysicalInfo.Handle, displayController.Brightness)
		if err != nil {
			return err
		}

		fmt.Printf("显示器当前亮度: %d, 最大亮度: %d\n", current, max)
		fmt.Println()
	}

	return nil
}

func selectMonitor() (useAll bool, handle int, err error) {
	buf := bufio.NewReader(os.Stdin)
	fmt.Println("输入你需要控制的显示器 handle 编号,直接回车默认选择所有显示器")
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
		return false, -1, fmt.Errorf("选择显示器错误,请重新输入你需要控制的显示器")
	}
	return false, handle, nil
}

func setMonitor(useAll bool, handle int) (err error) {
	buf := bufio.NewReader(os.Stdin)
	fmt.Println("输入设置的新的亮度")
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
