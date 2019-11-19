package main

import (
	"fmt"
	"myProject/videoWater/deal/config"
	"myProject/videoWater/deal/factory"
)

func main() {
	Run()
}

func Run()  {
	con := config.ReadConfig()
	fmt.Println(con)

	factory.DoFactory(con)

}