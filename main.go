package main

import (
	"flag"
	"fmt"
	"github.com/gek64/displayController"
	"log"
	"os"
)

var (
	cliShowDisplay bool
	cliGetDisplay  bool
	cliSetDisplay  bool

	cliDisplayID     int
	cliDisplayHandle int
	cliAll           bool
	cliHelp          bool
	cliVersion       bool
)

func init() {
	// 对显示器执行的操作
	flag.BoolVar(&cliShowDisplay, "show", false, "show display infos")
	flag.BoolVar(&cliGetDisplay, "get", false, "get display data")
	flag.BoolVar(&cliSetDisplay, "set", false, "set display data")

	// 按特征值选择显示器
	flag.IntVar(&cliDisplayID, "id", -1, "Display ID")
	flag.IntVar(&cliDisplayHandle, "handle", -1, "Display Handle")
	flag.BoolVar(&cliAll, "all", false, "All Display")

	// cli基础
	flag.BoolVar(&cliHelp, "h", false, "show help")
	flag.BoolVar(&cliVersion, "v", false, "show version")
	flag.Parse()

	// 重写显示用法函数
	flag.Usage = func() {
		var helpInfo = `Usage:
    dccli [Command] [Arguments]

Command:
	-show          : show
	-get        : get
	-set           : set
	
	-h                : show help
	-v                : show version

Arguments:
	-id      <int_number>      : by id
	-handle <int_number>  : by handle
    -all   : all

Example:
1) 
2) 
3) 
4) 
5) `
		fmt.Println(helpInfo)
	}

	// 如果无 args 或者 指定 h 参数,打印用法后退出
	if len(os.Args) == 1 || cliHelp {
		flag.Usage()
		os.Exit(0)
	}

	// 打印版本信息
	if cliVersion {
		fmt.Println("v1.00")
		os.Exit(0)
	}
}

func showChangelog() {
	var versionInfo = `Changelog:
  1.00:
    - First release`
	fmt.Println(versionInfo)
}

func main() {
	if cliShowDisplay {
		if cliDisplayID != -1 {
			err := showDisplayByID(cliDisplayID)
			if err != nil {
				log.Fatalln(err)
			}
		} else if cliDisplayHandle != -1 {
			err := showDisplayByHandle(cliDisplayHandle)
			if err != nil {
				log.Fatalln(err)
			}
		} else {
			err := showAllDisplay()
			if err != nil {
				log.Fatalln(err)
			}
		}

	}

	if cliGetDisplay {
		err := getAllDisplay(displayController.Contrast)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
