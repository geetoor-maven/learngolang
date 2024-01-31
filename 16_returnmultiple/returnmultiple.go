package main

import "fmt"

func main(){
	/*
		RETURN MULTIPLE VALUE

		di bahasa pemerograman golang, sebuah function tidak hanya mengembalikan 1 value, tapi bisa mengembalikan multiple value ( banyak value ).

		agar function dapat mengembalikan multiple value ( lebih dari satu data ), maka wajib menyebutkan semua tipe dan return valuenya.
	*/
	
	// berikut cara memanggil function yang terdapat return multiple value
	// urutkan sesuai tipe data return value
	// (name)string, (age)int, (isMerried)bool
	name, age, isMerried := getProfile();
	fmt.Println("Name,", name)
	fmt.Println("Age,", age)
	fmt.Println("isMerried,", isMerried)


	// berikut cara menghiraukan atau mengignore sebuah return value di function, dengan menggunakan tanda _
	// jadi return value bitbucket akan di hiraukan oleh program
	github, _, gitlab := getWebsite()
	fmt.Println(github)
	fmt.Println(gitlab)


}

// contoh sebuah function yang mengembalikan multiple value.
// note : ketika sebuah function mengembalikan data yang lebih dari satu, maka wajib menyebutkan tipe data apa saya yang 
// akan di kembalikan di dalam tanda kurung setelah nama function getProfile()
func getProfile()(string, int, bool){
	return "Agus Kurniawan", 25, false;
}

func getWebsite()(string, string, string){
	return "Github", "BitBucket", "Gitlab"
}