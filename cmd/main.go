package main

import (
	"fmt"
	"log"
	"os"

	"github.com/disintegration/imaging"
	"github.com/gokaybiz/degrador/internal/effect"
	"github.com/gokaybiz/degrador/internal/transform"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: degrador <input_image>")
	}

	src, err := imaging.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Failed to load image: %v", err)
	}

	srcImgSize := src.Bounds().Size()
	fx := effect.Compute(srcImgSize) // Effect adjustments based on resolution

	fmt.Printf("Reduce Ratio: %v, Pixelation Size: %v, Noise Level: %v\nSource Image Size: %v\n", fx.ScaleFactor, fx.BlockSize, fx.DistortionLevel, srcImgSize)
	transform := transform.Pipeline(
		transform.AddNoise(fx.DistortionLevel),
		transform.Blockify(fx.BlockSize),
		transform.Downscale(fx.ScaleFactor),
	)

	transformedImage := transform(src)

	if err := imaging.Save(transformedImage, "result.jpg", imaging.JPEGQuality(50)); err != nil {
		log.Fatalf("Failed to save image: %v", err)
	}
}
