package encoder

import (
	"image"
	"log"
	"os"
)

type Encoding interface {
	Encode()      // Takes any data and encodes it
	LoadKey()     // TODO when the encryption module is added
	Mode() string // Returns a helpful way to ask what mode it is
	NewEncoder()
}

type Encoder struct {
	Path     string
	Original image.Image
	Final    image.Image
}

func loadImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	return image, err
}

func NewEncoder(path string) *Encoder {
	img, err := loadImage(path)
	if err != nil {
		log.Fatal(err)
	}

	return &Encoder{path, img, img}
}
