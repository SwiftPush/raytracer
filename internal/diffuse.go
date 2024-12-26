package internal

import (
	"math/rand"
	"raytracer/internal/geometry"
)

type Diffuse struct {
	albedo geometry.Vector
}

func (d Diffuse) scatter(ray Ray, hitRecord HitRecord, rnd *rand.Rand) (result bool, attenuation geometry.Vector, scattered Ray) {
	target := hitRecord.p.Add(hitRecord.normal).Add(randomInUnitSphere(rnd))

	result = true
	scattered = Ray{hitRecord.p, target.Subtract(hitRecord.p)}
	attenuation = d.albedo

	return
}
