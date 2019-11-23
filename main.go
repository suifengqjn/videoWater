package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"myProject/videoWater/deal/config"
	"myProject/videoWater/deal/factory"
	"net/http"
)

func main() {
	Run()
}

func Run()  {

	if !check() {
		return
	}

	con := config.ReadConfig()
	fmt.Println(con)

	factory.DoFactory(con)

}

type Data struct {
	Code int	`json:"code"`
	Msg string	`json:"msg"`
}

func check() bool  {

	url := "https://raw.githubusercontent.com/suifengqjn/videoWater/master/source/txt.json"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("请检查网络")
		return false
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("请检查网络")
		return false
	}

	var res Data
	err = json.Unmarshal(buf, &res)
	if err != nil {
		fmt.Println("请检查网络")
		return false
	}
	if len(res.Msg) > 0 {
		fmt.Println("===========================")
		fmt.Println(res.Msg)
		fmt.Println("===========================")
	}

	if res.Code == 1 {
		return true
	}

	return false
}