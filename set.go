package main

import (
	"fmt"
	"github.com/gek64/displayController"
	"syscall"
)

func setDisplayByHandle(handle int, vcpCode byte, value int) (err error) {
	var found = false

	if len(monitors) == 0 {
		return fmt.Errorf("can't find any physical display monitor")
	}

	for _, monitor := range monitors {
		if monitor.PhysicalInfo.Handle == syscall.Handle(handle) {
			found = true
			err := setVCPFeatureValue(monitor, vcpCode, value)
			if err != nil {
				return err
			}
		}
	}

	if !found {
		return fmt.Errorf("can't find any physical display monitor with handle %d\n", handle)
	}
	return nil
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

func setVCPFeatureValue(monitor displayController.CompositeMonitorInfo, vcpCode byte, value int) (err error) {
	fmt.Printf("Display Monitor Handle: %d\n", monitor.PhysicalInfo.Handle)
	fmt.Printf("Display Monitor Description: %s\n", monitor.PhysicalInfo.Description)

	err = displayController.SetVCPFeature(monitor.PhysicalInfo.Handle, vcpCode, value)
	if err != nil {
		return err
	}

	fmt.Printf("Display Monitor VCP Code 0x%x, Set to Value %d\n", vcpCode, value)
	return nil
}
