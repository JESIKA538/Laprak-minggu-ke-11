package main

import "fmt"

func main() {
	var tipe_kendaraan  string
	var durasi, tarif int
	fmt.Print("Masukkan jenis tipe kendaraan (Motor/Mobil/Truk): ")
	fmt.Scan(&tipe_kendaraan)
	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scan(&durasi)

	if durasi < 1 {
		durasi = 1
	}
	switch tipe_kendaraan {
	case "motor":
		tarif = durasi * 2000
	case "mobil":
		tarif = durasi * 5000
	case "truk":
		tarif = durasi * 8000
	default:
		fmt.Println("Jenis kendaraan atau durasi parkir tidak valid")
	}
	fmt.Printf("Tarif Parkir: Rp %d\n", tarif)
}
