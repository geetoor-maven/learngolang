package main

import "fmt"

func main(){
	/*
		RECURSIVE FUNCTION
		adalah sebuah function memanggil dirinya sendiri di dalam function
	*/
	
	noRecursive := noRecursive(5)
	fmt.Println(noRecursive)

	sumWithRecursive := sumWithRecursive(5);
	fmt.Println(sumWithRecursive)
	fmt.Println(1+2+3+4+5)
}

// contoh code yang tidak menggunakan recursive function
func noRecursive(value int)int{
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

// contoh code yang menggunakan recursive function, tapi akan terjadi perulangan tanpa henti
func loopRecursive(value string)string{
	if value == "quit" {
		return "Logout"
	}else {
		// loopRecursive() di panggil di dalam function loopRecursive()
		return loopRecursive("welcome..");
	}
}

// contoh kasus yang benar ketika ingin menggunakan recursive function
func sumWithRecursive(value int)int{
	if value == 1 {
		return 1;
	}else {
		return value + sumWithRecursive(value - 1)
	}
}