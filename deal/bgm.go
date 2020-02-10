package deal

import (
	"fmt"
	"math/rand"
	"myTool/ffmpeg"
)

func addBgm(bgmPath ,videoPath string, cover bool)string  {

	bgmFiles := GetAllBgm(bgmPath)
	if len(bgmFiles) == 0 {
		fmt.Println("背景音乐为空")
		return videoPath
	}

	index := rand.Int() % len(bgmFiles)

	bgm :=  bgmFiles[index]

	info, err := ffmpeg.GetVideoInfo(fCmd,videoPath)
	if err != nil {
		return videoPath
	}

	return info.AddBgm(fCmd,videoPath,bgm,cover)

}