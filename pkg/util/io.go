package util

import (
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

//LoadImage takes path string and returns and image
func LoadImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)

	return image, err
}

//SaveImage outputs the image.Image
func SaveImage(path string, img image.Image) {
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

func LoadFileToByte(filename string) ([]byte, error) {

	var err error
	var data []byte
	data, err = ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}
