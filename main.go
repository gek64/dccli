package main

import (
	"flag"
	"fmt"
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

	cliVCPCode    int
	cliVCPFeature string
	cliVCPValue   int

	cliHelp    bool
	cliVersion bool
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

	// VCP Code 相关
	flag.IntVar(&cliVCPCode, "vcp", -1, "VCP Code")
	flag.StringVar(&cliVCPFeature, "feature", "", "VCP Feature")
	flag.IntVar(&cliVCPValue, "value", -1, "VCP Value")

	// cli基础
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
    -h             : show help
    -v             : show version

Arguments:
    -id      <int_number>         : select display monitor by id
    -handle  <int_number>         : select display monitor by handle
    -all                          : select all display monitor
    -vcp     <vcp_code>           : specify vcp code
    -feature <vcp_feature>        : specify vcp feature
    -value   <vcp_feature_value>  : specify vcp feature value

Example:
1) dccli -show
2) dccli -show -all
3) dccli -get -id 0 -feature Brightness
4) dccli -set -handle 0 -vcp 0x10 -value 50`
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

	// get功能
	if cliGetDisplay {
		// vcp 代码或者功能需要指定一个
		if cliVCPCode == -1 && cliVCPFeature == "" {
			log.Fatalln("no specify vcp code or vcp feature")
		}
	}

	// set功能
	if cliSetDisplay {
		// vcp 代码或者功能需要指定一个
		if cliVCPCode == -1 && cliVCPFeature == "" {
			log.Fatalln("no specify vcp code or vcp feature")
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
		if cliDisplayID != -1 {
			if VCPFeatures[cliVCPFeature] != 0 {
				err := getDisplayByID(cliDisplayID, VCPFeatures[cliVCPFeature])
				if err != nil {
					log.Fatalln(err)
				}
			} else if cliVCPCode != -1 {
				err := getDisplayByID(cliDisplayID, byte(cliVCPCode))
				if err != nil {
					log.Fatalln(err)
				}
			}
		} else if cliDisplayHandle != -1 {
			if VCPFeatures[cliVCPFeature] != 0 {
				err := getDisplayByHandle(cliDisplayHandle, VCPFeatures[cliVCPFeature])
				if err != nil {
					log.Fatalln(err)
				}
			} else if cliVCPCode != -1 {
				err := getDisplayByHandle(cliDisplayHandle, byte(cliVCPCode))
				if err != nil {
					log.Fatalln(err)
				}
			}
		} else {
			if VCPFeatures[cliVCPFeature] != 0 {
				err := getAllDisplay(VCPFeatures[cliVCPFeature])
				if err != nil {
					log.Fatalln(err)
				}
			} else if cliVCPCode != -1 {
				err := getAllDisplay(byte(cliVCPCode))
				if err != nil {
					log.Fatalln(err)
				}
			}
		}
	}

	if cliSetDisplay {
		if cliDisplayID != -1 {
			if VCPFeatures[cliVCPFeature] != 0 {
				err := setDisplayByID(cliDisplayID, VCPFeatures[cliVCPFeature], cliVCPValue)
				if err != nil {
					log.Fatalln(err)
				}
			} else if cliVCPCode != -1 {
				err := setDisplayByID(cliDisplayID, byte(cliVCPCode), cliVCPValue)
				if err != nil {
					log.Fatalln(err)
				}
			}
		} else if cliDisplayHandle != -1 {
			if VCPFeatures[cliVCPFeature] != 0 {
				err := setDisplayByHandle(cliDisplayHandle, VCPFeatures[cliVCPFeature], cliVCPValue)
				if err != nil {
					log.Fatalln(err)
				}
			} else if cliVCPCode != -1 {
				err := setDisplayByHandle(cliDisplayHandle, byte(cliVCPCode), cliVCPValue)
				if err != nil {
					log.Fatalln(err)
				}
			}
		} else {
			if VCPFeatures[cliVCPFeature] != 0 {
				err := setAllDisplay(VCPFeatures[cliVCPFeature], cliVCPValue)
				if err != nil {
					log.Fatalln(err)
				}
			} else if cliVCPCode != -1 {
				err := setAllDisplay(byte(cliVCPCode), cliVCPValue)
				if err != nil {
					log.Fatalln(err)
				}
			}
		}
	}
}
