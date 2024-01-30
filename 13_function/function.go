package main

import "fmt"

func main(){
	/*
		FUNCTION 
		function merupakan sebuah blok kode yang di buat agar dapat di gunakan secara berulang ulang.

		di karenakan golang tidak mendukung object oriented programming (oop) seperti java, maka nanti ketika membuat sebuah code yang rapi,
		dapat di pastikan kita akan banyak sekali membuat function agar kodingan terlihat sedikit rapi.

		di golang, untuk membuat function hanya membutuhkan kata kunci func kemudian di ikuti nama function nya apa dan isi blok kode dari function tersebut.

		kemudian ketika ingin menjalankan function tersebut, kita hanya perlu memanggil nama function nya saja.

	*/
	
	// cara memanggil function, cukup memanggil nama functionnya saja dan kemudian blok kode yang ada di dalam function tersebut akan di eksekusi oleh program
	iniNamaFunction();

	sayGood()
	sayGood()
	sayGood()
}

// contoh membuat function
// note : best praktis untuk menamai sebuah function adalah di awali dengan huruf kecil, dan jika ada kata baru, maka awali dengan huruf besar
func iniNamaFunction(){
	fmt.Println("ini adalah nama function")
}

func sayGood(){
	fmt.Println("Say Good...");
}