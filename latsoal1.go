package main

import (
	"fmt"
	"log"
)

func calculateX(fx float64) (float64, error) {
	// Cek apakah fx + 5 = 0, yang akan membuat pembagian menjadi tidak valid
	if fx == -5 {
		return 0, fmt.Errorf("nilai f(x) tidak boleh sama dengan -5 karena menghasilkan pembagian dengan nol")
	}
	// Hitung nilai x dari persamaan f(x)
	x := (2 / (fx + 5)) + 5
	return x, nil
}

func main() {
	var fx float64
	fmt.Print("Masukkan nilai f(x): ")
	_, err := fmt.Scan(&fx)
	if err != nil {
		log.Fatal("Input tidak valid")
	}

	// Hitung nilai x berdasarkan f(x)
	x, err := calculateX(fx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Nilai x adalah: %.4f\n", x)
}
