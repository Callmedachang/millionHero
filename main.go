package main

import (
	"github.com/hunterhug/go_image"
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"errors"
	"runtime/debug"
	"log"
	"os/exec"
	"MillionHeroes/utils"
)

var x, y, w, h int

var x1, y2, w3, h4 int

var path = "/Users/bighandsome/Documents/goWorkPlace/src/MillionHeroes"
//var path = "/Users/shuuharushi/millionHero"

func main() {
	defer func() {
		if e := recover(); e != nil {
			log.Printf("%s: %s", e, debug.Stack())
			fmt.Print("程序已崩溃，按任意键退出")
			var c string
			fmt.Scanln(&c)
		}
	}()
	fmt.Print("哪个场？\n1:-------芝士-------\n2:-------冲顶-------\n3:-------西瓜-------\n4:------一直播------\n选择:")
	var c int
	fmt.Scanln(&c)
	fmt.Println("参数配置成功")
	if c == 1 {
		//芝士超人参数
		x, y, w, h = 0, 100, 500, 150
		x1, y2, w3, h4 = 0, 100, 500, 400
	} else if c == 2 {
		x, y, w, h = 0, 160, 500, 200
		x1, y2, w3, h4 = 0, 160, 500, 450
	} else if c == 3 {
		x, y, w, h = 0, 150, 500, 155
		x1, y2, w3, h4 = 0, 160, 500, 450
	} else if c == 4 {
		x, y, w, h = 0, 150, 500, 155
		x1, y2, w3, h4 = 0, 160, 500, 450
	}
	for {
		fmt.Println("等题目出现完全时候按回车截图搜索")
		var c string
		fmt.Scanln(&c)
		fmt.Println("idevicescreenshot", path+"/1.jpg")
		icmd := exec.Command("idevicescreenshot", path+"/1.jpg")
		icmd.Run()
		err:=go_image.ScaleF2F(path+"/1.jpg", path+"/2.jpg", 500)
		fmt.Println("111111111111",err)
		ii, _, _ := go_image.LoadImage(path + "/2.jpg")
		ii1, _ := ImageCopy(ii, x, y, w, h)
		SaveImage(path+"/3.jpg", ii1)

		ii2, _ := ImageCopy(ii, x1, y2, w3, h4)
		SaveImage(path+"/4.jpg", ii2)

		go func() {
			result := utils.OCR(path + "/3.jpg")
			icmd2 := exec.Command("open", "https://www.baidu.com/s?chrome=UTF-8&wd="+result)
			icmd2.Run()
		}()
		go func() {
			result1 := utils.OCR(path + "/4.jpg")
			icmd1 := exec.Command("open", "-a", "/Applications/Google Chrome.app", "https://www.baidu.com/s?ie=UTF-8&wd="+result1)
			icmd1.Run()
		}()
	}
}

func ImageCopy(src image.Image, x, y, w, h int) (image.Image, error) {

	var subImg image.Image

	if rgbImg, ok := src.(*image.YCbCr); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+w, y+h)).(*image.YCbCr) //图片裁剪x0 y0 x1 y1
	} else if rgbImg, ok := src.(*image.RGBA); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+w, y+h)).(*image.RGBA) //图片裁剪x0 y0 x1 y1
	} else if rgbImg, ok := src.(*image.NRGBA); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+w, y+h)).(*image.NRGBA) //图片裁剪x0 y0 x1 y1
	} else {

		return subImg, errors.New("图片解码失败")
	}

	return subImg, nil
}

func SaveImage(p string, src image.Image) error {
	f, err := os.OpenFile(p, os.O_SYNC|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	err = jpeg.Encode(f, src, &jpeg.Options{Quality: 80})
	return err
}
