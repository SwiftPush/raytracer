package internal

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
	"raytracer/internal/geometry"
	"sync"
	"time"

	"github.com/cheggaaa/pb"
)

type ImageOptions struct {
	nX, nY int
	nS     int
}

func colour(ray Ray, objects HitableList, depth int, rnd *rand.Rand) geometry.Vector {
	if hit, hitRecord := objects.hit(ray, 0.001, math.MaxFloat64); hit {
		if depth >= 10 {
			return geometry.Vector{X: 0, Y: 0, Z: 0}
		}
		if result, attenuation, scattered := hitRecord.material.scatter(ray, hitRecord, rnd); result {
			return attenuation.Multiply(colour(scattered, objects, depth+1, rnd))
		}
		return geometry.Vector{X: 0, Y: 0, Z: 0}
	}
	unitDirection := ray.direction.Normalise()
	t := 0.5 * (unitDirection.Y + 1.0)
	return geometry.Vector{X: 1, Y: 1, Z: 1}.MultiplyScalar(1.0 - t).Add(geometry.Vector{X: 0.5, Y: 0.7, Z: 1.0}.MultiplyScalar(t))
}

func renderPixel(scene Scene, rnd *rand.Rand, io ImageOptions, i, j int) color.RGBA {
	col := geometry.Vector{X: 0, Y: 0, Z: 0}
	for s := 0; s < io.nS; s++ {
		u := (float64(i) + rnd.Float64()) / float64(io.nX)
		v := (float64(j) + rnd.Float64()) / float64(io.nY)
		ray := scene.camera.getRay(rnd, u, v)
		//p := ray.pointAtPatameter(2.0)
		col = col.Add(colour(ray, scene.objects, 0, rnd))
	}
	col = col.DivideScalar(float64(io.nS))
	// Gamera Correction
	/*col = geometry.Vector{
		math.Sqrt(col.x),
		math.Sqrt(col.y),
		math.Sqrt(col.z),
	}*/
	pixelColour := color.RGBA{
		R: uint8(255.99 * col.X),
		G: uint8(255.99 * col.Y),
		B: uint8(255.99 * col.Z),
		A: 255,
	}
	return pixelColour
}

func renderImage(io ImageOptions, scene Scene) *image.RGBA {
	frameBuffer := image.NewRGBA(image.Rect(0, 0, io.nX, io.nY))

	bar := pb.StartNew(io.nY)
	var wg sync.WaitGroup
	wg.Add(io.nY)
	for j := 0; j < io.nY; j++ {
		go func(j int) {
			defer wg.Done()
			rnd := rand.New(rand.NewSource(time.Now().Unix() + int64(j)))
			for i := 0; i < io.nX; i++ {
				pixelColour := renderPixel(scene, rnd, io, i, j)
				frameBuffer.SetRGBA(i, io.nY-j, pixelColour)
			}
			bar.Increment()
		}(j)
	}
	wg.Wait()
	bar.Finish()

	return frameBuffer
}

func writeFrameToFile(filename string, frameBuffer *image.RGBA) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("could not create output file: %w", err)
	}
	defer file.Close()

	err = png.Encode(file, frameBuffer)
	if err != nil {
		return err
	}

	return nil
}

func Render() error {
	io := ImageOptions{
		nX: 600, nY: 300,
		nS: 100,
	}
	scene := exampleScene1(io.nX, io.nY)
	frameBuffer := renderImage(io, scene)

	err := writeFrameToFile("out.png", frameBuffer)
	if err != nil {
		return err
	}

	return nil
}
