package decoder

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func Delta(img1 image.Image, img2 image.Image) {
	// Code to assure images are same size
	mx2, my2 := img2.Bounds().Max.X, img2.Bounds().Max.Y
	mx1, my1 := img1.Bounds().Max.X, img1.Bounds().Max.Y
	if mx1 != mx2 || my1 != my2 {
		log.Fatal("Images are not same size")
	}

	output := image.NewRGBA(img1.Bounds())

	for i := 0; i < my1; i++ {
		for j := 0; j < mx1; j++ {

			r1, g1, b1, _ := img1.At(j, i).RGBA()
			r2, g2, b2, a2 := img2.At(j, i).RGBA()

			newColor := color.RGBA{uint8(r2 - r1), uint8(g2 - g1), uint8(b2 - b1), uint8(a2)}

			output.Set(j, i, newColor)

		}
	}

	saveImage("delta.png", output)

}

func saveImage(path string, img image.Image) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Encode to `PNG` with `DefaultCompression` level
	// then save to file
	err = png.Encode(f, img)
	if err != nil {
		log.Fatal(err)
	}

}
