package main

import (
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"unicode/utf8"

	"github.com/disintegration/imaging"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

const (
	FontSize     = 8
	SquareHeight = 200
)

func main() {
	HandleUserImage()
}

// HandleUserImage paste user image onto background
func HandleUserImage() (string, error) {
	bm, err := imaging.Open("bg.png")
	if err != nil {
		fmt.Println("open file failed", err)
		return "", err
	}

	// 图片按比例缩放
	dst := imaging.Resize(bm, bm.Bounds().Max.X, SquareHeight, imaging.Lanczos)
	// 将图片粘贴到背景图的固定位置 最后的参数控制虚化程度
	//result := imaging.Overlay(bm, dst, image.Pt(0, bm.Bounds().Max.Y-SquareHeight), 0.3)
	writeOnImage(dst)

	fileName := fmt.Sprintf("%d.jpg", 1234)
	err = imaging.Save(dst, fileName)
	if err != nil {
		return "", err
	}

	return fileName, nil
}

func rotate90(m image.Image) image.Image {
	rotate90 := image.NewRGBA(image.Rect(0, 0, m.Bounds().Dy(), m.Bounds().Dx()))
	// 矩阵旋转
	for x := m.Bounds().Min.Y; x < m.Bounds().Max.Y; x++ {
		for y := m.Bounds().Max.X - 1; y >= m.Bounds().Min.X; y-- {
			//  设置像素点
			rotate90.Set(m.Bounds().Max.Y-x, y, m.At(y, x))
		}
	}
	return rotate90
}

func writeOnImage(target *image.NRGBA) {
	c := freetype.NewContext()

	// 设置屏幕每英寸的分辨率
	c.SetDPI(256)
	// 背景
	c.SetClip(target.Bounds())
	// 设置目标图像
	c.SetDst(target)
	c.SetHinting(font.HintingFull)

	// 设置文字颜色、字体、字大小
	c.SetSrc(image.NewUniform(color.RGBA{R: 220, G: 220, B: 220, A: 220}))
	// 以磅为单位设置字体大小
	c.SetFontSize(FontSize)
	fontFam, err := getFontFamily()

	if err != nil {
		fmt.Println("get font family error")
	}
	// 设置用于绘制文本的字体
	c.SetFont(fontFam)

	drawStr := "测试10211212"

	// 获取字体的尺寸大小
	fixed := c.PointToFixed(FontSize)
	// fixed.Ceil() 字体大小
	// utf8.RuneCountInString(drawStr) 获取字符串的实际大小，而不是以byte算
	//pt := freetype.Pt(target.Rect.Max.X/2-(utf8.RuneCountInString(drawStr)/2)*fixed.Ceil(), target.Rect.Max.Y-SquareHeight+SquareHeight/2)
	//
	//fontPin, err := c.DrawString(drawStr, pt)
	//if err != nil {
	//	fmt.Printf("draw error: %v \n", err)
	//	return
	//}

	_, err = c.DrawString(drawStr, freetype.Pt(target.Rect.Max.X/2-(utf8.RuneCountInString(drawStr)/2)*fixed.Ceil(), target.Rect.Max.X-(utf8.RuneCountInString(drawStr))*fixed.Ceil()))
	if err != nil {
		fmt.Printf("draw error: %v \n", err)
		return
	}
}

// 获取字符集，仅调用一次
func getFontFamily() (*truetype.Font, error) {
	// 这里需要读取中文字体，否则中文文字会变成方格
	fontBytes, err := ioutil.ReadFile("Alibaba-PuHuiTi-Regular.ttf")
	if err != nil {
		fmt.Println("read file error:", err)
		return &truetype.Font{}, err
	}

	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		fmt.Println("parse font error:", err)
		return &truetype.Font{}, err
	}

	return f, err
}
