package internal

import (
	"math"
	"math/rand"

	"raytracer/internal/geometry"
)

type Camera struct {
	origin, vertical, horizontal, lowerLeftCorner, u, v geometry.Vector
	lensRadius                                          float64
}

func NewCamera(lookFrom, lookAt, vup geometry.Vector, verticalFov, aspect, aperture, focusDist float64) Camera {
	theta := verticalFov * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight
	w := lookFrom.Subtract(lookAt).Normalise()
	u := vup.Cross(w).Normalise()
	v := w.Cross(u)
	lowerLeftCorner := lookFrom.
		Subtract(u.MultiplyScalar(halfWidth * focusDist)).
		Subtract(v.MultiplyScalar(halfHeight * focusDist)).
		Subtract(w.MultiplyScalar(focusDist))
	return Camera{
		origin:          lookFrom,
		vertical:        v.MultiplyScalar(halfHeight * 2 * focusDist),
		horizontal:      u.MultiplyScalar(halfWidth * 2 * focusDist),
		lowerLeftCorner: lowerLeftCorner,
		lensRadius:      aperture / 2,
		u:               u,
		v:               v,
	}
}

func (c Camera) getRay(rnd *rand.Rand, s, t float64) Ray {
	rd := randomInUnitDisk(rnd).MultiplyScalar(c.lensRadius)
	offset := c.u.MultiplyScalar(rd.X).Add(c.v.MultiplyScalar(rd.Y))
	return Ray{
		origin:    c.origin.Add(offset),
		direction: c.lowerLeftCorner.Add(c.horizontal.MultiplyScalar(s).Add(c.vertical.MultiplyScalar(t))).Subtract(c.origin).Subtract(offset),
	}
}

func randomInUnitDisk(rnd *rand.Rand) geometry.Vector {
	p := geometry.Vector{X: rnd.Float64(), Y: rnd.Float64(), Z: 0}.MultiplyScalar(2).Subtract(geometry.Vector{X: 1, Y: 1, Z: 0})
	for p.Dot(p) >= 1.0 {
		p = geometry.Vector{X: rnd.Float64(), Y: rnd.Float64(), Z: 0}.MultiplyScalar(2).Subtract(geometry.Vector{X: 1, Y: 1, Z: 0})
	}
	return p
}
