package main

type Camera struct {
}

func (c Camera) getRay(u, v float64) Ray {
	lowerLeftCorner := Vector{-2, -1, -1}
	horizontal := Vector{4, 0, 0}
	vertical := Vector{0, 2, 0}
	origin := Vector{0, 0, 0}

	return Ray{
		origin,
		lowerLeftCorner.Add(horizontal.MultiplyScalar(u).Add(vertical.MultiplyScalar(v))),
	}
}
