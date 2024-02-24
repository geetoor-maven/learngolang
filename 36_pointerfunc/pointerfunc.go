package main

import "fmt"

func main(){
	/*
		POINTER DI FUNCTION
	*/
	
	country := Country{"Jakarta", "Indonesia"}
	ChangeCountry(country)

	// data country.Capital tidak berubah menjadi bandung. di karenakan parameter yang di terima oleh function ChangeCountry akan mejadi copyan data tapi tidak bisa merubah data aslinya
	fmt.Println(country)

	// untuk merubah data asli yang di terima oleh parameter ChangeCountry, dapat kita gunakan pointer function.
	var countryPointer *Country = &country 
	ChangeCountryWithPointerParam("Makassar", countryPointer)
	// data capital Jakarta akan berubah sesuai set data parameter pertama ("Makassar")
	fmt.Println(country)
}

type Country struct{
	Capital, Country string
}

// function tidak menggunakan pointer
func ChangeCountry(country Country){
	country.Capital = "Bandung"
}

func ChangeCountryWithPointerParam(capital string, country *Country){
	country.Capital = capital;
}

