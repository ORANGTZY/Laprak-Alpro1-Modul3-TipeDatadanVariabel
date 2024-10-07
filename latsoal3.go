package main

import "fmt"

func isKabisat(tahun int) bool {
	if tahun%400 == 0 {
		return true
	} else if tahun%100 == 0 {
		return false
	} else if tahun%4 == 0 {
		return true
	}
	return false
}

func main() {
	var tahun int
	fmt.Print("Masukkan tahun: ")
	fmt.Scan(&tahun)

	if isKabisat(tahun) {
		fmt.Printf("%d adalah tahun kabisat: true\n", tahun)
	} else {
		fmt.Printf("%d bukan tahun kabisat: false\n", tahun)
	}
}
