// Casting tipe data adalah proses mengubah tipe data dari suatu variabel ke tipe data lain
package main

import "fmt"

func main()  {
	//Mengubah int ke float64
	var_int := 10 //Contoh variable bertipe integer
	var_float := float64(var_int)
	fmt.Println(var_float)

	//Mengubah float64 ke int
	var_float2 := 4.6 //Contoh variable bertipe float64
	var_int2 := int(var_float2)
	fmt.Println(var_int2)
}
