package main

type Ray struct {
	origin, direction Vector
}

func (ray Ray) pointAtPatameter(t float64) Vector {
	return ray.origin.Add(ray.direction.MultiplyScalar(t))
}
