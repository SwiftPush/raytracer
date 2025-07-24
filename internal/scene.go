package internal

import "raytracer/internal/geometry"

type Scene struct {
	objects HitableList
	camera  Camera
}

func exampleScene1(nx, ny int) Scene {
	objects := HitableList{
		[]Hitable{
			Sphere{geometry.Vector{X: 0, Y: 0, Z: -1}, 0.5, Diffuse{geometry.Vector{X: 0.8, Y: 0.15, Z: 0.15}}},    // middle sphere
			Sphere{geometry.Vector{X: 0, Y: -100.5, Z: -1}, 100, Diffuse{geometry.Vector{X: 0.8, Y: 0.8, Z: 0.8}}}, // floor
			Sphere{geometry.Vector{X: 1.1, Y: 0, Z: -1}, 0.5, Metal{geometry.Vector{X: 0.8, Y: 0.6, Z: 0.2}, 0.3}}, // right sphere
			Sphere{geometry.Vector{X: -1, Y: 0, Z: -1}, 0.5, Dielectric{1.5}},                                      // glass sphere (1/2)
			Sphere{geometry.Vector{X: -1, Y: 0, Z: -1}, -0.45, Dielectric{1.5}},                                    // glass sphere (2/2)
		},
	}
	lookFrom := geometry.Vector{X: -3, Y: 1, Z: 4}
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
