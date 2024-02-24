package main

import "fmt"

func main(){
	/*
		POINTER METHOD
	*/
	
	lenguange := Lenguange{"Java"}
	lenguange.NoPointer()
	fmt.Println(lenguange.Name)


	lenguangeWithPointer := Lenguange{"Golang"}
	lenguangeWithPointer.WithPointer()
	fmt.Println(lenguangeWithPointer.Name)
}

type Lenguange struct{
	Name string
}

func (lenguange Lenguange) NoPointer(){
	lenguange.Name = "Lenguange, " + lenguange.Name
	fmt.Println(lenguange.Name)
}

func (lenguange *Lenguange) WithPointer(){
	lenguange.Name = "Lenguange, " + lenguange.Name
	fmt.Println(lenguange.Name)
}