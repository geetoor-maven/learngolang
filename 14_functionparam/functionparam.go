package main

import "fmt"

func main(){
	/*
		FUNCTION PARAMETER.

		PARAMETER adalah data yang dapat di ambil oleh function dari luar . dan di dalam function, kita dapat menambahkan parameter sebanyak yang kita perlukan.
		jadi ketika sebuah function terdapat parameter, maka ketika function tersebut di inisialisasi, wajib memasukan data ke dalam parameternya.
	*/

	// pemanggilan function tanpa parameter
	sayHelloWorld()

	// pemanggilan function dengan parameter
	sayHello("Agus", "Kurniawan")

	firstName := "Gee"
	lastName := "Maven"
	sayHello(firstName, lastName)
}

// contoh function yang tidak memiliki parameter
func sayHelloWorld(){
	fmt.Println("Hello World...");
}

// contoh function yang memiliki parameter
func sayHello(firstName string, lastName string){
	fmt.Println("Hello,", firstName, lastName);
}