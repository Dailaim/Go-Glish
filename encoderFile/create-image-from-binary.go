package encoderFile

import (
	"image"
	"image/color"
)

func CreateImageFromBinary(binary []bool, imgProps *ImageProperties, rest int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, imgProps.Width, imgProps.Height))

	for y := 0; y < imgProps.Height; y += imgProps.PixelSize {
		for x := 0; x < imgProps.Width; x += imgProps.PixelSize {
			index := (y/imgProps.PixelSize)*(imgProps.Width/imgProps.PixelSize) + (x / imgProps.PixelSize)
			if index >= len(binary) {
				break
			}
			colorVal := color.White
			if binary[index] {
				colorVal = color.Black
			}
			for py := y; py < y+imgProps.PixelSize; py++ {
				for px := x; px < x+imgProps.PixelSize; px++ {
					if px < imgProps.Width && py < imgProps.Height {
						img.Set(px, py, colorVal)
					}
				}
			}
		}
	}

	// Llenar el área restante con un color de fondo (en este caso, morado)
	if rest > 0 {
		for y := imgProps.Height - rest; y < imgProps.Height; y += imgProps.PixelSize {
			for x := 0; x < imgProps.Width; x += imgProps.PixelSize {
				colorVal := color.RGBA{0x80, 0x00, 0x80, 0xff} // Púrpura
				for py := y; py < y+imgProps.PixelSize; py++ {
					for px := x; px < x+imgProps.PixelSize; px++ {
						if px < imgProps.Width && py < imgProps.Height {
							img.Set(px, py, colorVal)
						}
					}
				}
			}
		}
	}

	return img
}
