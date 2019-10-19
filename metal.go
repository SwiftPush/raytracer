package main

type Metal struct {
	albedo Vector
	fuzz   float64
}

func reflect(v, n Vector) Vector {
	return v.Subtract(n.MultiplyScalar(v.Dot(n)).MultiplyScalar(2))
}

func (m Metal) scatter(ray Ray, hitRecord HitRecord) (bool, Vector, Ray) {
	reflected := reflect(ray.direction.Normalise(), hitRecord.normal)
	scattered := Ray{hitRecord.p, reflected.Add(randomInUnitSphere().MultiplyScalar(m.fuzz))}
	attenuation := m.albedo
	result := scattered.direction.Dot(hitRecord.normal) > 0
	return result, attenuation, scattered
}
