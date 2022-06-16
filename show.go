package main

import (
	"fmt"
	"syscall"
)

func showDisplayByID(id int) (err error) {
	// 找到了标志物
	var found = false

	if len(monitors) == 0 {
		return fmt.Errorf("can't find any physical display monitor")
	}

	for _, monitor := range monitors {
		if monitor.id == id {
			found = true
			fmt.Printf("Display Monitor ID: %d\n", monitor.id)
			fmt.Printf("Display Monitor Handle ID: %d\n", monitor.physicalInfo.Handle)
			fmt.Printf("Display Monitor Description: %s\n", monitor.physicalInfo.Description)
			fmt.Printf("Display Monitor Resolution: %d x %d\n", monitor.sysInfo.RectAngle.Right-monitor.sysInfo.RectAngle.Left, monitor.sysInfo.RectAngle.Bottom-monitor.sysInfo.RectAngle.Top)
			fmt.Printf("Display Monitor Position: Bottom %d Top %d Right %d Left %d\n", monitor.sysInfo.RectAngle.Bottom, monitor.sysInfo.RectAngle.Top, monitor.sysInfo.RectAngle.Right, monitor.sysInfo.RectAngle.Left)
			fmt.Println()
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

func showDisplayByHandle(handle int) (err error) {
	// 找到了标志物
	var found = false

	if len(monitors) == 0 {
		return fmt.Errorf("can't find any physical display monitor")
	}

	for _, monitor := range monitors {
		if monitor.physicalInfo.Handle == syscall.Handle(handle) {
			found = true
			fmt.Printf("Display Monitor ID: %d\n", monitor.id)
			fmt.Printf("Display Monitor Handle ID: %d\n", monitor.physicalInfo.Handle)
			fmt.Printf("Display Monitor Description: %s\n", monitor.physicalInfo.Description)
			fmt.Printf("Display Monitor Resolution: %d x %d\n", monitor.sysInfo.RectAngle.Right-monitor.sysInfo.RectAngle.Left, monitor.sysInfo.RectAngle.Bottom-monitor.sysInfo.RectAngle.Top)
			fmt.Printf("Display Monitor Position: Bottom %d Top %d Right %d Left %d\n", monitor.sysInfo.RectAngle.Bottom, monitor.sysInfo.RectAngle.Top, monitor.sysInfo.RectAngle.Right, monitor.sysInfo.RectAngle.Left)
			fmt.Println()
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

func showAllDisplay() (err error) {
	if len(monitors) == 0 {
		return fmt.Errorf("can't find any physical display monitor")
	} else {
		fmt.Printf("Find %d Display Monitor\n", len(monitors))
	}

	for _, monitor := range monitors {
		fmt.Printf("Display Monitor ID: %d\n", monitor.id)
		fmt.Printf("Display Monitor Handle ID: %d\n", monitor.physicalInfo.Handle)
		fmt.Printf("Display Monitor Description: %s\n", monitor.physicalInfo.Description)
		fmt.Printf("Display Monitor Resolution: %d x %d\n", monitor.sysInfo.RectAngle.Right-monitor.sysInfo.RectAngle.Left, monitor.sysInfo.RectAngle.Bottom-monitor.sysInfo.RectAngle.Top)
		fmt.Printf("Display Monitor Position: Bottom %d Top %d Right %d Left %d\n", monitor.sysInfo.RectAngle.Bottom, monitor.sysInfo.RectAngle.Top, monitor.sysInfo.RectAngle.Right, monitor.sysInfo.RectAngle.Left)
		fmt.Println()
	}
	return nil
}
