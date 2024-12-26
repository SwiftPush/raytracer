package internal

import "raytracer/internal/geometry"

type HitRecord struct {
	t         float64
	p, normal geometry.Vector
	material  Material
}
