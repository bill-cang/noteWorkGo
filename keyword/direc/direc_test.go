package direc

import (
	"fmt"
	"noteWork/bestLib/utils"
	"noteWork/bestLib/utils/encryption"
	"testing"
	"time"
)

func TestReadFileMd5AndSize(t *testing.T) {
	_, dirs, err := GetDirSon01(`E:\Tmp\Audio  aac`)
	defer func(err error) {
		if err != nil {
			fmt.Printf("[TestReadFileMd5AndSize] err =%+v", err)
		}
	}(err)

	if err != nil {
		return
	}

	for _, dir := range dirs {
		files, _, err := GetDirSon01(dir)
		if err != nil {
			return
		}
		//fmt.Printf("files =%+v\n", files)
		for _, file := range files {
			digest, err := encryption.Md5File(file)
			if err != nil {
				return
			}
			fileSize := utils.GetFileSize(file)
			fmt.Printf("%s:\t%s\t%d\n", file, digest, fileSize)
		}
	}
}

func TestGetFolderCore(t *testing.T) {
	files, _, err := GetDirSon02(`E:\Tmp\Greetings audio aac\男声`)
	if err != nil {
		fmt.Printf("err =%+v\n", err)
		return
	}
	fmt.Printf("files = %+v\n", files)
}

func TestGetDirSon01(t *testing.T) {
	fmt.Printf("ss = %d\n", int(time.Now().Unix()))
}
