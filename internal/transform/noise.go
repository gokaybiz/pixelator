package transform

import (
	"image"
	"image/color"
	"math/rand"
	"sync"

	"github.com/gokaybiz/degrador/internal/util"
)

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
// Guassian blur
func _processChunk(source image.Image, destination *image.RGBA, startY, endY int, bounds image.Rectangle, level int, wg *sync.WaitGroup) {
	defer wg.Done()
	for y := startY; y < endY; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := source.At(x, y).RGBA()
			noise := uint32(rand.Intn(level))
			destination.Set(x, y, color.RGBA{
				// 65535: maximum value of unsigned 16-bit integer
				uint8((util.Clamp(r+noise, 0, 65535)) >> 8),
				uint8((util.Clamp(g+noise, 0, 65535)) >> 8),
				uint8((util.Clamp(b+noise, 0, 65535)) >> 8),
				uint8(a >> 8),
			})
		}
	}
}
