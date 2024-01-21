package main

import "fmt"

func main(){
	/*
		ARRAY ( kumpulan data dalam satu variable. )
		variable array dapat menyimpan kapasitas data sesuai pengesetan pas pertama kali variable array di buat.

		variable array memiliki nilai indeks yang DIMULAI DARI INDEX 0, sampai dengan nilai indeks yang telah di set di dalam array tersebut.

		apabila variable array nya di berikan tipe data string, maka kumpulan data yang terdapat di array tersebut cuma bisa menampung nilai string. begitupun int, dan bool.
		berikut contoh variable array di bahasa pemerograman golang.
	*/

	fmt.Println("ARRAY");
	// variable array fruit memiliki 3 panjang indeks, jadi variable fruit di bawah ini hanya bisa kita masukan 3 buah data string di dalamnya.
	var fruits [3]string;
	// memasukan data indeks 0
	fruits[0] = "Apple";
	// memasukan data indeks 1
	fruits[1] = "WaterMelon";
	// memasukan data indeks 2
	fruits[2] = "Orange";

	/* 
		note : apabila anda memasukan indeks yang lebih seperti di bawah ini.
		fruit[4] = "Mango";
		maka program akan error karena nilai set array cuma di berikan 3 indeks.
	*/

	// berikut cara menampilkan array fruit dengan memanggil indeks nya
	fmt.Println(fruits[0], fruits[1], fruits[2]);

	/*
		MENGISI VARIABLE ARRAY DENGAN CARA YANG Ke dua.
	*/
	
	var countrys = [4] string{
		// indeks 0
		"Indonesia",
		// indeks 1
		"Philipina",
		// indeks 2
		"Thailand",
		// indeks 3
		"Singapore",
	}
	fmt.Println("Countrys,", countrys);
	// mencari tahu total indeks dari sebuah array dengan menggunakan fungsi len(variableArray)
	fmt.Println("Total Indeks Array:",len(countrys));


	/*
		Array tanpa inisialisasi set indeks.
		cukup menggunakan tanda (...) titik tiga di dalam kurung siku [] maka jumlah array akan menentukan seberapa jumlah data di dalam array.
		
		contohnya seperti di bawah ini.
	*/

	// seceara otomatis indeks array oddNumbers ada 6, karena data di dalam array oddNumbers ada berupa 6 data.
	var oddNumbers = [...]int{3,5,6,7,9,11};
	fmt.Println(oddNumbers);


	/*
		Deklarasi variable array menggunakan function make().
	*/
	var names = make([]string, 2)
	names[0] = "Agus"
	names[1] = "Kurniawan"
	fmt.Println(names)
	
}