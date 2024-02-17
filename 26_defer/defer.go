package main

import "fmt"

func main(){
	/*
		DEFER function.
		function yang dapat di jadwalkan setelah sebuah function lain di eksekusi.

		defer function akan selalu di eksekusi meskipun terjadi error di function yang di eksekusi sebelum defer function.
		
	*/

	run();
	divideNumber(1)

	// function divideNumber(0) akan error karena pembagian tidak boleh terjadi kalau angka 0, tapi deferRun() akan tetap di jalankan
	divideNumber(0)
}

// berikut contoh code defer function
func run(){
	// function defer deferRun akan di jalankan ketika function run() selesai di jalankan
	defer deferRun()

	fmt.Println("Run Application...")
	fmt.Println("On Progress.....")
}

func deferRun(){
	fmt.Println("Done Run....")
	fmt.Println("............")
}

// berikut contoh ke 2 code defer funcion yang terjadi error tapi defer function tetap di jalankan.
func divideNumber(value int){
	// function defer deferRun akan di jalankan ketika function divideNumber() selesai di jalankan walaupun terjadi error di function divideNumber
	defer deferRun()

	fmt.Println("Pembagian angka..")
	result := 5 / value;
	fmt.Println("Hasil Bagi :", result)
}
