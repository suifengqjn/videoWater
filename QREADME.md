
### 问题总览

以下是部分用户遇到的一些问题，这里写出来并提提供解决方案，如遇到问题，可先在本页面查看是否已存在。

也页面会持续更新

## 部分用户遇到的问题
#### 1. 软件莫名退出或者失效

* 检查是否被杀毒软件干掉

#### 2. 有些视频能处理，有些无法处理

* 检查视频格式是否是以下几种 "mp4","m4v","avi","flv","webm","mpeg","divx","mov","asf","wmf","rm","rmvb","ts"
* 修改文件名，改成数字或字母的形式，如 123.xxx, abc.xxx 12ab.xxx 等，部分电脑无法识别中文及一些特俗字符的文件

#### 3. 添加片头之后没有声音
* 检查片头是否有音轨，如果用样片中的片头没问题，那么可能是你的片头存在问题


## 大部分用户问过的问题

#### 1. 关于处理速度
处理速度主要有 视频数量的多少，视频的大小，时长，以及操作有关。不同的操作耗时不一样，比如修改帧率，码率，添加水印等是比较耗时的操作。
软件目前支持并发操作，并发数量根据电脑cpu核数相关，比如你的电脑是8核，那么可以同时对8个视频进行处理，相当于处理速度提高8倍。

#### 2. 关于功能
已经具有的功能都在软件介绍页面列出来了，没有列出来的就是没有，如果你需要添加什么额外的功能，可以微信发我，看情况添加。


#### 3. 密钥购买
打开软件，在介绍页面的下面就有连接
[https://www.kuaifaka.com/purchasing?link=3ZUpQ](https://www.kuaifaka.com/purchasing?link=3ZUpQ)

#### 4.关于出错
软件会打印出一些内容，方便查看，但那些东西不是出错了，我给出一个出错的例子给大家看

一般出错了就这下面两种样子
```
[Debug log] [[/mylog/log.go:79: myTool/mylog.(*DebugLogger).LogError] [转换格式失败 exit status 1]]
```

或者是这个样子
```

ffmpeg version git-2019-11-09-0f89a22 Copyright (c) 2000-2019 the FFmpeg developers
  built with Apple clang version 11.0.0 (clang-1100.0.33.8)
  configuration: --enable-gpl --enable-version3 --enable-sdl2 --enable-fontconfig --enable-gnutls --enable-iconv --enable-libass --enable-libdav1d --enable-libbluray --enable-libfreetype --enable-libmp3lame --enable-libopencore-amrnb --enable-libopencore-amrwb --enable-libopenjpeg --enable-libopus --enable-libshine --enable-libsnappy --enable-libsoxr --enable-libtheora --enable-libtwolame --enable-libvpx --enable-libwavpack --enable-libwebp --enable-libx264 --enable-libx265 --enable-libxml2 --enable-libzimg --enable-lzma --enable-zlib --enable-gmp --enable-libvidstab --enable-libvorbis --enable-libvo-amrwbenc --enable-libmysofa --enable-libspeex --enable-libxvid --enable-libaom --enable-appkit --enable-avfoundation --enable-coreimage --enable-audiotoolbox
  libavutil      56. 35.101 / 56. 35.101
  libavcodec     58. 60.100 / 58. 60.100
  libavformat    58. 34.101 / 58. 34.101
  libavdevice    58.  9.100 / 58.  9.100
  libavfilter     7. 66.100 /  7. 66.100
  libswscale      5.  6.100 /  5.  6.100
  libswresample   3.  6.100 /  3.  6.100
  libpostproc    55.  6.100 / 55.  6.100
Input #0, mov,mp4,m4a,3gp,3g2,mj2, from '/Users/qjn/Desktop/mac/video/123.mp4':
  Metadata:
    major_brand     : isom
    minor_version   : 512
    compatible_brands: isomiso2avc1mp41
    encoder         : Lavf58.10.100
  Duration: 00:00:24.64, start: 0.000000, bitrate: 904 kb/s
    Stream #0:0(und): Video: h264 (High) (avc1 / 0x31637661), yuv420p, 960x412, 829 kb/s, 30 fps, 30 tbr, 15360 tbn, 60 tbc (default)
    Metadata:
      handler_name    : VideoHandler
    Stream #0:1(und): Audio: aac (LC) (mp4a / 0x6134706D), 44100 Hz, stereo, fltp, 64 kb/s (default)
    Metadata:
      handler_name    : SoundHandler
Stream mapping:
  Stream #0:0 -> #0:0 (h264 (native) -> h264 (libx264))
  Stream #0:1 -> #0:1 (aac (native) -> aac (native))
Press [q] to stop, [?] for help
[Parsed_drawtext_0 @ 0x7fb1a6d2ef00] Could not set tabsize.
[AVFilterGraph @ 0x7fb1a6d2afc0] Error initializing filter 'drawtext' with args 'fontsize=1000000000:fontcolor=black@1:fontfile=/Users/qjn/Desktop/mac/source/simsun.ttc:text=adfafadfaf:x=(100):y=(100)'
Error reinitializing filters!
Failed to inject frame into filter network: Invalid argument
Error while processing the decoded data for stream #0:0
[aac @ 0x7fb1a784e000] Qavg: 65536.000
[aac @ 0x7fb1a784e000] 2 frames left in the queue on closing
Conversion failed!
]] 
```