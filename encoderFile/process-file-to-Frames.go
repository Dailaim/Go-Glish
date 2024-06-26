package encoderFile

import (
	"os"
	"sync"
)

func ProcessFileToFrames(filename string, framesDir string, imgProps *ImageProperties) {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	os.Mkdir(imgProps.TmpDir, os.ModePerm)

	defer file.Close()

	bufferSize := (1080 * 1920 / 8) / imgProps.PixelSize
	var wg sync.WaitGroup

	counter := 0

	print("Reading file... ")
	for {
		counter++

		data := make([]byte, bufferSize)

		n, err := file.Read(data)

		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}

		if n == 0 {
			break
		}

		bits := make([]bool, 0)

		// Iterar sobre cada valor en el buffer
		for _, b := range data {

			// Iterar sobre cada bit del valor
			for j := 7; j >= 0; j-- {
				// Imprimir cada bit del byte
				bit := ((b >> j) & 1) == 1

				// Enviar el bit a través del canal
				bits = append(bits, bit)
			}

		}

		wg.Add(1)

		go CreateFrames(imgProps, bufferSize-len(data), counter, bits, &wg)

	}

	print("File readed. \n")
	wg.Wait()
	println("Frames created.")
}

// create frame save original name and the fila

// originalPath := strings.Split(filename, "/")

// originalFile := strings.Split(originalPath[len(originalPath)-1], ".")

// originalName, originalExtension := originalFile[0], originalFile[1]

// frame1 := StringToBinary(originalName)

// frame2 := StringToBinary(originalExtension)

// wg.Add(1)

// go CreateFrames(imgProps, bufferSize-len(frame1), counter+1, frame1, &wg)

// wg.Add(1)

// go CreateFrames(imgProps, bufferSize-len(frame2), counter+2, frame2, &wg)
