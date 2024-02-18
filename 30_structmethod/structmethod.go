package main

import "fmt"

// deklarasi struct yang memiliki 3 buat properti atau atribut.
type Person struct{
	Name, Phone string
	Age int
}

// ini adalah struct function atau method struct, yang di mana ketika object person di buat, maka function getPersonName() akan di miliki oleh object person
func (person Person) getPersonName(){
	fmt.Println("My Name is :", person.Name)
}

func (person Person) getPersonPhone(){
	fmt.Println("Phone :", person.Phone)
}

func main(){
	/*
		STRUCT METHDO
		di karenakan struct adalah sebuah tipe data, maka type struct juga dapat di sisipkan ke dalam sebuah parameter function.
	*/
	person := Person{
		Name: "Agus",
		Phone: "29304343",
		Age: 25,
	}

	// function getPersonName() dapat di akses melalui object person struct
	person.getPersonName()
	person.getPersonPhone()
}