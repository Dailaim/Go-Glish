package encoderFile

import (
	"fmt"
	"image/png"
	"os"
	"sync"
)

func CreateFrames(imgProps *ImageProperties, resta int, id int, bits []bool, wg *sync.WaitGroup) {
	defer wg.Done()

	img := CreateImageFromBinary(bits, imgProps, resta)

	file, err := os.Create(fmt.Sprintf("%s/frame%04d.png", imgProps.TmpDir, id))
	defer file.Close()

	if err != nil {
		panic(err)
	}

	err = png.Encode(file, img)

	if err != nil {
		panic(err)
	}
}
