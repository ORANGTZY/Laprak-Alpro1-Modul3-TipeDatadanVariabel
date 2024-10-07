package main

import (
	"fmt"
	"log"
)

// Function to convert Celsius to other temperature scales
func convertTemperature(celsius float64) (fahrenheit, reamur, kelvin float64) {
	// Convert Celsius to Fahrenheit
	fahrenheit = (celsius * 9 / 5) + 32

	// Convert Celsius to Reamur
	reamur = celsius * 4 / 5

	// Convert Fahrenheit to Kelvin
	kelvin = (fahrenheit + 459.67) * 5 / 9

	return fahrenheit, reamur, kelvin
}

func main() {
	var celsius float64

	// Input temperature in Celsius
	fmt.Print("Masukkan suhu dalam Celsius: ")
	_, err := fmt.Scan(&celsius)
	if err != nil {
		log.Fatal("Input tidak valid. Masukkan angka yang valid.")
	}

	// Convert temperature
	fahrenheit, reamur, kelvin := convertTemperature(celsius)

	// Display results
	fmt.Println("\nHasil Konversi Suhu:")
	fmt.Printf("Derajat reamur dari %.2f°C = %.2f°R\n", celsius, reamur)
	fmt.Printf("Derajat fahrenheit dari %.2f°C = %.2f°F\n", celsius, fahrenheit)
	fmt.Printf("Derajat kelvin dari %.2f°C = %.2f°K\n", celsius, kelvin)
}
