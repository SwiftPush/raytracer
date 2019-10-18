package main

import (
	"fmt"
)

func color(ray Ray, world HitableList) Vector {
	MAXFLOAT := 999999999.9
	if hit, hitRecord := world.hit(ray, 0.0, MAXFLOAT); hit {
		return hitRecord.normal.AddScalar(1).MultiplyScalar(0.5)
	}
	unitDirection := ray.direction.Normalise()
	t := 0.5 * (unitDirection.y + 1.0)
	return Vector{1, 1, 1}.MultiplyScalar(1.0 - t).Add(Vector{0.5, 0.7, 1.0}.MultiplyScalar(t))
}

func main() {
	nx, ny := 200, 100
	fmt.Printf("P3\n%v %v \n255\n", nx, ny)
	lowerLeftCorner := Vector{-2, -1, -1}
	horizontal := Vector{4, 0, 0}
	vertical := Vector{0, 2, 0}
	origin := Vector{0, 0, 0}
	world := HitableList{
		[]Hitable{
			Sphere{Vector{0, 0, -1}, 0.5},
			Sphere{Vector{0, -100.5, -1}, 100},
		},
	}
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)
			r := Ray{
				origin,
				lowerLeftCorner.Add(horizontal.MultiplyScalar(u).Add(vertical.MultiplyScalar(v))),
			}
			col := color(r, world)
			ir := int(255.99 * col.x)
			ig := int(255.99 * col.y)
			ib := int(255.99 * col.z)
			fmt.Printf("%v %v %v\n", ir, ig, ib)
		}
	}
}
