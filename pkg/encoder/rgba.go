package encoder

import (
	"image"
	"image/color"
	"image/draw"
)

type DataChannel byte

const (
	Red DataChannel = iota
	Green
	Blue
	Alpha
)

type RGBA struct {
	*Encoder
	// In          *image.RGBA
	// Out         *image.RGBA
	DataChannel DataChannel // Which Channel stores the data
}

func imgtorgb(src image.Image) *image.RGBA {
	b := src.Bounds()
	m := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(m, m.Bounds(), src, b.Min, draw.Src)
	return m
}

func (enc *RGBA) Init(path string, data []byte) {
	enc.Encoder = NewEncoder(path, data, enc)
}

// // Transcode takes the image and whips together two new images to process later
// func (enc *RGBA) Transcode() {

// 	// b := enc.Original.Bounds()
// 	// m := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
// 	// draw.Draw(m, m.Bounds(), enc.Original, b.Min, draw.Src)
// 	enc.In = imgtorgb(enc.Original)
// 	enc.Out = imgtorgb(enc.Original)

// 	// enc.In = image.NewRGBA(enc.Original.Bounds())
// 	// enc.Out = image.NewRGBA(enc.Original.Bounds())
// }

func (enc *RGBA) Encode(data []byte) image.Image {

	upperBounds := enc.Rect
	upperX, upperY := upperBounds.Max.X, upperBounds.Max.Y

	index := 0

	for i := 0; i < upperY; i++ {
		for j := 0; j < upperX; j++ {
			// Error check for sanity
			if index < len(data) {
				enc.StoreByte(j, i, data[index])
			}
			// else {
			// log.Fatal("Wrote too much data")
			// break

			// }

			index++

		}
	}

	return enc.Final

}

func (enc *RGBA) StoreByte(x int, y int, data byte) {

	r, g, b, a := enc.Original.At(x, y).RGBA() // the old r , g , b values

	newColor := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}

	switch enc.DataChannel {
	case Red:
		newColor.R += data
	case Green:
		newColor.G += data
	case Blue:
		newColor.B += data

	}
	enc.Final.(*image.RGBA).Set(x, y, newColor)

}

func (img *RGBA) LoadKey() {

}

func (img *RGBA) Mode() string {
	return "RGBA"
}
