package transform

import (
	"image"
	"image/color"
	"math/rand"
	"sync"

	"github.com/gokaybiz/pixelator/internal/util"
)

// AddNoise returns an ImageTransform function that applies Gaussian noise to an image.
func AddNoise(level int) ImageTransform {
	return func(img image.Image) image.Image {
		bounds := img.Bounds()          // Size
		result := image.NewRGBA(bounds) // Empty canvas

		workers := 4
		height := bounds.Dy()
		chunkSize := height / workers

		var wg sync.WaitGroup
		wg.Add(workers)

		for i := 0; i < workers; i++ {
			startY := bounds.Min.Y + i*chunkSize
			endY := startY + chunkSize
			if i == workers-1 {
				endY = bounds.Max.Y
			}

			go _processChunk(img, result, startY, endY, bounds, level, &wg)
		}

		wg.Wait()
		return result
	}
}

// https://www.youtube.com/watch?v=-Vk23ye2o_I
// Each pixel value is increased by a random value up to the provided level.
func _processChunk(source image.Image, destination *image.RGBA, startY, endY int, bounds image.Rectangle, level int, wg *sync.WaitGroup) {
	defer wg.Done()
	for y := startY; y < endY; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := source.At(x, y).RGBA() // Read the color components (returned as 16-bit values)
			noise := uint32(rand.Intn(level))    // Generate noise value between 0 and level-1
			destination.Set(x, y, color.RGBA{
				// Add noise and clamp to the maximum value 65535 (16-bit limit).
				uint8((util.Clamp(r+noise, 0, 65535)) >> 8),
				uint8((util.Clamp(g+noise, 0, 65535)) >> 8),
				uint8((util.Clamp(b+noise, 0, 65535)) >> 8),
				uint8(a >> 8),
			})
		}
	}
}
