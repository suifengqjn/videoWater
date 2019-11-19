package factory

import (
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
	} else {
		return "./source/win/ffmpeg"
	}

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