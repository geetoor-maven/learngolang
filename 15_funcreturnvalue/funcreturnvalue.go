package main

import "fmt"

func main(){
	/*
		FUNCTION RETURN VALUE.

		note : FUNCTION dapat mengembalikan data.
		ini di sebut dengan return value pada sebuah function.

		agar function tersebut dapat mengembalikan data, maka ketika membuat function, harus menginisialisasi tipe data apa yang harus di kembalikan dari functin tersebut.

		jika function tersebut di inisialisasi dengan tipe data string, maka tipe kembalian data dari function tersebut adalah string.
		untuk mengembalikan data dari sebuah function, menggunakan kata kunci return.
	*/
	
	// memanggil function return value dan menyimpan value nya didalam sebuah variabel sentence.
	sentence := getSentence()
	// jadi isi data sentence adalah : this sentence. 
	// karena nilai kembalian dari function getSentence() adalah this sentence
	fmt.Println(sentence)

	goodName := sayGood("agus")
	fmt.Println(goodName)

	good()

	isGender := getGender(true)
	fmt.Println(isGender)
	
}

// contoh function yang mengembalikan data di bahasa pemerograman golang.
// setelah nama function() kemudian ketik tipe data apa yang harus di kembalikan function tesebut (misal string, maka return nya harus string)
func getSentence()string{
	sentence := "this sentence"
	return sentence;
}

// contoh function yang mengembalikan data dengan parameter
func sayGood(name string) string{
	return "Hello, " + name + " are you good ?"
}

// function biasa
func good(){
	fmt.Println("Yes, iam good");
}

// contoh function yang mengembalikan data dan terdapat kondisi di dalamnya
func getGender(isGender bool)string{
	if isGender {
		return "Your is man"
	}else {
		return "Your is woman"
	}
}