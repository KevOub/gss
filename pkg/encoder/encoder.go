package encoder

import (
	"image"

	"github.com/KevOub/gss/pkg/util"
)

// EncoderInterface = the generic interface
type EncoderInterface interface {
	Init(path string, data []byte)
	Encode(data []byte) image.Image    // Takes any data and encodes it
	StoreByte(x int, y int, data byte) // Helper that takes byte and stores it based on the color model
	LoadKey()                          // TODO when the encryption module is added
	Mode() string                      // Returns a helpful way to ask what mode it is
}

type Encoder struct {
	Path     string
	Rect     image.Rectangle
	Data     []byte
	Original image.Image
	Final    image.Image
	EncoderInterface
}

func NewEncoder(path string, data []byte, enc EncoderInterface) *Encoder {
	var out Encoder
	out.Path = path
	out.Data = data
	out.Original, _ = util.LoadImage(path)
	out.Final, _ = util.LoadImage(path)
	out.Rect = out.Original.Bounds()
	out.EncoderInterface = enc
	return &out

}

// TODO find concise way to do sanity check
func (enc Encoder) Fits(data []byte) bool {
	return true
}

// Gross code that says pass an interface to the Encoder (who should be initialized) and Encode the data within
func (enc Encoder) Encode(w EncoderInterface, outpath string) {
	enc.EncoderInterface.Init(enc.Path, enc.Data) // sets up the wrapper
	out := enc.EncoderInterface.Encode(enc.Data)  // Stores the data and returns it as image.Image
	util.SaveImage(outpath, out)                  // Finally, save the encoded image
}
