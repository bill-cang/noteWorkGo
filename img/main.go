package main

import (
	"fmt"
	opencv "github.com/go-opencv/go-opencv/opencv"
	"os"
	"path"
	"runtime"
)

func main() {
	_, currentfile, _, _ := runtime.Caller(0)
	filename := path.Join(path.Dir(currentfile), "E:/Tmp/mzx.png")
	if len(os.Args) == 2 {
		filename = os.Args[1]
	}

	image := opencv.LoadImage(filename)
	if image == nil {
		panic("LoadImage fail")
	}
	defer image.Release()

	win := opencv.NewWindow("Go-OpenCV")
	defer win.Destroy()

	win.SetMouseCallback(func(event, x, y, flags int) {
		fmt.Printf("event = %d, x = %d, y = %d, flags = %d\n",
			event, x, y, flags,
		)
	})
	win.CreateTrackbar("Thresh", 1, 100, func(pos int) {
		fmt.Printf("pos = %d\n", pos)
	})

	win.ShowImage(image)

	opencv.WaitKey(0)
}
