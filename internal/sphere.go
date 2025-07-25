package internal

import (
	"math"
	"math/rand"
	"raytracer/internal/geometry"
)

type Sphere struct {
	center   geometry.Vector
	radius   float64
	material Material
}

func (sphere Sphere) hit(ray Ray, tMin float64, tMax float64) (bool, HitRecord) {
	oc := ray.origin.Subtract(sphere.center)
	a := ray.direction.Dot(ray.direction)
	b := oc.Dot(ray.direction)
	c := oc.Dot(oc) - sphere.radius*sphere.radius
	discriminant := b*b - a*c
	if discriminant > 0 {
		temp := (-b - math.Sqrt(b*b-a*c)) / a
		if (temp < tMax) && (temp > tMin) {
			hr := HitRecord{}
			hr.material = sphere.material
			hr.t = temp
			hr.p = ray.pointAtParameter(hr.t)
			hr.normal = hr.p.Subtract(sphere.center).DivideScalar(sphere.radius)
			return true, hr
		}
		temp = (-b + math.Sqrt(b*b-a*c)) / a
		if (temp < tMax) && (temp > tMin) {
			hr := HitRecord{}
			hr.material = sphere.material
			hr.t = temp
			hr.p = ray.pointAtParameter(hr.t)
			hr.normal = hr.p.Subtract(sphere.center).DivideScalar(sphere.radius)
			return true, hr
		}
	}
	return false, HitRecord{}
}

func randomInUnitSphere(rnd *rand.Rand) geometry.Vector {
	v := geometry.Vector{X: rnd.Float64(), Y: rnd.Float64(), Z: rnd.Float64()}
	for v.SquaredLength() > 1 {
		v = geometry.Vector{X: rnd.Float64(), Y: rnd.Float64(), Z: rnd.Float64()}
	}
	return v
}
