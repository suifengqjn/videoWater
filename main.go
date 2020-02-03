package main

import (
	"flag"
	"fmt"
	"log"
	"myProject/videoWater/account"
	"myProject/videoWater/common"
	"myProject/videoWater/deal"
	"strings"
	"time"
)
var conFile = flag.String("f", "./config.toml", "config file")
var videoPath = flag.String("v", "", "config file")
var line = "********************************************************************************"
func main() {
	Run()
}

func Run()  {

	// check config
	flag.Parse()
	con := common.ReadConfig(*conFile)
	if con == nil {
		log.Println("配置文件有误")
		time.Sleep(time.Second * 5)
		return
	}

	// check version
	code, msg := account.CheckVersion()
	if code != 1 {
		if len(msg) > 0 {
			fmt.Println(msg)
		} else {
			fmt.Println("请检查网络，稍后再试")
		}
		time.Sleep(time.Second * 100)
		return
	} else if len(msg) > 0 {
		fmt.Println(msg)
		time.Sleep(time.Second * 3)
	}

	if len(con.AppId) == 0 {
		fmt.Println("    请购买密钥   ")
		time.Sleep(time.Second * 500)
		return
	}

	acc := account.GetAccount(con.AppId)

	fmt.Println(line)
	fmt.Println(line)
	fmt.Println(line)
	fmt.Println()
	fmt.Println()
	printInfo()
	fmt.Println()
	if acc.AccType < 0 {
		fmt.Println(formatline("账户密钥:"+acc.AppId))
		fmt.Println(formatline("密钥无效，请购买密钥"))
		fmt.Println(formatline("密钥购买地址："+"https://www.kuaifaka.com/purchasing?link=3ZUpQ"))
	} else {
		fmt.Println(formatline(fmt.Sprintf("账户 密钥：%v",acc.AppId)))
		fmt.Println(formatline(fmt.Sprintf("账户类型：%v",acc.TYPE())))
		fmt.Println(formatline(fmt.Sprintf("%v",acc.Time)))
		fmt.Println(formatline(acc.Msg))
	}

	fmt.Println()
	fmt.Println(line)
	fmt.Println(line)
	fmt.Println(line)

	if acc.AccType < 0 {
		time.Sleep(time.Second * 500)
		return
	}
	time.Sleep(time.Second * 5)
	if len(*videoPath) > 0 {
		con.VideoPath = *videoPath
		fmt.Println(*videoPath)
	}
	deal.DoFactory(con)
	if con.Task > 0 {
		ticker := time.NewTicker(time.Minute * time.Duration(con.Task))

		for range ticker.C {
			deal.DoFactory(con)
		}
	}

	time.Sleep(time.Second * 5)

}

func formatline(text string)string  {

	r := strings.Repeat(" ", 10)
	return r + text + r

}

func printInfo()  {
	fmt.Println(formatline(fmt.Sprintf("伪原创视频批量剪辑器 %v", account.Version)))
	fmt.Println(formatline("软件地址：https://github.com/suifengqjn/videoWater"))
}
