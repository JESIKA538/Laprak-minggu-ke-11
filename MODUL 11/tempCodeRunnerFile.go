package main

import "fmt"

func main() {
	var bilangan int
	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&bilangan)

	switch {
	case bilangan%2 != 0:
		fmt.Printf("Kategori: Bilangan Ganjil\n")
		hasil := bilangan + (bilangan + 1)
		fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n", bilangan, bilangan+1, hasil)

	case bilangan%2 == 0 && bilangan%10 != 0:
		fmt.Printf("Kategori: Bilangan Genap\n")
		hasil := bilangan * (bilangan + 1)
		fmt.Printf("Hasil perkalian dengan bilangan berikutnya %d * %d = %d\n", bilangan, bilangan+1, hasil)

	case bilangan%10 == 0:
		fmt.Printf("Kategori: Bilangan Kelipatan 10\n")
		hasil := bilangan / 10
		fmt.Printf("Hasil pembagian antara %d / 10 = %d\n", bilangan, hasil)

	case bilangan%5 == 0:
		fmt.Printf("Kategori: Bilangan Kelipatan 5\n")
		hasil := bilangan * bilangan
		fmt.Printf("Hasil kuadrat dari %d^2 = %d\n", bilangan, hasil)
	
}
}