package deal

import (
	"fmt"
	"myProject/videoWater/account"
	"myProject/videoWater/common"
	"myTool/ffmpeg"
	"myTool/file"
	"time"
)

func DoSection(con *common.Config) int  {
	if con.CutSection.Switch == 0 {
		return 0
	}

	files, err := file.GetAllFiles(con.CutSection.SectionPath)
	if err != nil {
		return 0
	}

	count := 0
	fmt.Println("进行视频分段处理")
	for _, f := range files {

		if ffmpeg.IsVideo(f) == false {
			continue
		}

		if account.VDAccount.CanUse() == false {
			fmt.Println("今日免费次数已用完")
			time.Sleep(time.Second * 30)
			break
		}

		info, err := ffmpeg.GetVideoInfo(fCmd,f)
		if err != nil {
			continue
		}

		info.CutSection(fCmd, f, con.CutSection.Duration)
		if account.VDAccount.AccType < account.AccTypeYear {
			account.VDAccount.AddAction()
		}
		count++

	}

	return count
}
