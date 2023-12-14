package src

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"sync"
)

func MainMulti() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	wg := sync.WaitGroup{}
	wg.Add(width * height)
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Launch a goroutine to compute the pixel.
			go func(px, py int, z complex128) {
				// Image point (px, py) represents complex value z.
				img.Set(px, py, (func(z complex128) color.Color {
					var v complex128
					for n := uint8(0); n < iterations; n++ {
						v = v*v + z
						if cmplx.Abs(v) > 2 {
							return color.Gray{255 - contrast*n}
						}
					}
					return color.Black
				})(z))
				wg.Done()
			}(px, py, z)
		}
	}
	wg.Wait()

	file, _ := os.Create("mandelbrot-multi.png")
	png.Encode(file, img)
	file.Close()
}
