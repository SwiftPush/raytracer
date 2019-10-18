package main

import "fmt"

func main() {
	nx, ny := 200, 100
	fmt.Printf("P3\n%v %v \n255\n", nx, ny)
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			col := Vector{
				float64(i) / float64(nx),
				float64(j) / float64(ny),
				0.2,
			}
			ir := int(255.99 * col.x)
			ig := int(255.99 * col.y)
			ib := int(255.99 * col.z)
			fmt.Printf("%v %v %v\n", ir, ig, ib)
		}
	}
}
