package main

import (
	"errors"
	"fmt"
)

func main(){
	/*
		ERROR INTERFACE GOLANG
	*/
	
	var theError error = errors.New("The Error")
	fmt.Println(theError.Error())


	result, err := Kali(0, 2)

	// validasi jika return err adalah nil (null), maka sudah pasti tidak ada error
	if err == nil {
		fmt.Println("Hasil Perkalian,", result)
	}else {
		fmt.Println(err)
	}
}


func Kali(value1 int, value2 int) (int , error){
	if value1 == 0 || value2 == 0{
		return 0, errors.New("Tidak boleh masukan angka 0");
	}else {
		return value1 * value2 , nil
	}
}