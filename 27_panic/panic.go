package main

import "fmt"

func main(){
	/*
		PANIC function.
		
		panic function adalah sebuah function yang dapat di gunakan untuk menghentikan program ketika terjadi error.

		ketika ada panic function, program akan berhenti, akan tetapi defer function akan tetap di jalankan
	*/

	run();
	divideNumber(1, false)

	// function divideNumber(0) akan error karena pembagian tidak boleh terjadi kalau angka 0, tapi deferRun() akan tetap di jalankan dan panic() akan di jalankan
	divideNumber(0, true)
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

// berikut contoh ke 2 code defer dan panic funcion yang terjadi error tapi defer function tetap di jalankan.
func divideNumber(value int, err bool){
	// function defer deferRun akan di jalankan ketika function divideNumber() selesai di jalankan walaupun terjadi error di function divideNumber
	defer deferRun()
	if err {
		panic("Stop Program.... ada error")
	}

	fmt.Println("Pembagian angka..")
	result := 5 / value;
	fmt.Println("Hasil Bagi :", result)
}
