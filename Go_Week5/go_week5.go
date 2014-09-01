package main

import (
	"bufio"
	"code.google.com/p/draw2d/draw2d"
	"fmt"
	"image"
	"image/png"
	"log"
	"math"
	"os"
)

func saveToPngFile(filePath string, m image.Image) {
	f, err := os.Create(filePath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	b := bufio.NewWriter(f)
	err = png.Encode(b, m)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Wrote %s OK.\n", filePath)
}

func main() {
	i := image.NewRGBA(image.Rect(0, 0, 600, 600))
	gc := draw2d.NewGraphicContext(i)
	xc, yc := 100.0, 100.0
	radiusX, radiusY := 100.0, 100.0
	startAngle := 45.0 * (math.Pi / 180.0) /* angles are specified */
	angle := 360 * (math.Pi / 180.0)       /* in radians           */
	gc.SetLineWidth(10)
	gc.SetLineCap(draw2d.ButtCap)
	gc.SetStrokeColor(image.Black)
	gc.ArcTo(xc, yc, radiusX, radiusY, startAngle, angle)
	gc.Stroke()
	// fill a circle
	gc.SetLineWidth(6)

	//	gc.MoveTo(xc, yc)
	//gc.LineTo(xc+math.Cos(startAngle)*radiusX, yc+math.Sin(startAngle)*radiusY)
	//	gc.MoveTo(xc, yc)
	//gc.LineTo(xc-radiusX, yc)
	gc.Stroke()

	gc.ArcTo(xc, yc, 10.0, 10.0, 0, 2*math.Pi)
	gc.Fill()
	saveToPngFile("TestDrawArc.png", i)

	gc.MoveTo(10.0, 10.0)
	gc.LineTo(100.0, 10.0)
	gc.Stroke()
	saveToPngFile("TestPath.png", i)
	draw2d.RoundRect(gc, 5, 5, 95, 95, 10, 10)
	gc.SetStrokeColor(image.Black)
	gc.FillStroke()
	gc.SetFontSize(18)
	gc.MoveTo(10, 52)
	gc.SetFontData(draw2d.FontData{"luxi", draw2d.FontFamilyMono, draw2d.FontStyleBold | draw2d.FontStyleItalic})
	width := gc.FillString("cou")
	fmt.Printf("width: %f\n", width)
	gc.RMoveTo(width+1, 0)
	gc.FillString("cou")
	saveToPngFile("TestFillString.png", i)
	width = gc.FillString("1234124512512512")

	fmt.Printf("width: %f\n", width)
	saveToPngFile("TestFillString.png", i)
}
