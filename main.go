package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"strings"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"

	"github.com/icza/mjpeg"
)

func main() {
	label := "Lionel Andrés Messi Cuccittini is an Argentine professional | footballer who plays as a forward and captains both | Spanish club Barcelona and the Argentina national team."
	img := image.NewRGBA(image.Rect(0, 0, 800, 500))
	// img := image.Alpha(image.Rect(0, 0, 800, 500))

	// img.SetRGBA(0, 0, color.RGBA{255, 255, 255, 250})

	addLabel(img, 50, 30, label)

	/* f, err := os.Create("hello-go.png")
	if err != nil {
		panic(err)
	}
	defer f.Close() */
	buf := &bytes.Buffer{}

	if err := jpeg.Encode(buf, img, nil); err != nil {
		panic(err)
	}

	renderVid(buf)
}

func addLabel(img *image.RGBA, x, y int, label string) {
	// col := color.RGBA{200, 100, 0, 255}

	background := image.NewUniform(color.Black)
	draw.Draw(img, img.Bounds(), background, image.ZP, draw.Src)
	labelArr := strings.Split(label, "|")
	for i, v := range labelArr {
		h := i * 13
		hf := fixed.Int26_6(y*64) + fixed.Int26_6(h*64)
		fmt.Println(hf)
		point := fixed.Point26_6{fixed.Int26_6(x * 64), hf}
		d := &font.Drawer{
			Dst:  img,
			Src:  image.NewUniform(color.White),
			Face: basicfont.Face7x13,
			Dot:  point,
		}
		d.DrawString(v)
	}

}

func renderVid(buf *bytes.Buffer) {
	aw, err := mjpeg.New("test.avi", 200, 100, 2)
	checkErr(err)
	// Create a movie from bytes frame

	checkErr(err)
	for i := 1; i <= 50; i++ {
		checkErr(aw.AddFrame(buf.Bytes()))

	}
	// }

	checkErr(aw.Close())
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
