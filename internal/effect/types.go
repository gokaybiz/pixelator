package effect

import "image"

type ImageEffect struct {
	ScaleFactor     float64
	BlockSize       int
	DistortionLevel int
}

type Dimensions = image.Point
