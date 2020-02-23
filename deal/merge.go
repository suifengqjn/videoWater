package deal

import (
	"fmt"
	"log"
	"myProject/videoWater/common"
	cm "myTool/common"
	"myTool/ffmpeg"
	"myTool/file"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func DoMerge(con *common.Config) int  {
	if con.MergeSection.Switch == 0 {
		return 0
	}

	files, dirs, err := file.GetCurrentFilesAndDirs(con.MergeSection.MergePath)
	if err != nil {
		return 0
	}

	fmt.Println("进行视频合成处理")
	// 先处理当前的文件
	count := 0
	if len(files) >= con.MergeSection.Count {
		count ++
		MergeRandom(files, con.MergeSection.Count)
	}


	// 处理文件夹
	for _, d := range dirs {

		// 获取子文件夹中的文件
		files, err := file.GetCurrentFiles(d)
		if err != nil || len(files) < con.MergeSection.Count {
			continue
		}

		MergeRandom(files, con.MergeSection.Count)
	}


	return count
}

func MergeRandom(videos []string, count int)  {
	if len(videos) == 0 {
		return
	}
	if len(videos) < count {
		log.Printf("当前文件夹视频数量少于指定数量：%v \n",filepath.Base(videos[0]))
	}

	// 找出随机的视频
	m := make(map[string]string)
	for _, v := range videos {
		n := time.Now().String()
		m[cm.MD5String(n)] = v
	}
	c := 0
	var res []string
	for _, v := range m {
		if count == c {
			 break
		}
		res = append(res, v)
		c++
	}

	// ./merge/111/1.mp4 -> ./merge/111/result/output.mp4
	fileName := filepath.Base(videos[0])
	suf := strings.Split(fileName, ".")[1]

	curDir := filepath.Dir(videos[0])

	outDir := curDir + "/" + "result"
	os.MkdirAll(outDir, os.ModePerm)


	output := fmt.Sprintf("%v/%v.%v", outDir, "output", suf)
	ffmpeg.MergeMultiVideoByResolution(fCmd,res,output,0,0)


}
