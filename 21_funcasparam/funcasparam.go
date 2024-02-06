package main

import "fmt"

func main(){
	/*
		FUNCTION AS PARAMETER

		sebuah function tidak hanya dapat di simpan dalam sebuah variable, tapi sebuah function juga dapat berguna untuk function yang memiliki parameter.
		jadi parameter yang di terima sebuah function adalah function. 
		ini di namakan function as parameter
	*/

	testFilter := filter
	// parameter ketiga menerima function filter() yang akan mengfilter pesan...
	messageFiltering("Budi", "Bro..", testFilter );
	messageFiltering("Budi", "Anjay", testFilter );

	// filter time
	filterTime := timeFilter
	sayTime(7, filterTime)
	sayTime(15, filterTime)
	sayTime(21, filterTime)
	
}

// berikut cara membuat function yang menerima parameter function ( parameter ketiga adalah parameter yang wajib menerima function ketika function di bawah ini di panggil )
func messageFiltering(from string, message string, filterMessage func(string) string){
	fmt.Println("from :", from)
	fmt.Println("message :", filterMessage(message))
}

// fungsi filter ini akan bekerja di dalam function messageFiltering()
func filter(message string)string{
	if message == "Anjay" {
		return "Oppss"
	}else if message == "Anjir" {
		return "Oppss"
	}else{
		return message;
	}
}

// cara simple untuk membuat function as parameter
// yaitu menggunakan type alias. 
type FilterTime func(int) string
func sayTime(time int, filterTime FilterTime){
	println(filterTime(time))
}

func timeFilter(time int)string{
	if time >= 6 && time <= 11 {
		return "Good Morning";
	}else if time > 11 && time <= 18 {
		return "Good Afternoon";
	}else if time > 18 && time <= 24 {
		return "Good Night";
	}else{
		return "No Time";
	}
}

