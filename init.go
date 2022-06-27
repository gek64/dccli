package main

import (
	"github.com/gek64/displayController"
	"log"
)

var (
	monitors []displayController.CompositeMonitorInfo
)

func init() {
	var err error
	monitors, err = displayController.GetCompositeMonitors()
	if err != nil {
		log.Fatalln(err)
	}
}
