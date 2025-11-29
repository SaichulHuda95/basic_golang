// Pengkondisian di Golang
package main

import "fmt"

func main() {
	nilai := 50 //Nilai awal yang akan dicek
	fmt.Println("Nilai anda :", nilai)

	// -----------------------------------------
	// Cek kondisi nilai
	//-----------------------------------------

	if nilai == 80 {
		//Kondisi terpenuhi jika nilai sama persis dengan 80
		fmt.Println("Nilai Anda A")
	}else if nilai >= 70 {
		//Kondisi terpenuhi jika nilai >= 70 (tapi < 80)
    	fmt.Println("Nilai Anda B")
	}else if nilai >= 60 {
		// Kondisi terpenuhi jika nilai >= 60 (tapi < 70)
    	fmt.Println("Nilai Anda C")
	}else if nilai >= 50 {
		// Kondisi terpenuhi jika nilai >= 50 (tapi < 60)
    	fmt.Println("Nilai Anda D")
	}else {
		// Jika semua kondisi di atas tidak terpenuhi
    	fmt.Println("Nilai Anda E")
	}

	// Switch Case
	switch nilai {
	case 80:
		fmt.Println("Nilai anda 80")
	case 70:
		fmt.Println("Nilai anda 70")
	case 60:
		fmt.Println("Nilai anda 60")
	default:
		fmt.Println("Nilai anda kurang dari 60")
	}
}
