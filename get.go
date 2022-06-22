package main

import (
	"fmt"
	"github.com/gek64/displayController"
	"syscall"
)

func getDisplayByID(id int, vcpCode byte) (err error) {
	// 找到了标志物
	var found = false

	if len(monitors) == 0 {
		return fmt.Errorf("can't find any physical display monitor")
	}

	for _, monitor := range monitors {
		if monitor.id == id {
			found = true
			err := getVCPFeatureValue(monitor, vcpCode)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
		// 找到了就跳出循环
		if found {
			break
		}
	}

	if found {
		return nil
	} else {
		return fmt.Errorf("can't find any physical display monitor with id %d\n", id)
	}

}

func getDisplayByHandle(handle int, vcpCode byte) (err error) {
	// 找到了标志物
	var found = false

	if len(monitors) == 0 {
		return fmt.Errorf("can't find any physical display monitor")
	}

	for _, monitor := range monitors {
		if monitor.physicalInfo.Handle == syscall.Handle(handle) {
			found = true
			err := getVCPFeatureValue(monitor, vcpCode)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
		// 找到了就跳出循环
		if found {
			break
		}
	}

	if found {
		return nil
	} else {
		return fmt.Errorf("can't find any physical display monitor with handle %d\n", handle)
	}
}

func getAllDisplay(vcpCode byte) (err error) {
	if len(monitors) == 0 {
		return fmt.Errorf("can't find any physical display monitor")
	}

	for _, monitor := range monitors {
		err := getVCPFeatureValue(monitor, vcpCode)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return nil
}

func getVCPFeatureValue(monitor monitor, vcpCode byte) (err error) {
	fmt.Printf("Display Monitor ID: %d\n", monitor.id)
	fmt.Printf("Display Monitor Handle ID: %d\n", monitor.physicalInfo.Handle)
	fmt.Printf("Display Monitor Description: %s\n", monitor.physicalInfo.Description)

	currentValue, maximumValue, err := displayController.GetVCPFeatureAndVCPFeatureReply(monitor.physicalInfo.Handle, vcpCode)
	if err != nil {
		return err
	} else {
		fmt.Printf("Display Monitor VCP Code 0x%x, Current Value %d, Maximum Value %d\n", vcpCode, currentValue, maximumValue)
		fmt.Println()
	}

	return nil
}
