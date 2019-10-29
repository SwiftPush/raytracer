package main

type Diffuse struct {
	albedo Vector
}

func (d Diffuse) scatter(ray Ray, hitRecord HitRecord) (result bool, attenuation Vector, scattered Ray) {
	target := hitRecord.p.Add(hitRecord.normal).Add(randomInUnitSphere())

	result = true
	scattered = Ray{hitRecord.p, target.Subtract(hitRecord.p)}
	attenuation = d.albedo

	return
}
