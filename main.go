package main

import (
	"errors"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/cheggaaa/pb"
)

type ImageOptions struct {
	nX, nY int
	nS     int
}

func colour(ray Ray, objects HitableList, depth int, rnd *rand.Rand) Vector {
	if hit, hitRecord := objects.hit(ray, 0.001, math.MaxFloat64); hit {
		if depth >= 10 {
			return Vector{0, 0, 0}
		}
		if result, attenuation, scattered := hitRecord.material.scatter(ray, hitRecord, rnd); result {
			return attenuation.Multiply(colour(scattered, objects, depth+1, rnd))
		}
		return Vector{0, 0, 0}
	}
	unitDirection := ray.direction.Normalise()
	t := 0.5 * (unitDirection.y + 1.0)
	return Vector{1, 1, 1}.MultiplyScalar(1.0 - t).Add(Vector{0.5, 0.7, 1.0}.MultiplyScalar(t))
}

func renderPixel(scene Scene, rnd *rand.Rand, io ImageOptions, i, j int) color.RGBA {
	col := Vector{0, 0, 0}
	for s := 0; s < io.nS; s++ {
		u := (float64(i) + rnd.Float64()) / float64(io.nX)
		v := (float64(j) + rnd.Float64()) / float64(io.nY)
		ray := scene.camera.getRay(rnd, u, v)
		//p := ray.pointAtPatameter(2.0)
		col = col.Add(colour(ray, scene.objects, 0, rnd))
	}
	col = col.DivideScalar(float64(io.nS))
	// Gamera Correction
	/*col = Vector{
		math.Sqrt(col.x),
		math.Sqrt(col.y),
		math.Sqrt(col.z),
	}*/
	pixelColour := color.RGBA{
		R: uint8(255.99 * col.x),
		G: uint8(255.99 * col.y),
		B: uint8(255.99 * col.z),
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
		return errors.New("could not create output file")
	}
	defer file.Close()

	err = png.Encode(file, frameBuffer)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	io := ImageOptions{
		nX: 600, nY: 300,
		nS: 100,
	}
	scene := exampleScene1(io.nX, io.nY)
	frameBuffer := renderImage(io, scene)

	err := writeFrameToFile("out.png", frameBuffer)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
}
