package main

import (
	"fmt"
	"os"

	"github.com/dailaim/Go-Glish/encoderFile"
)

const (
	framesDir       = "./test/frames"
	outputVideoFile = "./test/video_salida2.mp4"
)

func main() {
	inputData := getInput("convert file to video press 1\nconvert video to file press 2\ndownload video from YouTube press 3\n")

	imgProps := &encoderFile.ImageProperties{
		Width:     1920,
		Height:    1080,
		PixelSize: 12,
		TmpDir:    framesDir,
	}

	switch inputData {
	case "1":
		filename := "./test/picxReal_10.safetensors" // getInput("Enter the name of the file to convert: ")

		os.RemoveAll(framesDir)
		os.Remove(outputVideoFile)

		encoderFile.ProcessFileToFrames(filename, framesDir, imgProps)

		fps := 20

		encoderFile.ConvertFramesToVideo(fps, framesDir, outputVideoFile)

		// os.RemoveAll(framesDir)
	case "2":
		fmt.Println("Not implemented yet.")
	case "3":
		fmt.Println("Not implemented yet.")
	default:
		fmt.Println("Invalid input.")
	}
}

func getInput(message string) string {
	var input string
	fmt.Print(message)
	fmt.Scanln(&input)
	return input
}
