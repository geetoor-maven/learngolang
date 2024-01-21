package main

import "fmt"

func main(){
	/*
		Sebelum ke module ini, silahkan untuk menyelesaikan module yang ada di dalam package array terlebih dahulu.

		di module 6, kita sudah mempelajari tentang perulangan atau for looping di golang.
		for looping erat kaitannya dengan array, dengan menggunakan for looping, kita dapat mengambil semua nilai dalam variable array secara berurutan.

		berikut cara pertama untuk mengambil nilai dalam array menggunakan for looping.
	*/
	var names [5]string;
	names[0] = "Gee";
	names[1] = "Andi";
	names[2] = "Ihsan";
	names[3] = "Alim";
	names[4] = "Fatur";

	// kita gunakan function len() untuk mengambil total berapa element dari array names.
	for i := 0; i < len(names); i++ {
		fmt.Println("Nama Indeks ke, ", i, " : ", names[i]);
	}

	fmt.Println("----------------")

	// cara yang kedua mengakses data dalam array for looping adalah menggunakan looping range. ( cara yang lebih sederhana )
	var cities = [4]string{"Jakarta", "Surabaya", "Makassar", "Bandung"};

	// nilai array di dalam variable cities akan di ambil secara berurutan dan di simpan ke dalam variable citie.
	// dan variable i adalah variable yang akan bertambah 1 (i++) ketika program melakukan perulangan.
	// terlihat lebih singkat dengan cara yang pertama.
	for i, citie := range cities{
		fmt.Println("Kota ke, ", i, " : ", citie);
	}


	fmt.Println("----------------")
	// mengambil inputan user dan menyimpannya kedalam array.
	var userInput [5]string;
	for i:=0; i < len(userInput); i++{
		var name string;
		fmt.Print("Masukan nama ke,", (i+1), ":", )
		fmt.Scanln(&name);

		userInput[i] = name;
	}

	fmt.Println("----------------")
	// munculkan semua nama yanga telah di input
	for i, name := range userInput{
		fmt.Println("Nama ke, ", (i+1), " : ", name);
	}
}