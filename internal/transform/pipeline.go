package transform

import "image"

func Pipeline(transforms ...ImageTransform) ImageTransform {
    return func(img image.Image) image.Image {
        for _, t := range transforms {
            img = t(img)
        }
        return img
    }
}
