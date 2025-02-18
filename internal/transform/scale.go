package transform

import (
	"image"
	"image/color"
	"sync"

	"github.com/disintegration/imaging"
)

// Downscale returns an ImageTransform that shrinks the image by a factor of 'scale',
// then upscales it back to its original dimensions. This simulates a loss of detail. But not rough as Blockify's
func Downscale(scale float64) ImageTransform {
	return func(img image.Image) image.Image {
		width := int(float64(img.Bounds().Dx()) * scale)
		height := int(float64(img.Bounds().Dy()) * scale)
		small := imaging.Resize(img, width, height, imaging.Lanczos) // Downscale using Lanczos filter for quality.
		return imaging.Resize(small, img.Bounds().Dx(), img.Bounds().Dy(), imaging.NearestNeighbor)
	}
}

// Blockify returns an ImageTransform that reduces the image resolution by dividing it into blocks,
// then upscales it back to its original size to give a pixelated pixelated look.
func Blockify(blockSize int) ImageTransform {
	return func(img image.Image) image.Image {
		bounds := img.Bounds()
		width, height := bounds.Dx(), bounds.Dy()

		// Let's compute block counts.
		blocksX := width / blockSize
		if width%blockSize != 0 {
			blocksX++
		}
		blocksY := height / blockSize
		if height%blockSize != 0 {
			blocksY++
		}

		// A smaller image where each pixel is one block.
		small := image.NewRGBA(image.Rect(0, 0, blocksX, blocksY))

		var wg sync.WaitGroup

		// Loop over blocks in parallel.
		for by := 0; by < blocksY; by++ {
			by := by // capture loop variable
			wg.Add(1)
			go func() {
				defer wg.Done()
				for bx := 0; bx < blocksX; bx++ {
					// Boundaries of the block.
					startX := bounds.Min.X + bx*blockSize
					startY := bounds.Min.Y + by*blockSize

					// End coordinates.
					endX := startX + blockSize
					if endX > bounds.Max.X {
						endX = bounds.Max.X
					}
					endY := startY + blockSize
					if endY > bounds.Max.Y {
						endY = bounds.Max.Y
					}

					// Accumulate values to average calculation.
					var sumR, sumG, sumB, sumA uint32
					var count uint32

					for y := startY; y < endY; y++ {
						for x := startX; x < endX; x++ {
							r, g, b, a := img.At(x, y).RGBA()
							sumR += r
							sumG += g
							sumB += b
							sumA += a
							count++
						}
					}

					// every block has at least one pixel, so...
					avgR := uint8((sumR / count) >> 8)
					avgG := uint8((sumG / count) >> 8)
					avgB := uint8((sumB / count) >> 8)
					avgA := uint8((sumA / count) >> 8)

					// Set the computed block to the small image.
					small.Set(bx, by, color.NRGBA{R: avgR, G: avgG, B: avgB, A: avgA})
				}
			}()
		}
		wg.Wait()

		// Upscale the small block image back to original size using nearest neighbor.
		return imaging.Resize(small, width, height, imaging.NearestNeighbor)
	}
}
