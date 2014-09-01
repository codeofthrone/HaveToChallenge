package main

import (
	"code.google.com/p/draw2d/draw2d"
	qrcode "code.google.com/p/go-qrcode"
	"flag"
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func main() {
	var url string
	var target_file string
	var AndroidApi string
	var ScreenSize string
	var Density string
	var DeviceID string
	flag.StringVar(&url, "Url", "url", "Qr code content")
	flag.StringVar(&target_file, "Output", "Target", "output file name")
	flag.StringVar(&DeviceID, "ID", "Device ID", "Device ID")
	flag.StringVar(&AndroidApi, "API", "API", "Android api level")

	flag.StringVar(&ScreenSize, "Size", "ScreenSize", "screen size")
	flag.StringVar(&Density, "Den", "Den", "Density")
	flag.Parse()
	fmt.Println("url\t:", url)
	fmt.Println("Output\t:", target_file)
	fmt.Println("ID\t:", DeviceID)
	fmt.Println("API\t:", AndroidApi)
	fmt.Println("Size\t:", ScreenSize)
	fmt.Println("Den\t:", Density)
	var q *qrcode.QRCode
	q, err := qrcode.New(url, qrcode.Medium)
	qrimg := q.Image(300)
	m := image.NewRGBA(image.Rect(0, 0, 1200, 270)) //*NRGBA (image.Image interface)
	draw.Draw(m, m.Bounds(), qrimg, image.ZP, draw.Over)
	gc := draw2d.NewGraphicContext(m)
	draw2d.Rect(gc, 280, -50, 1200, 270)
	gc.SetStrokeColor(image.Black)
	gc.SetFillColor(image.White)
	gc.FillStroke()
	gc.SetFillColor(image.Black)
	gc.SetFontSize(200)
	gc.Translate(280, 240)
	width := gc.FillString(DeviceID)
	gc.Translate(width+10, -143)
	gc.SetFontSize(64)
	gc.FillString("API   : " + AndroidApi)
	gc.Translate(0, 80)
	gc.FillString("SIZE : " + ScreenSize)
	gc.Translate(0, 80)
	gc.FillString(Density)
	w, err := os.Create(target_file + ".jpg")
	defer w.Close()
	jpeg.Encode(w, m, nil)
	file, err := os.Open(target_file + ".jpg")
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	im := resize.Resize(180, 0, img, resize.Lanczos3)

	out, err := os.Create(target_file + ".jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, im, nil)

}
