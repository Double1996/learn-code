package main

import (
	"flag"
	"fmt"
	"github.com/fogleman/gg" // 需要安装这个包
	"github.com/golang/freetype/truetype"
	"io/ioutil"
	"log"
	"math"
)

var (
	dpi      = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
	fontfile = flag.String("fontfile", "./Alibaba-PuHuiTi-Regular.ttf", "filename of the ttf font")
	size     = flag.Float64("size", 22, "font size in points")
)

// 知乎防水印 https://studygolang.com/articles/13632
func main() {
	drawImageBygg()
}

func drawImageBygg() {
	dc := gg.NewContext(320, 180) // 56 => w*sin(45) + h*sin(45)  45度时，字体达到最大高度
	dc.Clear()
	//dc.SetRGBA(0, 0, 0, 1) // 暗水印
	dc.SetRGBA(235, 235, 235, 0.1) // 明水印

	fontBytes, err := ioutil.ReadFile(*fontfile)
	if err != nil {
		log.Println(err)
		return
	}
	font, err := truetype.Parse(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}
	face := truetype.NewFace(font, &truetype.Options{
		Size: *size,
		DPI:  *dpi,
	})
	dc.SetFontFace(face)

	// 初始化用于计算坐标的变量
	fm := face.Metrics()
	ascent := float64(fm.Ascent.Round()) + 120 // 字体的基线到顶部距离
	decent := float64(fm.Descent.Round()) + 80 // 字体的基线到底部的距离
	w := float64(fm.Height.Round())            // 方块字，大多数应为等宽字，即和高度一样
	h := float64(fm.Height.Round())
	totalWidth := 30.0 // 目前已累积的图片宽度（需要用来计算字体位置）

	// 随机取汉字，定位倒立的字
	words := []string{"49测试0219887"} // 取8个字

	for _, word := range words {
		degree := -30.0
		x, y, leftCutSize, rightCS := getCoordByQuadrantAndDegree(w, h, ascent, decent, degree, totalWidth)
		dc.RotateAbout(gg.Radians(degree), 0, 0)
		dc.DrawStringAnchored(word, x, y, 0, 0)
		dc.RotateAbout(-1*gg.Radians(degree), 0, 0)
		totalWidth = totalWidth + leftCutSize + rightCS
		fmt.Println("x:", x, "y:", y, "total:", totalWidth, "degree:", degree)
	}

	dc.Stroke()
	dc.SavePNG("out.png")
	fmt.Println("Wrote out.png OK.")
}

func getCoordByQuadrantAndDegree(w, h, ascent, descent, degree, beforTotalWidth float64) (x, y, leftCutSize, rightCutSize float64) {
	var totalWidth float64
	switch {
	case degree <= 0 && degree >= -40: // 第一象限：逆时针 -30度 ~ 0  <=>  330 ~ 360 （目前参数要传入负数）
		rd := -1 * degree // 转为正整数，便于计算
		leftCutSize = w * getDegreeSin(90-rd)
		rightCutSize = h * getDegreeSin(rd)

		offset := (leftCutSize + rightCutSize - w) / 2 // 横向偏移量（角度倾斜越厉害，占宽越多，通过偏移量分摊给它的左右边距来收窄）
		leftCutSize, rightCutSize = leftCutSize-offset, rightCutSize-offset

		totalWidth = beforTotalWidth + leftCutSize
		x = getDegreeSin(90-rd)*totalWidth - w
		y = ascent + getDegreeSin(rd)*totalWidth
	case degree >= 0 && degree <= 40: // 第四象限：顺时针 0 ~ 30度
		leftCutSize = h * getDegreeSin(degree)
		rightCutSize = w * getDegreeSin(90-degree)

		offset := (leftCutSize + rightCutSize - w) / 2
		leftCutSize, rightCutSize = leftCutSize-offset, rightCutSize-offset

		totalWidth = beforTotalWidth + leftCutSize // 现在totalwidth = 前面的宽 + 自己的左切边
		x = getDegreeSin(90-degree) * totalWidth
		y = ascent - getDegreeSin(degree)*totalWidth
	case degree >= 180 && degree <= 220: // 第二象限：顺时针 180 ~ 210度
		rd := degree - 180
		leftCutSize = h * getDegreeSin(rd)
		rightCutSize = w * getDegreeSin(90-rd)

		offset := (leftCutSize + rightCutSize - w) / 2
		leftCutSize, rightCutSize = leftCutSize-offset, rightCutSize-offset

		totalWidth = beforTotalWidth + leftCutSize
		x = -1 * (getDegreeSin(90-rd)*totalWidth + w)
		y = getDegreeSin(rd)*totalWidth - descent
	case degree >= 140 && degree <= 180: // 第三象限：顺时针 150 ~ 180度
		rd := 180 - degree
		leftCutSize = w * getDegreeSin(90-rd)
		rightCutSize = h * getDegreeSin(rd)

		offset := (leftCutSize + rightCutSize - w) / 2
		leftCutSize, rightCutSize = leftCutSize-offset, rightCutSize-offset

		totalWidth = beforTotalWidth + leftCutSize
		x = -1 * (getDegreeSin(90-rd) * totalWidth)
		y = -1 * (getDegreeSin(rd)*totalWidth + descent)
	default:
		panic(fmt.Sprintf("非法的参数：%f", degree))
	}
	return
}

func getDegreeSin(degree float64) float64 {
	return math.Sin(degree * math.Pi / 180)
}
