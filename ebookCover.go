package ebookdownloader

import (
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/goki/freetype"
)

const (
	fontSize = 40 //字体尺寸
)

//GenerateCover 生成封面 cover.jpg
func GenerateCover(this BookInfo) {

	//需要添加内容的图片
	coverAbs, _ := filepath.Abs("./tpls/cover.jpg")
	//fmt.Println(coverAbs)
	imgfile, err := os.Open(coverAbs)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer imgfile.Close()

	jpgimg, _ := jpeg.Decode(imgfile)
	img := image.NewNRGBA(jpgimg.Bounds())

	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			img.Set(x, y, jpgimg.At(x, y))
		}
	}

	//需要一个ttf字体文件
	fontAbs, _ := filepath.Abs("./fonts/WenQuanYiMicroHei.ttf")
	fontBytes, err := ioutil.ReadFile(fontAbs)
	if err != nil {
		log.Println(err.Error())
	}

	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err.Error())
	}

	f := freetype.NewContext()
	f.SetDPI(72)
	f.SetFont(font)
	f.SetFontSize(fontSize)
	f.SetClip(jpgimg.Bounds())
	f.SetDst(img)
	f.SetSrc(image.Black) //设置字体颜色

	pt := freetype.Pt(img.Bounds().Dx()-370, img.Bounds().Dy()-590) //字体出现的位置
	//尝试把字符串，坚着写入图片中
	NameRune := []rune(this.Name)
	f.DrawString(string(NameRune[0]), pt) // 第一个中文字符
	for index := 1; index < len(NameRune); index++ {
		pt.Y += f.PointToFixed(50)
		f.DrawString(string(NameRune[index]), pt) //写入 小说名
	}

	f.SetFontSize(35)                                                     //重新设置 字体大小为35
	ptAuthor := freetype.Pt(img.Bounds().Dx()-320, img.Bounds().Dy()-500) //字体出现的位置
	f.DrawString(this.Author+" ©著", ptAuthor)                             //写入小说作者名

	newCoverpath, _ := filepath.Abs("./cover.jpg")
	newfile, err := os.Create(newCoverpath)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer newfile.Close()

	err = jpeg.Encode(newfile, img, &jpeg.Options{Quality: 100})
	if err != nil {
		fmt.Println(err.Error())
	}
}
