package main

import "fmt"

func main() {
	kontak := make(map[string]int)
	kontak["Alice"] = 1234567890
	kontak["Bob"] = 9876543210
	v1 := kontak["Alice"]
	fmt.Println("Semua Kontak", kontak)
	fmt.Println("Nomor telepon Alice :", v1)
	kontak["Charlie"] = 5555555555
	fmt.Println("Setelah menambahkan kontak Charlie:", kontak)
	delete(kontak, "Bob")
	fmt.Println("Setelah menghapus kontak Bob:", kontak)
	fmt.Println("Data kontak:")
	for key, value := range kontak {
		fmt.Printf("%s : %d \n", key, value)
	}

}
