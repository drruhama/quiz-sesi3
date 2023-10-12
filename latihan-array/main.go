package main

import "fmt"

func main() {
	//buah := [3]string{"Apel", "pisang", "mangga"}
	var buah [3]string
	buah[0] = "Apel"
	buah[1] = "Pisang"
	buah[2] = "Mangga"
	fmt.Println("nama buah pertama:", buah[0])
	fmt.Println("nama buah kedua:", buah[1])
	fmt.Println("nama buah ketiga:", buah[2])
}
