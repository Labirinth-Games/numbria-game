package utils

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
)

const MAP_TEMPLATE = "./assets/map.png"

func LoadImageTemplate() (width int, height int, pixels [][]Pixel) {

	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	file, err := os.Open(MAP_TEMPLATE)

	if err != nil {
		fmt.Println("Error to get file")
		os.Exit(1)
	}

	width, height, pixels, err = getPixels(file)

	if err != nil {
		fmt.Println("Error: Image could not be decoded")
		os.Exit(1)
	}

	defer file.Close()
	return
}

func getPixels(file io.Reader) (width int, height int, pixels [][]Pixel, err error) {
	img, _, err := image.Decode(file)

	if err != nil {
		return
	}

	bounds := img.Bounds()
	width, height = bounds.Max.X, bounds.Max.Y

	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
		}

		pixels = append(pixels, row)
	}

	return
}

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{
		int(r / 257),
		int(g / 257),
		int(b / 257),
		int(a / 257),
	}
}

type Pixel struct {
	R int
	G int
	B int
	A int
}
