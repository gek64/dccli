package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	selectedID      int    = -1
	selectedAll     bool   = false
	vcpCode         byte   = 0x0
	vcpFeature      string = ""
	vcpFeatureValue int    = -1
)

func console() (err error) {
	err = showBasicInfo()
	if err != nil {
		return err
	}
	waitCommand()

	return nil
}

func showBasicInfo() (err error) {
	if len(monitors) == 0 {
		return fmt.Errorf("can't find any physical display monitor")
	} else {
		fmt.Printf("Find %d Display Monitor\n", len(monitors))
	}

	for _, monitor := range monitors {
		fmt.Printf("Display Monitor ID: %d\n", monitor.id)
		fmt.Printf("Display Monitor Description: %s\n\n", monitor.physicalInfo.Description)
	}

	return nil
}

func waitCommand() {
	buf := bufio.NewReader(os.Stdin)
	fmt.Println("请输入你需要控制的显示器ID,或选择全部,ID请输入上述屏幕对应的,选择全部请输入all")
	fmt.Print("> ")
	sentence, err := buf.ReadBytes('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(sentence))
	}
}
