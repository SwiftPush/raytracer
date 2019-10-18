package main

import "fmt"

func main() {
	nx, ny := 200, 100
	fmt.Printf("P3\n%v %v \n255\n", nx, ny)
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			r := float64(i) / float64(nx)
			g := float64(j) / float64(ny)
			b := 0.2
			ir := int(255.99 * r)
			ig := int(255.99 * g)
			ib := int(255.99 * b)
			fmt.Printf("%v %v %v\n", ir, ig, ib)
		}
	}
}
