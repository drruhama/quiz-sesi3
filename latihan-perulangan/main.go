package main

import "fmt"

func main() {
	var hasil = ""
	for i := 0; i < 5; i++ {
		for y := 5; y > i; y-- {
			hasil += "*"
		}
		hasil += "\n"
	}
	fmt.Println(hasil)
	y := "*"

	for x := "*"; x <= "*****"; x = x + y {
		fmt.Println(x)
	}

}
