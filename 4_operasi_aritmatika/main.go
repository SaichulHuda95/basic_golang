// Operasi Aritmatika Dasar di Golang
package main

import "fmt"

func main()  {
	//----------------------------
	//Inisialisasi variabel angka
	//----------------------------
	angka1 := 10
	angka2 := 3
	fmt.Println("Angka 1 : ", angka1)
	fmt.Println("Angka 2 : ", angka2)

	//operasi tambah
	tambah := angka1 + angka2
	fmt.Println("Hasil penjumlahan : ", tambah)

	//operasi kurang
	kurang := angka1 - angka2
	fmt.Println("Hasil pengurangan : ", kurang)

	//operasi kali
	kali := angka1 * angka2
	fmt.Println("Hasil perkalian : ", kali)

	//operasi bagi
	var bagi float32 = float32(angka1) / float32(angka2)
	fmt.Println("Hasil pembagian : ", bagi)
}
