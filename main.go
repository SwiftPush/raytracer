package main

import "fmt"

func color(ray Ray) Vector {
	if hitSphere(Vector{0, 0, -1}, 0.5, ray) {
		return Vector{1, 0, 0}
	}
	unitDirection := ray.direction.Normalise()
	t := 0.5 * (unitDirection.y + 1.0)
	return Vector{1, 1, 1}.MultiplyScalar(1.0 - t).Add(Vector{0.5, 0.7, 1.0}.MultiplyScalar(t))
}

func hitSphere(center Vector, radius float64, ray Ray) bool {
	oc := ray.origin.Subtract(center)
	a := ray.direction.Dot(ray.direction)
	b := 2.0 * oc.Dot(ray.direction)
	c := oc.Dot(oc) - radius*radius
	discriminant := b*b - 4*a*c
	return discriminant > 0
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
