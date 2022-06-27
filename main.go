package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	cliShowDisplay bool
	cliGetDisplay  bool
	cliSetDisplay  bool

	cliDisplayHandle int
	cliUseHandle     bool = false
	cliUseAll        bool = false

	cliVCPCode  int
	cliVCPValue int

	cliDemo    bool
	cliHelp    bool
	cliVersion bool
)

func init() {
	// 对显示器执行的操作
	flag.BoolVar(&cliShowDisplay, "show", false, "show display infos")
	flag.BoolVar(&cliGetDisplay, "get", false, "get display data")
	flag.BoolVar(&cliSetDisplay, "set", false, "set display data")

	// 按特征值选择显示器
	flag.IntVar(&cliDisplayHandle, "handle", -1, "Display Handle")
	flag.BoolVar(&cliUseAll, "all", false, "All Display(default)")

	// VCP Code 相关
	flag.IntVar(&cliVCPCode, "vcp", -1, "VCP Code")
	flag.IntVar(&cliVCPValue, "value", -1, "VCP Value")

	// cli基础
	flag.BoolVar(&cliDemo, "demo", false, "brightness setting demo")
	flag.BoolVar(&cliHelp, "h", false, "show help")
	flag.BoolVar(&cliVersion, "v", false, "show version")
	flag.Parse()

	// 重写显示用法函数
	flag.Usage = func() {
		var helpInfo = `Usage:
    dccli [Command] [Arguments]

Command:
    -show          : show display monitor info
    -get           : get vcp feature value from select display monitors
    -set           : set vcp feature value for select display monitors
    -demo          : simple demo to control brightness
    -h             : show help
    -v             : show version

Arguments:
    -handle  <int_number>         : select display monitor by handle
    -all                          : select all display monitor
    -vcp     <vcp_code>           : specify vcp code
    -value   <vcp_feature_value>  : specify vcp feature value

Example:
1) dccli -demo
2) dccli -show
3) dccli -get -handle 0 -vcp 0x10
4) dccli -set -handle 0 -vcp 0x10 -value 50`
		fmt.Println(helpInfo)
	}

	// demo控制亮度
	if len(os.Args) == 1 || cliDemo {
		err := runDemo()
		if err != nil {
			log.Println(err)
			time.Sleep(3 * time.Second)
		}
		os.Exit(0)
	}

	// help
	if cliHelp {
		flag.Usage()
		os.Exit(0)
	}

	// 打印版本信息
	if cliVersion {
		fmt.Println("v1.01")
		os.Exit(0)
	}

	// 显示器 handle 指定
	if cliDisplayHandle != -1 {
		cliUseHandle = true
	} else {
		cliUseAll = true
	}

	// get 功能检测
	if cliGetDisplay {
		// vcp 代码需要指定
		if cliVCPCode == -1 {
			log.Fatalln("no specify vcp code")
		}
	}

	// set 功能检测
	if cliSetDisplay {
		// vcp 代码需要指定
		if cliVCPCode == -1 {
			log.Fatalln("no specify vcp code")
		}
		// vcp 功能值需要指定
		if cliVCPValue == -1 {
			log.Fatalln("no specify vcp feature value")
		}
	}
}

func showChangelog() {
	var versionInfo = `Changelog:
  1.00:
    - First release
  1.0.1:
    - Compatible to displayController v1.01`
	fmt.Println(versionInfo)
}

func main() {
	if cliShowDisplay {
		if cliUseHandle {
			err := showDisplayByHandle(cliDisplayHandle)
			if err != nil {
				log.Fatalln(err)
			}
		} else if cliUseAll {
			err := showAllDisplay()
			if err != nil {
				log.Fatalln(err)
			}
		}
	}

	if cliGetDisplay {
		if cliUseHandle {
			err := getDisplayByHandle(cliDisplayHandle, byte(cliVCPCode))
			if err != nil {
				log.Fatalln(err)
			}
		} else if cliUseAll {
			err := getAllDisplay(byte(cliVCPCode))
			if err != nil {
				log.Fatalln(err)
			}
		}
	}

	if cliSetDisplay {
		if cliUseHandle {
			err := setDisplayByHandle(cliDisplayHandle, byte(cliVCPCode), cliVCPValue)
			if err != nil {
				log.Fatalln(err)
			}
		} else if cliUseAll {
			err := setAllDisplay(byte(cliVCPCode), cliVCPValue)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}
