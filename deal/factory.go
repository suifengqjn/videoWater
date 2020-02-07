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

/*

如果有多个文件，则最终所有视频都合并到一起
1. 不剪辑，视频和文件都移入result中，视频和信息都在video 下
2. 剪辑，视频和文件都移入result中，视频在video/result 中,信息在video 下

VideoPath 1种是相对路径，一种是绝对路径
*/

func doEdit(con *common.Config) int {

	result := createResultDir(con.VideoPath)


	files, err := file.GetAllFiles(con.VideoPath)
	if err != nil || len(files) == 0 {
		fmt.Printf("当前目录：%v 没有文件", con.VideoPath)
		time.Sleep(time.Second * 5)
	}
	count := 0
	_, oriDirs, _ := file.GetCurrentFilesAndDirs(con.VideoPath)
	var delDirs []string

	if !Contains(oriDirs,result) {
		oriDirs = append(oriDirs, result)
	}
	for _, f := range files {

		temp := f
		if ffmpeg.IsVideo(f) == false {
			continue
		}

		if account.VDAccount.CanUse() == false {
			fmt.Println("今日免费次数已用完")
			time.Sleep(time.Second * 30)
			break
		}

		fmt.Println("正在处理", f)

		f = deal(f,result, con)

		if account.VDAccount.AccType < account.AccTypeYear {
			account.VDAccount.AddAction()
		}

		// 把最终的视频移入result 中
		if temp != f {
			to := result + "/" + filepath.Base(f)
			_ = file.MoveFile(f, to)

		}
		//临时生成的目录
		/*
		./video/2/123_2.mp4
		./video/2/cut_front/add_head/123_2.mp4
		*/
		tempdir := filepath.Dir(temp)
		newDir := filepath.Dir(f)
		str := strings.TrimPrefix(newDir,tempdir)
		str =strings.TrimPrefix(str,"/")
		arr := strings.Split(str,"/")
		if len(str) > 0 && len(arr) > 0 {
			str = arr[0]
			delDir := ""
			if strings.HasPrefix(f,"./") {
				delDir = "./" + tempdir + "/" + str
			} else {
				delDir = tempdir + "/" + str
			}
			if !Contains(delDirs, delDir) {
				delDirs = append(delDirs,delDir)
			}

		}


		//clean
		count ++

	}

	//_, dirs, _ := file.GetCurrentFilesAndDirs(con.VideoPath)
	//delDirs = append(delDirs, dirs...)
	//删除临时目录
	for _, d := range delDirs {
		if !Contains(oriDirs, d) {
			_ = os.RemoveAll(d)
		}
	}


	return count
}


func deal(f ,resultDir string, con *common.Config)string  {

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
		if con.Crop1.Duration < 0 {
			con.Crop1.Duration = int64(info.Duration) - con.Crop1.Start + con.Crop1.Duration
		}
		f = info.CropVideoWithSpan(fCmd, f, con.Crop1.Start,con.Crop1.Duration, con.Crop1.Left,con.Crop1.Right,con.Crop1.Top,con.Crop1.Bottom)
	}

	// 7. clear water
	if con.ClearWater.Switch {
		f = ffmpeg.ClearWater(fCmd,f,con.ClearWater.X,con.ClearWater.Y,con.ClearWater.W,con.ClearWater.H)
	}

	if con.ClearWater1.Switch {
		f = ffmpeg.ClearWater(fCmd,f,con.ClearWater1.X,con.ClearWater1.Y,con.ClearWater1.W,con.ClearWater1.H)
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
	if con.Task > 0 && len(con.Output) > 0 {
		if f != temp {
			_ = os.RemoveAll(temp)
		} else {

			//将原始视频移到result 中
			fileName := filepath.Base(temp)

			dest := resultDir + "/" + fileName
			_ = file.MoveFile(temp,dest)

		}

		fileName := filepath.Base(temp)
		dir := filepath.Dir(temp)
		preFile := strings.Split(fileName,".")[0]

		txtPath := dir + "/" + preFile + ".txt"

		if file.PathExist(txtPath) {
			dest := resultDir + "/"  + preFile + ".txt"
			_ = file.MoveFile(txtPath,dest)
		}
	}

	return f
}

func move(from ,to string)  {
	_ = file.MoveDirFiles(from,to)
}
