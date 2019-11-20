package factory

import (
	"fmt"
	"myProject/videoWater/deal/config"
	"myTool/ffmpeg"
	"myTool/file"
	"os"
	"path/filepath"
	"strings"
)

var fCmd = ""

func DoFactory(con *config.Config)  {

	fCmd = GetFCmd(con.System)

	files, err := file.GetAllFiles(con.VideoPath)
	if err != nil || len(files) == 0 {
		fmt.Printf("当前目录：%v 没有文件", con.VideoPath)
	}

	_, oriDirs, _ := file.GetCurrentFilesAndDirs(con.VideoPath)
	var delDirs []string
	for _, f := range files {

		fmt.Println("当前处理", f)
		if ffmpeg.IsVideo(f) == false {
			continue
		}



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

		to := result + "/" + filepath.Base(f)

		fmt.Println("from",f, "to", to)
		_ = file.MoveFile(f, to)

		fmt.Println("处理结束", f)
		//clean

		_, dirs, _ = file.GetCurrentFilesAndDirs(dir)
		for _, d := range dirs {
			if !Contains(delDirs, d) {
				delDirs = append(delDirs, d)
			}
		}

	}

	_, dirs, _ := file.GetCurrentFilesAndDirs(con.VideoPath)
	delDirs = append(delDirs, dirs...)

	for _, d := range delDirs {
		if !Contains(oriDirs, d) {
			_ = os.RemoveAll(d)
		}
	}

}

func deal(f string, con *config.Config)string  {
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

	// 6. clear water
	if con.ClearWater.Switch {
		f = ffmpeg.ClearWater(fCmd,f,con.ClearWater.X,con.ClearWater.Y,con.ClearWater.W,con.ClearWater.H)
	}

	//7. Resolution
	if con.Resolution.Switch {
		f = ffmpeg.UpdateResolution(fCmd, f, con.Resolution.W, con.Resolution.H)
	}

	//8. water text
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

	//9. water image
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

	//10. film title
	if con.FilmTitle.Switch {
		info, err := ffmpeg.GetVideoInfo(fCmd, f)
		if err != nil {
			return f
		}

		newHeader := ffmpeg.UpdateResolution(fCmd, con.FilmTitle.Path, info.W, info.H)

		f = info.MergeVideoHeader(fCmd, newHeader,f)
	}
	//10. film end
	if con.FilmEnd.Switch {
		info, err := ffmpeg.GetVideoInfo(fCmd, f)
		if err != nil {
			return f
		}
		newFooter := ffmpeg.UpdateResolution(fCmd, con.FilmEnd.Path, info.W, info.H)
		f = info.MergeVideoFooter(fCmd, newFooter,f)
	}

	return f
}
