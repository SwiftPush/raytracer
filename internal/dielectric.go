package internal

import (
	"math"
	"math/rand"
	"raytracer/internal/geometry"
)

type Dielectric struct {
	refIdx float64
}

func schlick(cosine, refIdx float64) float64 {
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}

func refract(v, n geometry.Vector, niOverNt float64) (result bool, refracted geometry.Vector) {
	uv := v.Normalise()
	dt := uv.Dot(n)
	discriminant := 1.0 - niOverNt*niOverNt*(1-dt*dt)
	if discriminant > 0 {
		refracted = uv.Subtract(n.MultiplyScalar(dt)).MultiplyScalar(niOverNt).Subtract(n.MultiplyScalar(math.Sqrt(discriminant)))
		return true, refracted
	}
	return
}

func (d Dielectric) scatter(ray Ray, hitRecord HitRecord, rnd *rand.Rand) (result bool, attenuation geometry.Vector, scattered Ray) {
	attenuation = geometry.Vector{X: 1.0, Y: 1.0, Z: 1.0}

	var outwardNormal geometry.Vector
	niOverNt := 0.0
	cosine := 0.0
	if ray.direction.Dot(hitRecord.normal) > 0 {
		outwardNormal = hitRecord.normal.MultiplyScalar(-1)
		niOverNt = d.refIdx
		cosine = d.refIdx * ray.direction.Dot(hitRecord.normal) / ray.direction.Length()
	} else {
		outwardNormal = hitRecord.normal
		niOverNt = 1.0 / d.refIdx
		cosine = (-1 * ray.direction.Dot(hitRecord.normal)) / ray.direction.Length()
	}

	reflected := reflect(ray.direction, hitRecord.normal)

	if doesRefract, refracted := refract(ray.direction, outwardNormal, niOverNt); doesRefract {
		reflectProbability := schlick(cosine, d.refIdx)
		if rnd.Float64() < reflectProbability {
			scattered = Ray{hitRecord.p, reflected}
		} else {
			scattered = Ray{hitRecord.p, refracted}
		}
	} else {
		scattered = Ray{hitRecord.p, reflected}
	}

	return true, attenuation, scattered
}
