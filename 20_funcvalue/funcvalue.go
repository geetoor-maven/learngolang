package main

import "fmt"

func main(){
	/*
		FUNCTION VALUE.

		di dalam bahasa pemerograman golang, function dapat di simpan di dalam sebuah variable. yang artinya
		sebuah function juga merupakan sebuah tipe data, yang dapat di simpan ke dalam variable.

		berikut contohnya
	*/
	
	// function getLenguangeProgramming akan di simpan ke dalam variable tanpa mengeset nilai parameternya
	getLenguange := getLenguangeProgramming
	// sekarang, variable getLenguange adalah sebuah function yang dapat menerima parameter
	fmt.Println(getLenguange("Golang"))

	totalSum := sum
	fmt.Println(totalSum(1,2,3,4,5))
}

// buat terlebih dahulu function biasa
func getLenguangeProgramming(name string)string{
	return "Programming lenguange, " + name;
}

func sum(numbers ...int)int{
	totalSum := 0
	for _, number := range numbers{
		totalSum += number
	}
	return totalSum;
}