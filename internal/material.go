package internal

import (
	"math/rand"
	"raytracer/internal/geometry"
)

type Material interface {
	scatter(Ray, HitRecord, *rand.Rand) (bool, geometry.Vector, Ray)
}
