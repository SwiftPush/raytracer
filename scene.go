package main

type Scene struct {
	objects HitableList
	camera  Camera
}

func exampleScene1(nx, ny int) Scene {
	objects := HitableList{
		[]Hitable{
			Sphere{Vector{0, 0, -1}, 0.5, Diffuse{Vector{0.8, 0.3, 0.3}}},
			Sphere{Vector{0, -100.5, -1}, 100, Diffuse{Vector{0.8, 0.8, 0}}},
			Sphere{Vector{1, 0, -1}, 0.5, Metal{Vector{0.8, 0.6, 0.2}, 0.3}},
			Sphere{Vector{-1, 0, -1}, 0.5, Dielectric{1.5}},
			Sphere{Vector{-1, 0, -1}, -0.45, Dielectric{1.5}},
		},
	}
	lookFrom := Vector{-2, 2, 1}
	lookAt := Vector{0, 0, -1}
	camera := NewCamera(
		lookFrom,
		lookAt,
		Vector{0, 1, 0},
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
