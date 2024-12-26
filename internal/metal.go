package internal

import (
	"math/rand"
	"raytracer/internal/geometry"
)

type Metal struct {
	albedo geometry.Vector
	fuzz   float64
}

func reflect(v, n geometry.Vector) geometry.Vector {
	return v.Subtract(n.MultiplyScalar(v.Dot(n)).MultiplyScalar(2))
}

func (m Metal) scatter(ray Ray, hitRecord HitRecord, rnd *rand.Rand) (result bool, attenuation geometry.Vector, scattered Ray) {
	reflected := reflect(ray.direction.Normalise(), hitRecord.normal)

	scattered = Ray{hitRecord.p, reflected.Add(randomInUnitSphere(rnd).MultiplyScalar(m.fuzz))}
	attenuation = m.albedo
	result = scattered.direction.Dot(hitRecord.normal) > 0

	return
}
