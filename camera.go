package main

import "math"

type Camera struct {
	origin, vertical, horizontal, lowerLeftCorner Vector
}

func NewCamera(lookFrom, lookAt, vup Vector, verticalFov, aspect float64) Camera {
	theta := verticalFov * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight
	w := lookFrom.Subtract(lookAt).Normalise()
	u := vup.Cross(w).Normalise()
	v := w.Cross(u)
	lowerLeftCorner := lookFrom.
		Subtract(u.MultiplyScalar(halfWidth)).
		Subtract(v.MultiplyScalar(halfHeight)).
		Subtract(w)
	return Camera{
		origin:          lookFrom,
		vertical:        v.MultiplyScalar(halfHeight * 2),
		horizontal:      u.MultiplyScalar(halfWidth * 2),
		lowerLeftCorner: lowerLeftCorner,
	}
}

func (c Camera) getRay(u, v float64) Ray {
	return Ray{
		origin:    c.origin,
		direction: c.lowerLeftCorner.Add(c.horizontal.MultiplyScalar(u).Add(c.vertical.MultiplyScalar(v))).Subtract(c.origin),
	}
}
