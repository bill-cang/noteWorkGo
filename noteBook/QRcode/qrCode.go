package QRcode

import (
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
	"strings"
)

func GetQRCode(SourceImgPath string) (posterPath string, err error) {

	// qr.Encode (编码内容，容错等级，编码形式)
	var qrCode barcode.Barcode
	if qrCode, err = qr.Encode("https://www.baidu.com", qr.H, qr.Auto); err != nil {
		log.Printf("[GetQRCode] Encode err =%+v", err)
		return
	}
	// barcode.Scale() 对二维码进行缩放
	qrCode, _ = barcode.Scale(qrCode, 128, 128)
	// 将图片写入文件 (如果是生成海报，可暂先不写入文件)
	index := strings.LastIndex(SourceImgPath, ".")
	posterPath = fmt.Sprintf("%s_cp%s", SourceImgPath[:index], SourceImgPath[index:])
	file, _ := os.Create(posterPath)
	png.Encode(file, qrCode)

	// 引入海报图片
	file, _ = os.Open(SourceImgPath)
	// 为了记录，这里就不检测 err 了
	img, _, err := image.Decode(file)

	// 开始绘制
	// image.NewNRGBA(图像的大小) 这里因为要把二维码放在海报上，所以传入海报的大小
	posterImg := image.NewNRGBA(img.Bounds())

	// draw.Draw(被绘制的图片, 绘制框的大小, 要绘制的图片, 绘制的位置, 绘制类型)
	// 先把海报背景画上
	draw.Draw(posterImg, posterImg.Bounds(), img, image.Pt(0, 0), draw.Over)
	// 再把二维码画上，需要注意的是坐标.
	// 这里的 qrCode 就是上面使用 barcode 生成的二维码
	draw.Draw(posterImg, posterImg.Bounds(), qrCode, image.Pt(-50, -1020), draw.Over)

	// 绘制好后保存到文件中
	posterFile, _ := os.Create("02.png")
	png.Encode(posterFile, posterImg)

	return
}
