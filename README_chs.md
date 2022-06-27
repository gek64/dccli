```
██████╗  ██████╗ ██████╗██╗     ██╗
██╔══██╗██╔════╝██╔════╝██║     ██║
██║  ██║██║     ██║     ██║     ██║
██║  ██║██║     ██║     ██║     ██║
██████╔╝╚██████╗╚██████╗███████╗██║
╚═════╝  ╚═════╝ ╚═════╝╚══════╝╚═╝
```

- 显示控制器命令行界面

## 使用说明
- 不带参数直接运行，将启动demo控制亮度
```
使用说明:
    dccli [指令] [参数]
                                                                       
指令:
    -show          : 显示显示器信息
    -get           : 为选中的显示器获取vcp功能对应的值
    -set           : 为选中的显示器设置vcp功能对应的值
    -demo          : 简单控制亮度demo
    -h             : 显示帮助
    -v             : 显示版本

参数:
    -handle  <int_number>         : 通过handle选择显示器
    -all                          : 选择全部显示器
    -vcp     <vcp_code>           : 指定vcp代码
    -value   <vcp_feature_value>  : 指定vcp功能的值

例子:
1) dccli -demo
2) dccli -show
3) dccli -get -handle 0 -vcp 0x10
4) dccli -set -handle 0 -vcp 0x10 -value 50
```


## 安装
- 从发行下载 https://github.com/gek64/dccli/releases

## 编译
```sh
# 下载源代码的依赖
git clone https://github.com/gek64/displayController.git
# 下载源代码
git clone https://github.com/gek64/dccli.git
# 编译源代码
cd dccli
go build -v -trimpath -ldflags "-s -w"
```

## 常见问题
### 这个模块支持哪些系统?
- 这个模块当前只支持 windows, 未来会考虑支持macOS、linux内核系统、freebsd等系统

### 可以正常获取显示驱动程序参数，但是无法获得和控制显示显示监视器参数。
- 本程序使用`vesa`在1998年定义的`DDC/CI`显示器通讯标准协议与显示器进行数据交换，绝大部分的现代显示器都默认支持并启用了这一项功能，但部分显示器的制造商可能因为某些特定因素的考量而默认关闭了这个选项，请确认显示器`OSD`菜单中是否已经开启了`DDC/CI`功能选项，或与您的显示器制造商联系获取更多有关的信息

### 能自定义查询、设置的参数除了已在库文件中定义的还有哪些？
- 请参考以下两篇文章来获取更多的自定义选项
- https://www.ddcutil.com/vcpinfo_output/
- https://www.hattelandtechnology.com/hubfs/pdf/misc/doc101681-1_8_and_13inch_dis_ddc_control.pdf

### 如何查找我自己的显示监视器支持的参数？
- 如果监视器不支持某个参数，则在调用命令时将返回错误，您可以使用错误信息来确定监视器是否支持某个参数
- 可以使用这个工具来检查你的显示支持的参数[ControlMyMonitor](https://www.nirsoft.net/utils/control_my_monitor.html)

## 许可证
- **GPL-3.0 License**
- 查看 `LICENSE` 获取详细内容

## 致谢
- [goland](https://www.jetbrains.com/go/)
- [vscode](https://code.visualstudio.com/)
