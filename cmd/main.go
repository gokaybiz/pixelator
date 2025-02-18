package main

import (
	"fmt"
	"log"
	"os"

	"github.com/disintegration/imaging"
	"github.com/gokaybiz/pixelator/internal/effect"
	"github.com/gokaybiz/pixelator/internal/transform"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: pixelator <input_image>")
	}

	src, err := imaging.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Failed to load image: %v", err)
	}

	srcSize := effect.Dimensions(src.Bounds().Size())
	fx := effect.Compute(srcSize) // Effect adjustments based on resolution (width x height)

	fmt.Printf("Reduce Ratio: %v, Pixelation Size: %v, Noise Level: %v\n", fx.ScaleFactor, fx.BlockSize, fx.DistortionLevel)

	transform := transform.Pipeline(
		transform.Downscale(fx.ScaleFactor),
		transform.AddNoise(fx.DistortionLevel),
		transform.Blockify(fx.BlockSize),
		transform.AddNoise(fx.DistortionLevel/2),
	)

	transformedImage := transform(src)
	transformedSize := effect.Dimensions(transformedImage.Bounds().Size())

	fmt.Printf("Source Image Size: %v x %v\n", srcSize.X, srcSize.Y)
	fmt.Printf("Generated Image Size: %v x %v\n", transformedSize.X, transformedSize.Y)

	if err := imaging.Save(transformedImage, "result.jpg", imaging.JPEGQuality(50)); err != nil {
		log.Fatalf("Failed to save image: %v", err)
	}
}
