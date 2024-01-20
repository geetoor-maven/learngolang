package main

import "fmt"

func main(){
	/*
		sama seperti di bahasa pemerograman lain, seperti bahasa pemerograman java, python, c++.
		golang juga terdapat function yang dapat di gunakan untuk memberikan perintah kepada user agar menginputkan angka atau huruf di console.
		Untuk mengambil inputan user di bahasa pemerograman golang, dapat menggunakan function fmt.Scanln(&namaVariable).

		jika anda menggunakan visual studio code untuk memerogram golang, gunakan terminal atau cmd untuk menjalankan program yang terdapat inputan di dalam nya.
		
		Berikut adalah contoh code mengambil inputan user di golang.
	*/

	var resultNumber int;

	// gunakan function print() agar program tidak membuat line baru, agar user dapat mengerti kalau user sedang di minta inputan oleh program.
	fmt.Print("Masukan Angka Random :");
	// gunakan function scanln() dan masukan karakter & kemudian nama variable yang ingin di isi value dari hasil inputan user
	fmt.Scanln(&resultNumber);
	fmt.Println("Angka yang di input :", resultNumber)

	var resultString string;
	fmt.Print("Masukan Karakter Random :");
	fmt.Scanln(&resultString);
	fmt.Println("Karakter yang di input :", resultString)

}