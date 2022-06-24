package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	selectedID         int  = 0
	selectedAll        bool = false
	brightnessVCPCode  byte = 0x10
	brightnessVCPValue int  = 0
)

func runDemo() (err error) {
	err = showBasicInfo()
	if err != nil {
		return err
	}

	err = selectMonitor()
	if err != nil {
		return err
	}

	return nil
}

func showBasicInfo() (err error) {
	if len(monitors) == 0 {
		return fmt.Errorf("找不到任何物理显示器")
	} else {
		fmt.Printf("找到 %d 个显示器\n", len(monitors))
	}

	for _, monitor := range monitors {
		fmt.Printf("显示器 ID: %d\n", monitor.id)
		fmt.Printf("显示器描述: %s\n\n", monitor.physicalInfo.Description)
		err := getBrightness(monitor.id)
		if err != nil {
			return err
		}
	}

	return nil
}

func selectMonitor() (err error) {
	buf := bufio.NewReader(os.Stdin)
	fmt.Println("输入你需要控制的显示器")
	fmt.Println("请输入上述找到的显示器对应的ID,或输入all来选择全部显示器,默认为选择全部显示器")
	fmt.Printf("> ")
	line, _, err := buf.ReadLine()
	if err != nil {
		return err
	} else {
		selectedID, err = strconv.Atoi(string(line))
		if err != nil {
			if string(line) == "all" || string(line) == "" {
				selectedAll = true
			} else {
				return fmt.Errorf("选择显示器错误,请重新输入你需要控制的显示器")
			}
		}
	}
	return nil
}

func getBrightness(id int) (err error) {
	return getDisplayByID(id, brightnessVCPCode)
}
