package main

import (
	"fmt"
	"github.com/gek64/displayController"
	"syscall"
)

func setDisplayByID(id int, vcpCode byte, value int) (err error) {
	// 找到了标志物
	var found = false

	if len(monitors) == 0 {
		return fmt.Errorf("can't find any physical display monitor")
	}

	for _, monitor := range monitors {
		if monitor.id == id {
			found = true
			err := setVCPFeatureValue(monitor, vcpCode, value)
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

func setDisplayByHandle(handle int, vcpCode byte, value int) (err error) {
	// 找到了标志物
	var found = false

	if len(monitors) == 0 {
		return fmt.Errorf("can't find any physical display monitor")
	}

	for _, monitor := range monitors {
		if monitor.physicalInfo.Handle == syscall.Handle(handle) {
			found = true
			err := setVCPFeatureValue(monitor, vcpCode, value)
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

func setAllDisplay(vcpCode byte, value int) (err error) {
	if len(monitors) == 0 {
		return fmt.Errorf("can't find any physical display monitor")
	}

	for _, monitor := range monitors {
		err := setVCPFeatureValue(monitor, vcpCode, value)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return nil
}

func setVCPFeatureValue(monitor monitor, vcpCode byte, value int) (err error) {
	fmt.Printf("Display Monitor ID: %d\n", monitor.id)
	fmt.Printf("Display Monitor Handle ID: %d\n", monitor.physicalInfo.Handle)
	fmt.Printf("Display Monitor Description: %s\n", monitor.physicalInfo.Description)

	err = displayController.SetVCPFeature(monitor.physicalInfo.Handle, vcpCode, value)
	if err != nil {
		return err
	} else {
		fmt.Printf("Display Monitor VCP Code 0x%x, Set to Value %d\n", vcpCode, value)
		fmt.Println()
	}

	return nil
}
