package deal

import (
	"fmt"
	"myProject/videoWater/account"
	"myProject/videoWater/common"
	"myTool/ffmpeg"
	"myTool/file"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var fCmd = ""


func DoFactory(con *common.Config)  {

	fCmd = GetFCmd(con.System)

	c1 := DoSection(con)

	c2 := doEdit(con)

	t := c1 + c2
	if t == 0 {
		fmt.Println("没有视频需要处理")
	} else {
		fmt.Printf("一共处理视频 %v 个",t)
	}

	if con.Task > 0 && len(con.Output) > 0 {
		move(con.VideoPath + "/result",con.Output)
	}


}



func doEdit(con *common.Config) int {
	files, err := file.GetAllFiles(con.VideoPath)
	if err != nil || len(files) == 0 {
		fmt.Printf("当前目录：%v 没有文件", con.VideoPath)
		time.Sleep(time.Second * 5)
	}
	count := 0
	_, oriDirs, _ := file.GetCurrentFilesAndDirs(con.VideoPath)
	var delDirs []string
	for _, f := range files {

		if ffmpeg.IsVideo(f) == false {
			continue
		}

		if account.VDAccount.CanUse() == false {
			fmt.Println("今日免费次数已用完")
			time.Sleep(time.Second * 30)
			break
		}

		fmt.Println("正在处理", f)

		dir := filepath.Dir(f)
		result := filepath.Join(dir,"result")
		if !strings.HasPrefix(result,"/") && !strings.HasPrefix(result,"./")  {
			result = "./" + result
			oriDirs = append(oriDirs, result)
		}
		if file.PathExist(result) == false {
			_ = os.MkdirAll(result, os.ModePerm)
		}

		_, dirs, _ := file.GetCurrentFilesAndDirs(dir)
		for _, d := range dirs {
			if !Contains(oriDirs, d) {
				oriDirs = append(oriDirs, d)
			}
		}

		f = deal(f, con)

		if account.VDAccount.AccType < account.AccTypeYear {
			account.VDAccount.AddAction()
		}

		// 把最终的视频移入result 中
		to := result + "/" + filepath.Base(f)
		_ = file.MoveFile(f, to)

		//clean

		_, dirs, _ = file.GetCurrentFilesAndDirs(dir)
		for _, d := range dirs {
			if !Contains(delDirs, d) {
				delDirs = append(delDirs, d)
			}
		}
		count ++

	}

	_, dirs, _ := file.GetCurrentFilesAndDirs(con.VideoPath)
	delDirs = append(delDirs, dirs...)
	//删除临时目录
	for _, d := range delDirs {
		if !Contains(oriDirs, d) {
			_ = os.RemoveAll(d)
		}
	}


	return count
}


func deal(f string, con *common.Config)string  {

	temp := f
	// 0. snip
	if con.Snip.Switch {
		ffmpeg.Snip(fCmd, f, strconv.Itoa(con.Snip.T),strconv.Itoa(con.Snip.R))
	}

	//1 . 格式转换
	if con.Format.Switch {
		f = ffmpeg.CoverToCustomFormat(fCmd, f, con.Format.Form)
	}

	// 2. frame code
	if con.FrameRate.Switch {
		info, err := ffmpeg.GetVideoInfo(fCmd, f)
		if err != nil {
			return f
		}
		value := StringToValue(con.FrameRate.Value)
		if value <= 0 {
			return f
		}
		value = info.Zhen + value
		f = ffmpeg.UpdateFrameRate(fCmd,f,value)

	}
	// 3. bit code
	if con.BitRate.Switch {
		info, err := ffmpeg.GetVideoInfo(fCmd, f)
		if err != nil {
			return f
		}
		value := StringToValue(con.FrameRate.Value)
		if value <= 0 {
			return f
		}
		value = info.BitrateValue + value
		f = ffmpeg.UpdateBitRate(fCmd,f,value)

	}
	if con.CutFront.Switch && con.CutBack.Switch {
		info, err := ffmpeg.GetVideoInfo(fCmd, f)
		if err != nil {
			return f
		}
		f = info.CutFrontAndBack(fCmd,f,con.CutFront.Value,con.CutBack.Value)
	} else {
		//4. cut front
		if con.CutFront.Switch {
			info, err := ffmpeg.GetVideoInfo(fCmd, f)
			if err != nil {
				return f
			}
			f = info.CutFront(fCmd,f,con.CutFront.Value)
		}

		//5. cut back
		if con.CutBack.Switch {
			info, err := ffmpeg.GetVideoInfo(fCmd, f)
			if err != nil {
				return f
			}
			f = info.CutBack(fCmd,f,con.CutBack.Value)
		}
	}


	// 6. crop
	if con.Crop.Switch {
		f = ffmpeg.CropVideo(fCmd,f,con.Crop.Start,con.Crop.Duration,con.Crop.X,con.Crop.Y,con.Crop.W,con.Crop.H)
	} else if con.Crop1.Switch {
		info, err := ffmpeg.GetVideoInfo(fCmd, f)
		if err != nil {
			return f
		}
		f = info.CropVideoWithSpan(fCmd, f, con.Crop1.Start,con.Crop1.Duration, con.Crop1.Left,con.Crop1.Right,con.Crop1.Top,con.Crop1.Bottom)
	}

	// 7. clear water
	if con.ClearWater.Switch {
		f = ffmpeg.ClearWater(fCmd,f,con.ClearWater.X,con.ClearWater.Y,con.ClearWater.W,con.ClearWater.H)
	}

	//  mirror
	if con.Mirror.Switch {
		f = ffmpeg.Mirror(fCmd, f,con.Mirror.Direction)
	}

	//8. Resolution
	if con.Resolution.Switch {
		f = ffmpeg.UpdateResolution(fCmd, f, con.Resolution.W, con.Resolution.H)
	}

	if con.Compress.Switch {
		f = ffmpeg.Compress(fCmd, f, con.Compress.Preset, con.Compress.Crf)
	}
	//9. water text
	if con.WaterText.Switch {
		info, err := ffmpeg.GetVideoInfo(fCmd, f)
		if err != nil {
			return f
		}
		f = info.AddTextWaterWithStyle(
			fCmd,
			f,
			con.WaterText.Path,
			con.WaterText.Size,
			con.WaterText.Content,
			con.WaterText.Style,
			con.WaterText.Sp1,
			con.WaterText.Sp2,
			con.WaterText.Color,
			con.WaterText.Alpha,
		)
	}

	// 9.1 runtext
	if con.RunWaterText.Switch {
		info, err := ffmpeg.GetVideoInfo(fCmd, f)
		if err != nil {
			return f
		}
		f = info.AddScrollTextWater(
			fCmd,
			f,
			con.RunWaterText.Path,
			con.RunWaterText.Content,
			con.RunWaterText.Color,
			con.RunWaterText.Size,
			con.RunWaterText.IsTop,
			con.RunWaterText.LeftToRight,
			con.RunWaterText.Sp,
			)
	}
	//10. water image
	if con.WaterImage.Switch {
		info, err := ffmpeg.GetVideoInfo(fCmd, f)
		if err != nil {
			return f
		}

		f = info.AddTextWaterImageWithStyle(
			fCmd,
			f,
			con.WaterImage.Path,
			con.WaterImage.Style,
			con.WaterImage.Sp1,
			con.WaterImage.Sp2,
		)
	}

	//10. speed
	if con.Speed.Switch {
		v,err := strconv.ParseFloat(con.Speed.V, 10)
		if err == nil {
			f = ffmpeg.Speed(fCmd,f,float32(v))
		}

	}
	//11. film title
	if con.FilmTitle.Switch {
		info, err := ffmpeg.GetVideoInfo(fCmd, f)
		if err != nil {
			return f
		}

		newHeader := ffmpeg.UpdateResolution(fCmd, con.FilmTitle.Path, info.W, info.H)

		f = info.MergeVideoHeader(fCmd, newHeader,f)
	}
	//12. film end
	if con.FilmEnd.Switch {
		info, err := ffmpeg.GetVideoInfo(fCmd, f)
		if err != nil {
			return f
		}
		newFooter := ffmpeg.UpdateResolution(fCmd, con.FilmEnd.Path, info.W, info.H)
		f = info.MergeVideoFooter(fCmd, newFooter,f)
	}


	// 两种情况 1. 视频不做任何处理 视频和信息都在video 下
	// 2. 经过剪辑 视频在video/result 中,信息在video 下
	if len(con.Output) > 0 {
		if f != temp {
			_ = os.RemoveAll(temp)
		} else {

			//将原始视频移到result 中
			fileName := filepath.Base(temp)
			dir := filepath.Dir(temp)

			dest := dir + "/result/" + fileName
			_ = file.MoveFile(temp,dest)

		}

		fileName := filepath.Base(temp)
		dir := filepath.Dir(temp)
		preFile := strings.Split(fileName,".")[0]

		txtPath := dir + "/" + preFile + ".txt"

		if file.PathExist(txtPath) {
			dest := dir + "/result/"  + preFile + ".txt"
			_ = file.MoveFile(txtPath,dest)
		}
	}

	return f
}

func move(from ,to string)  {
	_ = file.MoveDirFiles(from,to)
}
