package internal

import "raytracer/internal/geometry"

type Scene struct {
	objects HitableList
	camera  Camera
}

func exampleScene1(nx, ny int) Scene {
	objects := HitableList{
		[]Hitable{
			Sphere{geometry.Vector{X: 0, Y: 0, Z: -1}, 0.5, Diffuse{geometry.Vector{X: 0.8, Y: 0.3, Z: 0.3}}},
			Sphere{geometry.Vector{X: 0, Y: -100.5, Z: -1}, 100, Diffuse{geometry.Vector{X: 0.8, Y: 0.8, Z: 0}}},
			Sphere{geometry.Vector{X: 1, Y: 0, Z: -1}, 0.5, Metal{geometry.Vector{X: 0.8, Y: 0.6, Z: 0.2}, 0.3}},
			Sphere{geometry.Vector{X: -1, Y: 0, Z: -1}, 0.5, Dielectric{1.5}},
			Sphere{geometry.Vector{X: -1, Y: 0, Z: -1}, -0.45, Dielectric{1.5}},
		},
	}
	lookFrom := geometry.Vector{X: -2, Y: 2, Z: 1}
	lookAt := geometry.Vector{X: 0, Y: 0, Z: -1}
	camera := NewCamera(
		lookFrom,
		lookAt,
		geometry.Vector{X: 0, Y: 1, Z: 0},
		20,
		float64(nx)/float64(ny),
		1.0,
		lookFrom.Subtract(lookAt).Length(),
	)
	return Scene{
		objects: objects,
		camera:  camera,
	}
}
