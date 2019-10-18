package main

import (
	"fmt"
	"math"
)

func color(ray Ray) Vector {
	t := hitSphere(Vector{0, 0, -1}, 0.5, ray)
	if t > 0 {
		n := ray.pointAtPatameter(t).Subtract(Vector{0, 0, -1}).Normalise()
		return n.AddScalar(1).MultiplyScalar(0.5)
	}
	unitDirection := ray.direction.Normalise()
	t = 0.5 * (unitDirection.y + 1.0)
	return Vector{1, 1, 1}.MultiplyScalar(1.0 - t).Add(Vector{0.5, 0.7, 1.0}.MultiplyScalar(t))
}

func hitSphere(center Vector, radius float64, ray Ray) float64 {
	oc := ray.origin.Subtract(center)
	a := ray.direction.Dot(ray.direction)
	b := 2.0 * oc.Dot(ray.direction)
	c := oc.Dot(oc) - radius*radius
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return -1.0
	}
	return (-b - math.Sqrt(discriminant)) / (2.0 * a)
}

func main() {
	nx, ny := 200, 100
	fmt.Printf("P3\n%v %v \n255\n", nx, ny)
	lowerLeftCorner := Vector{-2, -1, -1}
	horizontal := Vector{4, 0, 0}
	vertical := Vector{0, 2, 0}
	origin := Vector{0, 0, 0}
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)
			r := Ray{
				origin,
				lowerLeftCorner.Add(horizontal.MultiplyScalar(u).Add(vertical.MultiplyScalar(v))),
			}
			col := color(r)
			ir := int(255.99 * col.x)
			ig := int(255.99 * col.y)
			ib := int(255.99 * col.z)
			fmt.Printf("%v %v %v\n", ir, ig, ib)
		}
	}
}
