package encoderFile

import (
	"fmt"
	"os/exec"
)

func ConvertFramesToVideo(fps int, tmpDir string, outputVideoFile string) {
	println("Creating video... ")

	imagePattern := fmt.Sprintf("%s/frame%s.png", tmpDir, "%04d")
	fmt.Println(imagePattern)

	frames := fmt.Sprintf("%d", fps)

	cmd := exec.Command("ffmpeg", "-framerate", frames, "-i", imagePattern, "-c:v", "av1_qsv", "-pix_fmt", "nv12", outputVideoFile)
	fmt.Println(cmd.Args)

	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	fmt.Println("Video created:", outputVideoFile)
}
