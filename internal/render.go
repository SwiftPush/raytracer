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
	"runtime"
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
	return renderPixelAdaptive(scene, rnd, io, i, j)
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

func Render(outputFilename string) error {
	// Record start time
	startTime := time.Now()
	
	io := ImageOptions{
		nX: 600, nY: 300,
		nS: 300,
	}
	scene := exampleScene1(io.nX, io.nY)
	
	fmt.Printf("Rendering %dx%d image with up to %d samples per pixel...\n", io.nX, io.nY, io.nS)
	fmt.Printf("Using %d CPU cores\n", runtime.NumCPU())
	
	frameBuffer := renderImage(io, scene)
	
	// Calculate timing
	wallTime := time.Since(startTime)
	
	err := writeFrameToFile(outputFilename, frameBuffer)
	if err != nil {
		return err
	}
	
	// Calculate some interesting stats
	totalPixels := int64(io.nX * io.nY)
	maxPossibleSamples := totalPixels * int64(io.nS)
	
	// Print timing statistics
	fmt.Printf("\n=== Render Statistics ===\n")
	fmt.Printf("Wall time:        %s\n", formatDuration(wallTime))
	fmt.Printf("Estimated CPU:    %s (%.1fx parallelism)\n", 
		formatDuration(time.Duration(float64(wallTime)*float64(runtime.NumCPU())*0.8)),
		float64(runtime.NumCPU())*0.8)
	fmt.Printf("Total pixels:     %d\n", totalPixels)
	fmt.Printf("Max samples:      %d per pixel (%d total possible)\n", io.nS, maxPossibleSamples)
	fmt.Printf("Performance:      %.0f pixels/second\n", float64(totalPixels)/wallTime.Seconds())
	fmt.Printf("Throughput:       %.1f megasamples/second\n", float64(maxPossibleSamples)/wallTime.Seconds()/1e6)
	fmt.Printf("Output file:      %s\n", outputFilename)

	return nil
}

func formatDuration(d time.Duration) string {
	if d >= time.Minute {
		return fmt.Sprintf("%.2gm", d.Minutes())
	} else if d >= time.Second {
		return fmt.Sprintf("%.2gs", d.Seconds())
	} else if d >= time.Millisecond {
		return fmt.Sprintf("%.2gms", float64(d.Nanoseconds())/1e6)
	} else {
		return fmt.Sprintf("%.2gμs", float64(d.Nanoseconds())/1e3)
	}
}

