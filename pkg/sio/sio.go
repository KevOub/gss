package sio

import (
	"image"
	"os"
)

// Enum to determine the mode to use
type ModeSelector uint8

const (
	Encode ModeSelector = iota
	Decode
)

// prints the mode
func (d ModeSelector) String() string {
	return [...]string{"Encode", "Decode"}[d]
}

// TODO flesh out to solidify whether it should be IO or
type IO struct {
	Path string
	Mode ModeSelector
}

func LoadImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	return image, err
}
