# Pixelify Your High-Res Images
Original                   |       Pixelated
:-------------------------:|:-------------------------:
<img src="./.readme/mark-harpur-K2s_YE031CA-unsplash.jpg" height="638" width="426"/>      |     <img src="./.readme/mark-harpur-K2s_YE031CA-unsplash_pixelated.jpg" height="638" width="426"/>

Pixelator is a fun and experimental project that intentionally degrades an image over several cycles.
Imagine your pristine image suffering through repeated JPEG compression and a bit of intentional randomization until it transforms into an abstract artwork. (pixel arts may be?)

## Features

- AddNoise(level int) ImageTransform
  Applies random Gaussian noise to an image. The noise level determines the maximum variation added to the pixel values.

- Downscale(scale float64) ImageTransform
  Downscales an image by the specified factor and then upscales it back to its original dimensions. This simulates data loss and creates a pixelated appearance.

- Blockify(size int) ImageTransform
  Reduces the resolution of the image by grouping pixels into blocks (based on the given size) and then scaling it back up, resulting in a blocky, pixelated effect.


## Usage

1. Clone the repository:
  ```bash
    git clone https://github.com/gokaybiz/pixelator.git
    cd pixelator
  ```
2. Run the application directly or build it:
  For direct execution, use:
  ```bash
    go run cmd/main.go
  ```
  To build the executable, run:
  ```bash
    go build -o pixelator ./cmd/main.go
    ./pixelator
  ```

## Contributing
Feel free to submit pull requests or report issues.

## License
This project is licensed under the GPL-3.0-or-later License.
