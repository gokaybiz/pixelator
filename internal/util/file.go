package util

import (
	"errors"
	"fmt"
	"image"
	"os"
	"path"
	"strings"

	"github.com/disintegration/imaging"
)

func CheckFileAndLoad(input string) (image.Image, error) {
	// Check if input file exists
	if _, err := os.Stat(input); err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to find image: %v", err))
	}

	src, err := imaging.Open(input)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to load image: %v", err))
	}

	return src, nil
}

func GenerateOutputLocation(input string) string {
	inputDir, inputFile := path.Split(input)
	ext := path.Ext(inputFile)
	outputFileName := strings.TrimSuffix(inputFile, ext) + "_pixelated" + ext

	return path.Join(inputDir, outputFileName)
}
