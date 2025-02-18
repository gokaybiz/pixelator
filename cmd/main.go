package main

import (
	"log"
	"os"

	"github.com/disintegration/imaging"
	"github.com/gokaybiz/pixelator/internal/effect"
	"github.com/gokaybiz/pixelator/internal/transform"
	"github.com/gokaybiz/pixelator/internal/util"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: pixelator <input_image>")
	}

	inputFileLocation := os.Args[1]
	outputFileLocation := util.GenerateOutputLocation(inputFileLocation)

	src, err := util.CheckFileAndLoad(inputFileLocation)
	if err != nil {
		log.Fatal(err)
	}

	srcSize := effect.Dimensions(src.Bounds().Size())
	fx := effect.Compute(srcSize) // Effect ratio adjustments based on resolution (width x height)

	log.Printf("Reduce Ratio: %v, Pixelation Size: %v, Noise Level: %v\n", fx.ScaleFactor, fx.BlockSize, fx.DistortionLevel)

	transform := transform.Pipeline(
		transform.Downscale(fx.ScaleFactor),
		transform.AddNoise(fx.DistortionLevel),
		transform.Blockify(fx.BlockSize),
		transform.AddNoise(fx.DistortionLevel/2),
		transform.Downscale(fx.ScaleFactor/2),
	)

	transformedImage := transform(src)
	transformedSize := effect.Dimensions(transformedImage.Bounds().Size())

	log.Printf("Source Image Size: %v x %v\n", srcSize.X, srcSize.Y)
	log.Printf("Generated Image Size: %v x %v\n", transformedSize.X, transformedSize.Y)

	if err := imaging.Save(transformedImage, outputFileLocation, imaging.JPEGQuality(68)); err != nil {
		log.Fatalf("Failed to save image: %v", err)
	}

	log.Println("But check the output:")
	log.Printf("Here you go => \"%v\"", outputFileLocation)
}
