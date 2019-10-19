package main

type Diffuse struct {
	albedo Vector
}

func NewDiffuse(albedo Vector) Diffuse {
	return Diffuse{
		albedo,
	}
}

func (d Diffuse) scatter(ray Ray, hitRecord HitRecord) (bool, Vector, Ray) {
	target := hitRecord.p.Add(hitRecord.normal).Add(randomInUnitSphere())
	scattered := Ray{hitRecord.p, target.Subtract(hitRecord.p)}
	attenuation := d.albedo
	return true, attenuation, scattered
}
