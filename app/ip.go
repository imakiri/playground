package app

import (
	"github.com/imakiri/playground/core"
	"gocv.io/x/gocv"
	"image"
	"image/color"
)

func Detect(rImg []byte) (reImg []byte, err error) {
	img, err := gocv.IMDecode(rImg, gocv.IMReadColor)
	if err != nil {
		return nil, core.NewStatus(core.AppDetecterIncorrectImageError{}, err)
	}
	defer img.Close()

	imgGrey := gocv.NewMat()
	defer imgGrey.Close()

	gocv.CvtColor(img, &imgGrey, gocv.ColorBGRToGray)

	harrcascade := "app\\haarcascades\\haarcascade_frontalface_alt.xml"
	classifier := gocv.NewCascadeClassifier()
	classifier.Load(harrcascade)
	defer classifier.Close()

	//rects := classifier.DetectMultiScaleWithParams(img, 1.5, 2, 0, image.Point{}, image.Point{X: img.Size()[0], Y: img.Size()[1]})
	rects := classifier.DetectMultiScale(img)

	for _, r := range rects {
		gocv.Rectangle(&img, r.Add(image.Point{
			X: 1,
			Y: 1,
		}), color.RGBA{B: 255, G: 255}, 1)
		gocv.Rectangle(&img, r, color.RGBA{R: 255}, 1)
	}

	data, err := gocv.IMEncode(".jpg", img)
	if err != nil {
		return nil, err
	}

	return data, nil
}
