package main

import (
	"math"
	"math/rand"
)

type Camera struct {
	origin, vertical, horizontal, lowerLeftCorner, u, v Vector
	lensRadius                                          float64
}

func NewCamera(lookFrom, lookAt, vup Vector, verticalFov, aspect, aperture, focusDist float64) Camera {
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
	offset := c.u.MultiplyScalar(rd.x).Add(c.v.MultiplyScalar(rd.y))
	return Ray{
		origin:    c.origin.Add(offset),
		direction: c.lowerLeftCorner.Add(c.horizontal.MultiplyScalar(s).Add(c.vertical.MultiplyScalar(t))).Subtract(c.origin).Subtract(offset),
	}
}

func randomInUnitDisk(rnd *rand.Rand) Vector {
	p := Vector{rnd.Float64(), rnd.Float64(), 0}.MultiplyScalar(2).Subtract(Vector{1, 1, 0})
	for p.Dot(p) >= 1.0 {
		p = Vector{rnd.Float64(), rnd.Float64(), 0}.MultiplyScalar(2).Subtract(Vector{1, 1, 0})
	}
	return p
}
