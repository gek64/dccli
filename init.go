package main

import (
	"github.com/gek64/displayController"
	"log"
)

type monitor struct {
	id           int
	physicalInfo displayController.PhysicalMonitorInfo
	sysInfo      displayController.SystemMonitorInfo
}

var (
	monitors    []monitor
	VCPFeatures = map[string]byte{
		"Brightness":            displayController.Brightness,
		"Contrast":              displayController.Contrast,
		"Red":                   displayController.Red,
		"Green":                 displayController.Green,
		"Blue":                  displayController.Blue,
		"InputSource":           displayController.InputSource,
		"Volume":                displayController.Volume,
		"Sharpness":             displayController.Sharpness,
		"ColorSaturation":       displayController.ColorSaturation,
		"MuteORScreenBlank":     displayController.MuteORScreenBlank,
		"HorizontalFrequency":   displayController.HorizontalFrequency,
		"VerticalFrequency":     displayController.VerticalFrequency,
		"DisplayTechnologyType": displayController.DisplayTechnologyType,
		"DisplayUsageTime":      displayController.DisplayUsageTime,
		"PowerMode":             displayController.PowerMode,
	}
)

func init() {
	// 获取所有系统显示设备
	displayMonitorInfos, err := displayController.GetSystemMonitors()
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
