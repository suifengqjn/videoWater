package factory

import (
	"fmt"
	"myTool/sys"
	"strconv"
	"strings"
)

func GetFCmd(system int) string {

	if system == 0 {
		info := sys.GetSysInfo()
		system = info.PlatForm
	}
	if system == sys.MacOS {
		return "./source/mac/ffmpeg"
	} else if system == sys.Win64{
		return "./source/win/64/ffmpeg.exe"
	} else if system == sys.Win32 {
		return "./source/win/32/ffmpeg.exe"
	}

	fmt.Println("系统类型无法识别，请在配置中指定：1:mac 3: win32  4:win64")
	return ""

}

func StringToValue(str string)int  {
	v, err := strconv.Atoi(str)
	if err != nil {
		return v
	}

	if strings.HasPrefix(str,"+") {
		str = strings.TrimPrefix(str,"+")
		v, err := strconv.Atoi(str)
		if err != nil {
			return 0
		}
		return v
	}

	if strings.HasPrefix(str,"-") {
		str = strings.TrimPrefix(str,"-")
		v, err := strconv.Atoi(str)
		if err != nil {
			return 0
		}
		return v
	}
	return 0
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}