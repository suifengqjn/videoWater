package common

import (
	"fmt"
	"github.com/BurntSushi/toml"
	file2 "myTool/file"
	"os"
	"path/filepath"
	"strings"
)

var VideoWaterCon *Config

type Config struct {
	VideoPath string `toml:"videoPath"`
	System int `toml:"system"`
	Task int `toml:"task"`
	Output string `toml:"output"`
	AppId string `toml:"appId"`
	CutSection CutSection
	Snip Snip
	Format Format
	FrameRate FrameRate
	BitRate BitRate
	CutFront CutFront	`toml:"cutFront"`
	CutBack CutBack	`toml:"cutBack"`
	Crop   Crop
	Crop1   Crop1 `toml:"crop1"`
	ClearWater ClearWater
	ClearWater1 ClearWater `toml:"clearWater1"`
	Mirror Mirror
	Resolution Resolution
	Compress Compress
	WaterText WaterText
	RunWaterText RunWaterText `toml:"RunWaterText"`
	WaterImage WaterImage
	AddBgm AddBgm `json:"addBgm"`
	Speed Speed
	FilmTitle FilmTitle
	FilmEnd FilmEnd

}


type CutSection struct {
	SectionPath string `toml:"sectionPath"`
	Switch int
	Duration int
}


type Snip struct {
	Switch int
	T int
	R int
}

type Format struct {
	Switch int
	Form string
}

type FrameRate struct {
	Switch int
	Value string
}

type BitRate struct {
	Switch int
	Value string
}

type CutFront struct {
	Switch int
	Value int
}

type CutBack struct {
	Switch int
	Value int
}

type Crop struct {
	Switch int
	Start int64
	Duration int64
	X int
	Y int
	W int
	H int
}

type Crop1 struct {
	Switch int
	Start int64
	Duration int64
	Left int
	Right int
	Top int
	Bottom int
}

type ClearWater struct {
	Switch int
	X int
	Y int
	W int
	H int
}

type Mirror struct {
	Switch int
	Direction string
}

type Resolution struct {
	Switch int
	W int
	H int
}

type Compress struct {
	Switch int
	Preset string
	Crf int
}

type WaterText struct {
	Switch int
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
	Switch  int
	Content string
	Path    string
	Size    int
	Color   string
	IsTop   int `toml:"isTop"`
	LeftToRight   int `toml:"leftToRight"`
	Sp      int
}

type WaterImage struct {
	Switch int
	Path string
	Style int
	Sp1 int
	Sp2 int
}

type AddBgm struct {
	Switch int
	Cover int
	Bgm string
}

type Speed struct {
	Switch int
	V string
}

type FilmTitle struct {
	Switch int
	Path   string
}

type FilmEnd struct {
	Switch int
	Path   string
}


func ReadConfig(file string) *Config  {
	if VideoWaterCon != nil {
		return VideoWaterCon
	}

	conPath := ""
	cur, _ := os.Getwd()
	if file == "" {
		conPath = filepath.Join(cur, "config.toml")

		if file2.PathExist(conPath) == false {
			conPath = os.ExpandEnv("$HOME") + "/Desktop/vm/config.toml"
		}

		file = conPath

	}

	_, err := toml.DecodeFile(file, &VideoWaterCon)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if strings.HasPrefix(VideoWaterCon.VideoPath,"./") {

	} else {
		VideoWaterCon.VideoPath = filepath.Join(cur,VideoWaterCon.VideoPath)
	}

	return VideoWaterCon

}