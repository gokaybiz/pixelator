package transform

import (
	"image"

	"github.com/disintegration/imaging"
)

func Downscale(scale float64) ImageTransform {
	return func(img image.Image) image.Image {
		width := int(float64(img.Bounds().Dx()) * scale)
		height := int(float64(img.Bounds().Dy()) * scale)
		small := imaging.Resize(img, width, height, imaging.Lanczos)
		return imaging.Resize(small, img.Bounds().Dx(), img.Bounds().Dy(), imaging.NearestNeighbor)
	}
}

func Blockify(size int) ImageTransform {
	return func(img image.Image) image.Image {
		width := img.Bounds().Dx() / size
		height := img.Bounds().Dy() / size
		small := imaging.Resize(img, width, height, imaging.NearestNeighbor)
		return imaging.Resize(small, img.Bounds().Dx(), img.Bounds().Dy(), imaging.NearestNeighbor)
	}
}
