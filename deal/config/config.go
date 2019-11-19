package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"path/filepath"
)

var VideoWaterCon *Config
type Config struct {
	VideoPath string `toml:"videoPath"`
	System int
	Format Format
	FrameRate FrameRate
	BitRate BitRate
	CutFront CutFront	`toml:"cutFront"`
	CutBack CutBack	`toml:"cutBack"`
	ClearWater ClearWater
	Resolution Resolution
	WaterText WaterText
	WaterImage WaterImage
	FilmTitle FilmTitle
	FilmEnd FilmEnd


}

type Format struct {
	Switch bool
	Form string
}

type FrameRate struct {
	Switch bool
	Value string
}

type BitRate struct {
	Switch bool
	Value string
}

type CutFront struct {
	Switch bool
	Value int
}

type CutBack struct {
	Switch bool
	Value int
}

type ClearWater struct {
	Switch bool
	X int
	Y int
	W int
	H int
}

type Resolution struct {
	Switch bool
	W int
	H int
}

type WaterText struct {
	Switch bool
	Content string
	Path string
	Size int
	Color string
	Alpha float32
	Style int
	Sp1 int
	Sp2 int
}

type WaterImage struct {
	Switch bool
	Path string
	Style int
	Sp1 int
	Sp2 int
}


type FilmTitle struct {
	Switch bool
	Path   string
}

type FilmEnd struct {
	Switch bool
	Path   string
}







func init() {
	ReadConfig()
}

func ReadConfig() *Config  {
	if VideoWaterCon != nil {
		return VideoWaterCon
	}

	cur, _ := os.Getwd()
	fmt.Println("工程路径：", cur)

	con_path := filepath.Join(cur, "config.toml")

	_, err := toml.DecodeFile(con_path, &VideoWaterCon)
	if err != nil {
		panic(err)
	}

	fmt.Println("项目配置", *VideoWaterCon)

	return VideoWaterCon

}