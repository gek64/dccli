package main

import (
	"github.com/gek64/displayController"
	"log"
)

type monitor struct {
	id           int
	physicalInfo displayController.PhysicalMonitorInfo
	sysInfo      displayController.DisplayMonitorInfo
}

var (
	monitors []monitor
)

func init() {
	// 获取所有系统显示设备
	displayMonitorInfos, err := displayController.GetAllMonitors()
	if err != nil {
		log.Fatalln(err)
	}

	for i, displayMonitorInfo := range displayMonitorInfos {
		// 获取物理显示器信息
		physicalMonitorInfo, err := displayController.GetPhysicalMonitor(displayMonitorInfo.Handle)
		if err != nil {
			continue
		}

		// 拼接到完整的显示器类中
		monitors = append(monitors, monitor{id: i, physicalInfo: physicalMonitorInfo, sysInfo: displayMonitorInfo})
	}
}
