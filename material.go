package main

import "math/rand"

type Material interface {
	scatter(Ray, HitRecord, *rand.Rand) (bool, Vector, Ray)
}
