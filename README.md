# videoWater

由于常常在网上搜集视频，然后将视频进行一系列的处理，然后发布到头条等自媒体，
视频处理往往要花大量的时间，现在写了个程序，可以将视频进行批量处理。

下面的任意一个操作，都会修改视频的md5，也就是所谓的md5去重。
经过下面的2，3，4，5，6，8，9，10，11之后，基本上可以算是自己的视频了

* 已完成
* [x] [1. 视频格式转换] 
* [x] [2. 帧率设置] 
* [x] [3. 比特率设置] 
* [x] [4. 减掉视频前几秒] 
* [x] [5. 减掉视频最后几秒] 
* [x] [6. 去除水印] 
* [x] [7. 调整视频分辨率] 
* [x] [8. 添加文字水印] 
* [x] [9. 添加图片水印]
* [x] [10. 添加片头] 
* [x] [11. 添加片尾] 

* TODO
* [ ] [1. 去掉视频原有的配音] 
* [ ] [2. 添加自己的配音] 
* [ ] [3. 跑马灯水印]
* [ ] [4. 添加字幕]
* [ ] [5. 加速减速]

## 程序使用

打开 config.toml 文件，修改配置，每一项都有一个开关，选择自己需要操作的选项。
将需要处理的视频全部放入video文件夹，然后运行

`./vm run`


## 水印样式说明

目前支持5种样式

** 1. 左上角 2.右上角 3.右下角 4.左下角 5.正中间 **
** sp1 水平方向距离边界的距离 **
** sp2 竖直方向距离边界的距离 **

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
![](https://github.com/suifengqjn/videoWater/blob/master/image/style1.jpg?raw=true)

#### style 2
![](https://github.com/suifengqjn/videoWater/blob/master/image/style2.jpg?raw=true)

#### style 3
![](https://github.com/suifengqjn/videoWater/blob/master/image/style3.jpg?raw=true)

#### style 4
![](https://github.com/suifengqjn/videoWater/blob/master/image/style4.jpg?raw=true)

#### style 5
![](https://github.com/suifengqjn/videoWater/blob/master/image/style5.jpg?raw=true)



 

