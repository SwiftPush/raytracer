package main

import "math/rand"

type Diffuse struct {
	albedo Vector
}

func (d Diffuse) scatter(ray Ray, hitRecord HitRecord, rnd *rand.Rand) (result bool, attenuation Vector, scattered Ray) {
	target := hitRecord.p.Add(hitRecord.normal).Add(randomInUnitSphere(rnd))

	result = true
	scattered = Ray{hitRecord.p, target.Subtract(hitRecord.p)}
	attenuation = d.albedo

	return
}
