package main

import (
	"fmt"
	"math"
)

func main() {
	var radius int
	fmt.Print("Masukkan jari-jari bola: ")
	fmt.Scan(&radius)

	// Konstanta pi
	const pi float64 = 3.1415926535

	// Hitung volume bola
	volume := (4.0 / 3.0) * pi * math.Pow(float64(radius), 3)

	// Hitung luas kulit bola
	area := 4.0 * pi * math.Pow(float64(radius), 2)

	fmt.Printf("Bola dengan jejari %d memiliki volume: %.2f\n", radius, volume)
	fmt.Printf("Bola dengan jejari %d memiliki Luas kulit bola: %.2f\n", radius, area)
}
