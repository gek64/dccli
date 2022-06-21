package main

import (
	"fmt"
	"github.com/gek64/displayController"
)

func getAllDisplay(vcpCode byte) (err error) {
	if len(monitors) == 0 {
		return fmt.Errorf("can't find any physical display monitor")
	} else {
		fmt.Printf("Find %d Display Monitor\n", len(monitors))
	}

	for _, monitor := range monitors {
		fmt.Printf("Display Monitor ID: %d\n", monitor.id)
		fmt.Printf("Display Monitor Handle ID: %d\n", monitor.physicalInfo.Handle)
		fmt.Printf("Display Monitor Description: %s\n", monitor.physicalInfo.Description)

		currentValue, maximumValue, err := displayController.GetVCPFeatureAndVCPFeatureReply(monitor.physicalInfo.Handle, vcpCode)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Display Monitor VCP Code 0x%x, Current Value %d, Maximum Value %d\n", vcpCode, currentValue, maximumValue)

		fmt.Println()
	}
	return nil
}
