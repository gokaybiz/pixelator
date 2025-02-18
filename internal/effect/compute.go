package effect

import "github.com/gokaybiz/pixelator/internal/util"

// Calculate reduction ratio, pixelation size, and noise level dynamically
func Compute(dim Dimensions) ImageEffect {
	minEffect := ImageEffect{0.74, 8, 20}
	maxEffect := ImageEffect{0.44, 45, 85}

	minDim, maxDim := 1000.0, 3000.0
	currentDim := float64(util.Max(dim.X, dim.Y))

	if currentDim <= minDim {
		return minEffect
	}
	if currentDim >= maxDim {
		return maxEffect
	}

	// Linear interpolation factor (scales between 0 and 1)
	scale := (maxDim - currentDim) / (maxDim - minDim)

	// Interpolating each parameter
	return ImageEffect{
		ScaleFactor:     minEffect.ScaleFactor + scale*(maxEffect.ScaleFactor-minEffect.ScaleFactor),
		BlockSize:       int(float64(minEffect.BlockSize) + scale*float64(maxEffect.BlockSize-minEffect.BlockSize)),
		DistortionLevel: int(float64(minEffect.DistortionLevel) + scale*float64(maxEffect.DistortionLevel-minEffect.DistortionLevel)),
	}
}
