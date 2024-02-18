package main

import "fmt"

// deklarasi interface di golang
type Profile interface{
	// nilai default interface di bawah ini adalah nil jika belum ada kontrak struct nya.
	GetName() string
	GetAddress() string
	IsMerried() bool
}

func getProfile(profile Profile){
	fmt.Println("name :", profile.GetName())
	fmt.Println("address :", profile.GetAddress())
	fmt.Println("merried :", profile.IsMerried())
}


type Person struct{
	Name, Address string
	Merried bool
}

// implementasi interface GetName 
func (person Person) GetName()string{
	return person.Name
}

// implementasi interface GetAddress 
func (person Person) GetAddress()string{
	return person.Address
}

func (person Person) IsMerried()bool{
	return person.Merried
}

func main(){
	/*
		INTERFACE GOLANG

		interface adalah kumpulan defenisi function/method yang tidak memiliki body ( isi function ).
		di bahasa pemerograman golang, interface adalah tipe data yang di mana dafault value nya adalah nil, dan hanya dapat di gunakan
		jika interface tersebut sudah terisi.
	*/

	var person Person
	person.Name = "Agus"
	person.Address = "Makassar"
	person.Merried = false

	// di function getProfile terdapat paramter struct Person yang mengimplementasi interface Profile
	getProfile(person)


	// memanggil interface kosong
	fmt.Println(returnAny())
	

	// parameter kedua sebagai interface dapat di berikan nilai tipe data apapun
	paramAny(1, true)
	paramAny(1, "true")
	paramAny(1, 20)
}


// ada juga yang di namakan interface kosong di bahasa pemerograman golang.
// ketika kita menggunakan function yang terdapat interface kosong, maka return type nya akan bebas di berikan tipe datanya
// sekilas mirip tipe data Object kalau di bahasa pemerograman java.
// berikut contohnya
func returnAny() interface{}{
	// return tipe nya bebas ( bisa string, int, bool )
	return "Golang Interface"
}

// contoh kedua untuk interface kosong
func paramAny(value int, inter interface{}){
	fmt.Println("Value :", value)
	fmt.Println("Interface :", inter)
}