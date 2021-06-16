package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/KevOub/gss/pkg/decoder"
	"github.com/KevOub/gss/pkg/encoder"
	"github.com/KevOub/gss/pkg/util"
)

/*
ENCODING:
The generalization of the encoder
	{ Encoder + &[EncoderInterface] }

Takes an encoder struct, with an interface
That interface has functions other ColorSpaces impelment so that they can encode
Then,

*/

func main() {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))

	// The colorspace to encode in
	var mode encoder.RGBA

	// The channel of the color space
	mode.DataChannel = encoder.Green

	// The data to send over
	junk := make([]byte, 720000)
	for i := 0; i < len(junk); i++ {
		junk[i] = byte(r.Intn(64))
	}

	buff := encoder.NewEncoder("assets/cat.png", junk, &mode)

	buff.Encode(&mode, "test.png")

	img1, _ := util.LoadImage("assets/cat.png")
	img2, _ := util.LoadImage("test.png")

	decoder.Delta(img1, img2)
	// fmt.Printf("%d %d", img.Out.Rect.Dx(), img.Out.Rect.Dy())

	fmt.Println("hello world")
}
