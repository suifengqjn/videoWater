

### 视频演示

视频介绍了每种功能的参数改如何配置，及软件整体是如何使用

[1. 下载安装使用及简单演示(必看)](https://www.bilibili.com/video/av84085197/)

[2. 视频格式转换](https://www.bilibili.com/video/av84090158/)

[3. 改变视频帧率和码率](https://www.bilibili.com/video/av84090567/)

[4. 去除视频片头和片尾](https://www.bilibili.com/video/av84090675/)

[5. 视频剪裁](https://www.bilibili.com/video/av84090816/)

[6. 视频去水印](https://www.bilibili.com/video/av84093352/)

[7. 视频镜像制作](https://www.bilibili.com/video/av84093482/)

[8. 修改视频分辨率](https://www.bilibili.com/video/av84093628/)

[9. 视频压缩](https://www.bilibili.com/video/av84093725/)

[10. 文字水印和图片水印](https://www.bilibili.com/video/av84093826/)

[11. 倍速播放](https://www.bilibili.com/video/av84093943/)

[12. 添加片头和片尾](https://www.bilibili.com/video/av84094016/)

[13. 伪原创参数配置推荐](https://www.bilibili.com/video/av84094116/)

[20. 独立功能：视频分段](https://www.bilibili.com/video/av84094229/)



### windows 系统使用

这是软件解压后的结构

![](https://github.com/suifengqjn/videoWater/blob/master/image/r_1.png?raw=true)

* config.toml 需要对视频操作的配置
* source 软件依赖工具，不要动
* video 需要处理的视频放到这个文件夹
* vm.exe 启动程序

config.toml中针对每一项操作都有一个开关，不需要哪项操作，关闭即可

操作示例：

![](https://github.com/suifengqjn/videoWater/blob/master/image/r_2.png?raw=true)

比例现在在video 目录中放了一个视频，并且config.toml也已经配置好了
直接双击 vm.exe 打开程序。

![](https://github.com/suifengqjn/videoWater/blob/master/image/r_3.png?raw=true)

运行完毕后，video 下会多出一个result 文件夹，处理后的视频就在这里面


### mac 系统使用

mac 的打开方式不太一样

使用终端进入到程序所在文件夹

![](https://github.com/suifengqjn/videoWater/blob/master/image/r_4.png?raw=true)

运行程序
`./vm`

如果出现 permission denied 
则执行 `chmod 777 vm`

再执行 `./vm`