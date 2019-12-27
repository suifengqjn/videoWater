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
	Crop   Crop
	ClearWater ClearWater
	Resolution Resolution
	Compress Compress
	WaterText WaterText
	RunWaterText RunWaterText `toml:"RunWaterText"`
	WaterImage WaterImage
	Speed Speed
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

type Crop struct {
	Switch bool
	Start int64
	Duration int64
	X int
	Y int
	W int
	H int
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

type Compress struct {
	Switch bool
	Preset string
	Crf int
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

type RunWaterText struct {
	Switch  bool
	Content string
	Path    string
	Size    int
	Color   string
	IsTop   bool `toml:"isTop"`
	LeftToRight   bool `toml:"leftToRight"`
	Sp      int
}

type WaterImage struct {
	Switch bool
	Path string
	Style int
	Sp1 int
	Sp2 int
}

type Speed struct {
	Switch bool
	V string
}

type FilmTitle struct {
	Switch bool
	Path   string
}

type FilmEnd struct {
	Switch bool
	Path   string
}







func ReadConfig(file string) *Config  {
	if VideoWaterCon != nil {
		return VideoWaterCon
	}

	cur, _ := os.Getwd()
	fmt.Println("工程路径：", cur)

	conPath := filepath.Join(cur, "config.toml")

	if file == "" {
		file = conPath
	}
	_, err := toml.DecodeFile(file, &VideoWaterCon)
	if err != nil {
		panic(err)
	}

	fmt.Println("项目配置", *VideoWaterCon)

	return VideoWaterCon

}