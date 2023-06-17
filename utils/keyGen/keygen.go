package keygen

import (
	"math/rand"

	"gocv.io/x/gocv"
)

type KeyGen struct {
}

func New() KeyGen {
	return KeyGen{}
}

func (k *KeyGen) Generate() []byte {
	// Open the video file
	video, err := gocv.VideoCaptureFile("assets/video/Screen_Recording.mp4")
	if err != nil {
		panic(err)
	}
	defer video.Close()

	frameCount := int(video.Get(gocv.VideoCaptureFrameCount))

	randomFrameNumber := rand.Intn(frameCount)

	video.Set(gocv.VideoCapturePosFrames, float64(randomFrameNumber))

	frame := gocv.NewMat()

	if ok := video.Read(&frame); !ok {
		panic("Failed to read the frame")
	}

	videoFrame, _ := gocv.IMEncode(".jpg", frame)

	frameBytes := videoFrame.GetBytes()
	last32Bytes := frameBytes[len(frameBytes)-32:]

	return last32Bytes
}
