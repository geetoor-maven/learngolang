package main

import "fmt"

func main(){
	/*
		package pembelajaran di pisah2 
		di karenakan golang sendiri akan mendeteksi cukup 1 fungsi main() di dalam satu package, 
		lebih dari satu akan terjadi error
	*/
	fmt.Println("Hello World");
}