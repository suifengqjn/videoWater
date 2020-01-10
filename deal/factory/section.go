package factory

import (
	"fmt"
	"myProject/videoWater/account"
	"myProject/videoWater/deal/config"
	"myTool/ffmpeg"
	"myTool/file"
	"time"
)

func DoSection(con *config.Config)  {
	if len(con.SectionPath) == 0 {
		return
	}

	files, err := file.GetAllFiles(con.VideoPath)
	if err != nil {
		return
	}

	if con.CutSection.Switch == false {
		return
	}

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
		if account.VDAccount.AccType == account.AccTypeFree {
			account.VDAccount.AddAction()
		}

	}


}
