package main

import "fmt"

func main(){
	/*
		ANONYMOUS FUNCTION

		anonymous function merupakan metode mendeklarasikan function secara langsung ke variable tanpa harus buat functionnya di luar
		blok code main(), biasanya di sebut dengan function tanpa nama (anonymous function)
	*/
	
	// contoh code anonymous function yang langsung di assign ke dalam sebuah variable
	name := func (name string) string {
		return name;
	}
	fmt.Println(name("Budi"))

	user := func(userName string) string  {
		return userName
	}
	selectUser("Admin", user)
	selectUser("Bro", user)
	selectUser("test", func(s string) string {
		return "Root"
	});

}

type DetectionUser func(string) string

func selectUser(userName string, detectUser DetectionUser){
	if detectUser(userName) == "Admin" {
		fmt.Println("Welcome Admin")
	}else if detectUser(userName) == "Root" {
		fmt.Println("Welcome Root")
	}else {
		fmt.Println("Welcome....", userName)
	}
}