package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
)

func colour(ray Ray, world HitableList, depth int) Vector {
	MAXFLOAT := 999999999.9
	if hit, hitRecord := world.hit(ray, 0.001, MAXFLOAT); hit {
		if depth >= 50 {
			return Vector{0, 0, 0}
		}
		if result, attenuation, scattered := hitRecord.material.scatter(ray, hitRecord); result {
			return attenuation.Multiply(colour(scattered, world, depth+1))
		}
		return Vector{0, 0, 0}
	}
	unitDirection := ray.direction.Normalise()
	t := 0.5 * (unitDirection.y + 1.0)
	return Vector{1, 1, 1}.MultiplyScalar(1.0 - t).Add(Vector{0.5, 0.7, 1.0}.MultiplyScalar(t))
}

func main() {
	nx, ny := 200, 100
	ns := 100
	image := image.NewRGBA(image.Rect(0, 0, nx, ny))
	world := HitableList{
		[]Hitable{
			Sphere{Vector{0, 0, -1}, 0.5, Diffuse{Vector{0.8, 0.3, 0.3}}},
			Sphere{Vector{0, -100.5, -1}, 100, Diffuse{Vector{0.8, 0.8, 0}}},
			Sphere{Vector{1, 0, -1}, 0.5, Metal{Vector{0.8, 0.6, 0.2}, 0.3}},
			Sphere{Vector{-1, 0, -1}, 0.5, Metal{Vector{0.8, 0.8, 0.8}, 1.0}},
		},
	}
	camera := Camera{}
	for j := 0; j < ny; j++ {
		for i := 0; i < nx; i++ {
			col := Vector{0, 0, 0}
			for s := 0; s < ns; s++ {
				u := (float64(i) + rand.Float64()) / float64(nx)
				v := (float64(j) + rand.Float64()) / float64(ny)
				ray := camera.getRay(u, v)
				//p := ray.pointAtPatameter(2.0)
				col = col.Add(colour(ray, world, 0))
			}
			col = col.DivideScalar(float64(ns))
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
			image.SetRGBA(i, ny-j, pixelColour)
		}
	}
	file, err := os.Create("out.png")
	defer file.Close()
	if err != nil {
		log.Fatalf("Could not create output file\n")
		os.Exit(1)
	}
	png.Encode(file, image)
}
