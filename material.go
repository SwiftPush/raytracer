package main

type Material interface {
	scatter(Ray, HitRecord) (bool, Vector, Ray)
}
