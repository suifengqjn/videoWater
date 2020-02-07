# 批量伪原创视频剪辑器

[快速入门视频教程](https://www.bilibili.com/video/av84085197/)

[针对小白超级详细的视频教程及文档说明](https://github.com/suifengqjn/videoWater/blob/master/SREADME.md)

## 其他相关软件

[视频采集器](https://github.com/suifengqjn/videoCollector)
[视频全自动发布器](https://github.com/suifengqjn/mediaBot)

流程：视频自动采集->伪原创剪辑->全自动发布： 全自动流水线生产，躺着赚钱
采集器和发布器点击上面链接，本页面软件是全自动视频剪辑软件


## 注意事项（请仔细查看）

许多用户反馈容易出错的地方

* 配置文件不要用记事本直接打开，请用sublime Text 打开修改 
* 配置文件中的引号(") 逗号(,) 都是英文输入法 

## sublime Text 下载

链接:https://pan.baidu.com/s/1b7NM83ii66IaeibfXlNATQ  密码:wfys


## 功能介绍

视频自媒体批量处理软件，秒杀市面上绝大部分收费软件。
针对视频自媒体运营者，包括原创和伪原创，这一款软件就足够了。

填好配置信息后，一键启动即可处理成千上万个视频。每次使用只需双击启动即可

下面的任意一个操作，都会修改视频的md5，也就是所谓的md5去重。
经过本软件处理后，极高概率通过各大自媒体平台的原创审核

* 已完成
* [x] [视频分段] 
* [x] [视频截图] 
* [x] [视频格式转换] 
* [x] [帧率设置] 
* [x] [比特率设置] 
* [x] [剪掉视频前几秒] 
* [x] [剪掉视频最后几秒] 
* [x] [视频剪裁]
* [x] [去除水印] 
* [x] [视频镜像] 
* [x] [调整视频分辨率] 
* [x] [视频压缩] 
* [x] [添加文字水印] 
* [x] [添加跑马灯文字水印] 
* [x] [添加图片水印]
* [x] [倍速播放(加速减速)]
* [x] [添加片头] 
* [x] [添加片尾] 

* TODO
* [ ] [1. 去掉视频原有的配音] 
* [ ] [2. 添加自己的配音] 
* [ ] [3. 添加字幕]
* [ ] [4. 随机添加背景音乐]
* [ ] [5. 多镜头合并]
* [ ] [6. 多种滤镜]


## 软件收费说明

* 基础套餐：1元，有效期一个月，每天20个视频
* 月卡套餐: 15元，有效期一个月，每天200个视频
* 半年卡套餐: 100元，有效期180天，每天1000个视频
* 年卡套餐: 180元，有效期360天，无限制

可以先买基础套餐试用效果

密钥购买地址：[https://www.kuaifaka.com/purchasing?link=3ZUpQ](https://www.kuaifaka.com/purchasing?link=3ZUpQ)

## 打包程序下载

支持 mac win32 win64 系统
链接: https://pan.baidu.com/s/1B2J-ggsmkhW_cLo604VKbA  密码:a8f7

## 程序使用

打开 config.toml 文件，修改配置，每一项都有一个开关，选择自己需要操作的选项。
将需要处理的视频全部放入video文件夹(支持多目录)，然后运行

下载的程序包中已填好样例配置，只需填入密钥即可双击启动，查看效果

解压文件夹后
**mac 运行**
cd mac  进入到mac 文件夹下
./vm

**windows 运行**
双击vm.exe 文件即可

文件夹中已经配置了示例配置和演示用的视频，直接运行即可查看效果

## 水印样式说明

目前支持9种样式

 > 1. 左上角 2. 右上角 3. 右下角 4. 左下角 5. 正中间 6. 上正中间 7. 右正中间 8. 下正中间 9. 坐正中间 
 
 >  sp1 水平方向距离边界的距离
 
 >  sp2 竖直方向距离边界的距离

以图片水印举例
```
[waterImage]
    switch = true
    path = "./source/item.jpg"
    style = 1
    sp1 = 50
    sp2 = 100
```

下面是5种样式的效果

文字水印样式和图片水印样式一直

#### style 1
<img src="https://github.com/suifengqjn/videoWater/blob/master/image/style1.jpg?raw=true" width="120" height="212" alt="style 1"/>

#### style 2
<img src="https://github.com/suifengqjn/videoWater/blob/master/image/style2.jpg?raw=true" width="120" height="212" alt="style 2"/>

#### style 3
<img src="https://github.com/suifengqjn/videoWater/blob/master/image/style3.jpg?raw=true" width="120" height="212" alt="style 3"/>

#### style 4
<img src="https://github.com/suifengqjn/videoWater/blob/master/image/style4.jpg?raw=true" width="120" height="212" alt="style 4"/>

#### style 5
<img src="https://github.com/suifengqjn/videoWater/blob/master/image/style5.jpg?raw=true" width="120" height="212" alt="style 5"/>


## 版本更新记录

#### 1.0
* 视频格式转换
* 帧率设置
* 比特率设置
* 剪掉视频前几秒
* 剪掉视频最后几秒
* 去除水印
* 调整视频分辨率
* 添加文字水印
* 添加图片水
* 添加片头
* 添加片尾

#### 1.1
修复去除水印边界问题

#### 1.2
* 增加视频剪裁功能，此功能也适用于去水印
* 可以手动指定配置文件,使用于有多套配置文件  

mac `./vm -f config1.toml`
windows `./vm.exe -f config1.toml`

#### 1.3
* 增加倍速播放，可以将视频加速或者减速

#### 1.4 
* 文字水印，图片水印增加4种样式，一共9种
* 增加文字水印跑马灯效果

#### 1.5 
* 视频压缩

#### 2.0
* 视频镜像
* 视频分段
* 视频截图
#### 2.1
* 新增一种视频剪裁方式
* 可设置两个水印位置
* 调整结果输出
* 添加定时任务(主要是用于配合自动发布)
## 其他说明

我们的目标是软件一开，实现睡后收入
有任何问题 微信资讯

![](https://github.com/suifengqjn/videoWater/blob/master/image/wechat.jpeg?raw=true)





 

