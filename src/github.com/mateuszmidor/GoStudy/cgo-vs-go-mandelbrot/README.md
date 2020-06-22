# go vs cgo mandelbrot implementation performance

## Note

The speedup is not linear to number of goroutines as mandelbrot requires different amount of work depending on what part of image is being processed