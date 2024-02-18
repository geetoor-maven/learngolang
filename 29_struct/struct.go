package main

import "fmt"

// deklarasi struct yang memiliki 3 buat properti atau atribut.
type Person struct{
	Name, Phone string
	Age int
}

func main(){
	/*
		STRUCT GOLANG

		jika anda pernah belajar java sebelumnya atau bahasa pemerograman yang berorientasi kepada object (oop), maka tidak jarang anda pastinya selalu membuat object.

		di bahasa pemerograman golang, penggunaan OOP tidak di dukung, maka di buat tipe data struktur yang di namakan struct.
		dari sebuah tipe data struct, kita dapat membuat sebuah object ketika struct tersebut di inisialisasi. yang di mana object dari struct tersebut terdapat
		atribut sesuai yang telah di isi di dalam struct.

		lebih jelasnya bisa di lihat pada link di bawah ini : 
		https://www.w3schools.com/go/go_struct.php

	*/
	
	// inisialisasi struct person cara pertama
	var person Person

	person.Name = "Budi"
	person.Phone = "0989865665"
	person.Age = 20

	fmt.Println("Name :", person.Name)
	fmt.Println("Phone :", person.Phone)
	fmt.Println("Age :", person.Age)

	// inisialisasi object struct person dengan cara kedua ( mirip dengan array yaa )
	ikbal := Person{"ikbal", "934304343", 18}
	fmt.Println("Name :", ikbal.Name)
	fmt.Println("Phone :", ikbal.Phone)
	fmt.Println("Age :", ikbal.Age)


	// inisialisasi object struct person degan cara ketiga ( mirip dengan map yaa )
	riswan := Person{
		Name: "Riswan",
		Phone: "98389375935",
		Age: 15,
	}
	fmt.Println("Name :", riswan.Name)
	fmt.Println("Phone :", riswan.Phone)
	fmt.Println("Age :", riswan.Age)


}