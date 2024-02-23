package main

import "fmt"

func main(){
	/*
		TYPE ASSERTION
		type assertion merupakan kemampuan untuk merubah sebuah tipe data menjadi tipe data yang di inginkan.
	*/
	
	result := Assert()
	// merubah data interface menjadi ke string
	assertToString := result.(string)
	fmt.Println(assertToString)


	// kode di bawah ini akan menyebabkan panic, di karenakan function Assert memiliki return data string yang tidak bisa di ubah menjadi boolean
	// assertToBool := result.(bool)
	// fmt.Println(assertToBool)

	// untuk lebih aman, bisa menggunakan switch assertion ketika inginn merubah sebuah tipe data yang di inginkan
	resultWithSwitch := AssertWithSwitch()
	switch value := resultWithSwitch.(type){
	case string:
		fmt.Println("String Value", value)
	case int:
		fmt.Println("Int Value", value)
	case bool:
		fmt.Println("Bool Value", value)
	default:
		fmt.Println("Type Unknown")
	}
}

func Assert() interface{}{
	return "ASSERTION";
}

func AssertWithSwitch() interface{}{
	return true
}