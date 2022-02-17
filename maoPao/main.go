package main

import "fmt"

func main() {
	var maop = []int{1, -2, 22, 10, -5}
	for i := 0; i < len(maop); i++ {
		for j := 0; j < len(maop)-1-i; j++ {
			if maop[j] > maop[j+1] {
				maop[j], maop[j+1] = maop[j+1], maop[j]
			}
			fmt.Println("第", j, "次", maop)
		}
	}
	fmt.Println(maop)
}
