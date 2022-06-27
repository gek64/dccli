package main

import (
	"fmt"
	"github.com/gek64/displayController"
	"syscall"
)

func showDisplayByHandle(handle int) (err error) {
	var found bool = false

	if len(monitors) == 0 {
		return fmt.Errorf("can't find any physical display monitor")
	}

	for _, monitor := range monitors {
		if monitor.PhysicalInfo.Handle == syscall.Handle(handle) {
			found = true
			showDisplay(monitor)
			break
		}
	}

	if !found {
		return fmt.Errorf("can't find any physical display monitor with handle %d\n", handle)
	}

	return nil
}

func showAllDisplay() (err error) {
	if len(monitors) == 0 {
		return fmt.Errorf("can't find any physical display monitor")
	} else {
		fmt.Printf("Find %d Display Monitor\n", len(monitors))
	}

	for _, monitor := range monitors {
		showDisplay(monitor)
	}
	return nil
}

func showDisplay(monitor displayController.CompositeMonitorInfo) {
	fmt.Printf("Display Monitor Handle: %d\n", monitor.PhysicalInfo.Handle)
	fmt.Printf("Display Monitor Description: %s\n", monitor.PhysicalInfo.Description)
	fmt.Printf("Display Monitor Resolution: %d x %d\n", monitor.SysInfo.RectAngle.Right-monitor.SysInfo.RectAngle.Left, monitor.SysInfo.RectAngle.Bottom-monitor.SysInfo.RectAngle.Top)
	fmt.Printf("Display Monitor Position: Bottom %d Top %d Right %d Left %d\n", monitor.SysInfo.RectAngle.Bottom, monitor.SysInfo.RectAngle.Top, monitor.SysInfo.RectAngle.Right, monitor.SysInfo.RectAngle.Left)
}
