package main

import (
	"fmt"
	"github.com/gek64/displayController"
	"syscall"
)

func getDisplayByHandle(handle int, vcpCode byte) (err error) {
	var found bool = false

	if len(monitors) == 0 {
		return fmt.Errorf("can't find any physical display monitor")
	}

	for _, monitor := range monitors {
		if monitor.PhysicalInfo.Handle == syscall.Handle(handle) {
			found = true
			err := getVCPFeatureValue(monitor, vcpCode)
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

func getVCPFeatureValue(monitor displayController.CompositeMonitorInfo, vcpCode byte) (err error) {
	fmt.Printf("Display Monitor Handle: %d\n", monitor.PhysicalInfo.Handle)
	fmt.Printf("Display Monitor Description: %s\n", monitor.PhysicalInfo.Description)

	currentValue, maximumValue, err := displayController.GetVCPFeatureAndVCPFeatureReply(monitor.PhysicalInfo.Handle, vcpCode)
	if err != nil {
		return err
	}

	fmt.Printf("Display Monitor VCP Code 0x%x, Current Value %d, Maximum Value %d\n", vcpCode, currentValue, maximumValue)
	return nil
}
