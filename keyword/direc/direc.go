package direc

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GetDirSon01(dirPath string) (files []string, dirs []string, err error) {
	var dir []os.FileInfo
	dir, err = ioutil.ReadDir(dirPath)
	if err != nil {
		fmt.Printf("[ReadDir] err =%+v", err)
		return
	}

	files = make([]string, 0)
	dirs = make([]string, 0)
	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPath+PthSep+fi.Name())
			//GetFolderCore(dirPath + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			//ok := strings.HasSuffix(fi.Name(), ".go")
			files = append(files, dirPath+PthSep+fi.Name())
		}
	}
	return
}

//获得文件夹下的子文件夹和文件
func GetDirSon02(dirPth string) (files []string, dirs []string, err error) {
	files = make([]string, 0)
	dirs = make([]string, 0)
	PthSep := string(os.PathSeparator)
	err = filepath.Walk(dirPth, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			dir := fmt.Sprintf("%s%s%s", dirPth, PthSep, f.Name())
			dirs = append(dirs, dir)
		} else {
			file := fmt.Sprintf("%s%s%s", dirPth, PthSep, f.Name())
			files = append(files, file)
		}
		return nil
	})
	return
}

/*遍历文件夹下的文件，输出MD5值以及文件大小*/
