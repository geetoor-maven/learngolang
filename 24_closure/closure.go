package main

import "fmt"

func main(){
	/*
		CLOSURE GOLANG
		closure adalah sebuah metode di mana variable dapat di akses di dalam function selama variable tersebut berada di scope yang sama.
		contoh nya : 
	*/
	
	// variable name di inisialisasi dengan data "variable name"
	name := "variabel name"

	changeName := func ()  {
		// kemudian variable name di ganti di dalam function changeName ( ini bisa karena masih dalam satu ruang lingkup scope )
		name = "di ganti"
		fmt.Println(name)
	}

	// jadi name nya sekarang berubah menjadi kata "di ganti"
	changeName()
	changeName()

	fmt.Println(name)

}