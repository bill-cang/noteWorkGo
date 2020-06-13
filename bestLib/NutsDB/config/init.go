package config

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"strings"
)

//初始化文件保存路径
func InitConfig() {
	var err error
	defer func() {
		if err != nil {
			panic(err)
		}
	}()
	//解析Toml
	var dp *dbFilePath
	dp, err = analysisToml("conf.toml")
	if err != nil {
		return
	}

	var way string
	way, err = os.Getwd()
	if err != nil {
		log.Printf("[config] Getwd err =%+v", err)
		return
	}
	way = strings.Replace(way, "\\", "/", -1)

	Opt.Dir = fmt.Sprintf("%s%s%s", way, "/", dp.FilePath.FilePath)
	log.Printf("[config] init success. Opt.Dir =%s", Opt.Dir)
}

type dbFilePath struct {
	FilePath database `toml:"database"`
}

type database struct {
	FilePath string `toml:"file_path"`
}

func analysisToml(tPath string) (dp *dbFilePath, err error) {
	dp = new(dbFilePath)
	file := *flag.String("config", tPath, "Path to toml config file.")
	flag.Parse()
	if _, err = toml.DecodeFile(file, dp); err != nil {
		fmt.Printf("DecodeFile err =%+v\n", err)
		return nil, err
	}
	dp.FilePath.FilePath = strings.Replace(dp.FilePath.FilePath, "\\", "/", -1)
	return
}
