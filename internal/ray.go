package internal

import "raytracer/internal/geometry"

type Ray struct {
	origin, direction geometry.Vector
}

func (ray Ray) pointAtParameter(t float64) geometry.Vector {
	return ray.origin.Add(ray.direction.MultiplyScalar(t))
}
