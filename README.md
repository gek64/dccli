```
██████╗  ██████╗ ██████╗██╗     ██╗
██╔══██╗██╔════╝██╔════╝██║     ██║
██║  ██║██║     ██║     ██║     ██║
██║  ██║██║     ██║     ██║     ██║
██████╔╝╚██████╗╚██████╗███████╗██║
╚═════╝  ╚═════╝ ╚═════╝╚══════╝╚═╝
```

[中文说明](#)
- Display controller command line interface

## Usage
```
Usage:                                                                 
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
4) dccli -set -handle 0 -vcp 0x10 -value 50
```


## Install
- Download From https://github.com/gek64/dccli/releases

## Compile
```sh
# Download the source code of dependent packages
git clone https://github.com/gek64/displayController.git
# Download application source code
git clone https://github.com/gek64/dccli.git
# Compile the source code
cd dccli
go build -v -trimpath -ldflags "-s -w"
```

## FAQ
### What operating system does this app support?
- It only supports Windows now, and support systems such as macOS, Linux kernel system and freeBSD will be considered in the future.

### Get the display driver parameter normally, but the display monitor parameter cannot be obtained and controlled.
- This program uses the `VESA` `DDC/CI` Display communication standard protocol which release in 1998 to exchange data with physical display monitors. Most of the modern display supports and enables this feature by default, If you encounter this problem, please confirm whether the `DDC/CI` function has been opened in OSD menu, or contact your display manufacturer to get more relevant information

### What other parameters can be customized?
- Please refer to the following articles to get more custom parameters
- https://www.ddcutil.com/vcpinfo_output/
- https://www.hattelandtechnology.com/hubfs/pdf/misc/doc101681-1_8_and_13inch_dis_ddc_control.pdf

### How to find parameters supported by my own display monitor？
- If the monitor does not support a certain parameter, the error will be returned when calling the command. You can use the error information to determine whether the monitor supports a certain parameter
- You can use this tool to check which parameters that your monitor supported [ControlMyMonitor](https://www.nirsoft.net/utils/control_my_monitor.html)

## License
- **GPL-3.0 License**
- See `LICENSE` for details

## Credits
- [goland](https://www.jetbrains.com/go/)
- [vscode](https://code.visualstudio.com/)
