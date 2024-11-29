package main

import "fmt"

func main() {
	var kadar_Ph float64
	fmt.Print("Masukkan nilai pH air: ")
	fmt.Scan(&kadar_Ph)

	switch {
	case kadar_Ph >= 6.5 && kadar_Ph <= 8.6:
		fmt.Println("Air layak minum.")
	case kadar_Ph < 0 || kadar_Ph > 14:
		fmt.Println("Nilai pH tidak valid. Harus antara 0 dan 14.")
	default:
		fmt.Println("Air tidak layak minum.")
	}
}
