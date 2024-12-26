package internal

type Hitable interface {
	hit(Ray, float64, float64) (bool, HitRecord)
}
