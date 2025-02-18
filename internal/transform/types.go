package transform

import "image"

type ImageTransform = func(image.Image) image.Image
