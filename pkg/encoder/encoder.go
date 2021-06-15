package encoder

import (
	"image"
	"log"
)

type Encoding interface {
	Encode() error  // Takes any data and encodes it
	LoadKey() error // TODO when the encryption module is added
	Mode()          // Returns a helpful way to ask what mode it is
	NewEncoder()
}

type Encoder struct {
	path     string
	original image.Image
	final    image.Image
}

func NewEncoder(path string) *Encoder {
	img, err := sio.LoadImage(path)
	if err != nil {
		log.Fatal(err)
	}

	return &Encoder{path, img, img}
}
