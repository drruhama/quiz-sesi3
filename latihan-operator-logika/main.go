package main

import "fmt"

func main() {
	gender := "female"
	age := 21
	if gender == "female" || age >= 21 {
		fmt.Println("Pelamar dapat dipekerjakan")
	} else {
		fmt.Println("Pelamar tidak dapat dipekerjakan")
	}
}
