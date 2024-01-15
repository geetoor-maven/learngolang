package main

import (
	"fmt"
)

func main(){
	/*
		if kondisi ini gunanya untuk mengatur alur program. analoginya seperti saat kita sedang lapar, maka kita harus makan.
		atau seperti rambu lalu lintas, jika rambunya belok ke kiri maka kita harus belok kekiri juga.

		yang menjadi acuan dalam proses seleksi kondisi adalah nilai yang bertipe boolean yang berasal dari variable ataupun 
		hasil dari operator perbandingan. nah, nilai tersebutlah yang akan menentukan jalannya program yang kita buat.

		dalam bahasa golang, ada 2 cara membuat kondisi. 
		- if else kondisi
		- switch case

		if kondisi di golang kurang lebih sama dengan bahasa pemerograman yang lain.
		yang membedakan di golang tidak perlu menulis tanda kurung kondisi setelah if.
	*/

	fmt.Println("IF ELSE KONDISI");

	nilaiUjian := 77;

	// coba sebutkan apa yang akan muncul ketika program di jalankan ? 
	if nilaiUjian == 100 {
		fmt.Println("Anda lulus dengan sempurna");
	}else if nilaiUjian >= 80 && nilaiUjian <= 90 {
		fmt.Println("Anda lulus dengan sangat baik");
	}else if nilaiUjian >= 70 && nilaiUjian < 80 {
		fmt.Println("Anda lulus lumayan baik");
	}else {
		fmt.Println("Anda tidak lulus");
	}

	/*
		output yang akan muncul dari kondisi di atas adalah "Anda lulus lumayan baik", karena nilai ujian lebir besar sama dengan 70 ( true ),
		dan nilai ujian lebih kecil dari 80 ( true ) dan di apit oleh operator perbandingan AND yang di mana jika dua duanya true 
		maka hasil nya pun akan true.
	*/


	/*
		VARIABLE PENAMPUNG DALAM IF ELSE.
		variable penampung dalam if else pada bahasa pemerograman golang hanya bisa di tempatkan di dalam blok kondisi pada if else.
		langsung saja ke prakteknya.
	*/

	skor := 9000;
	// variable persen di sini adalah variable penampung didalam kondisi if else, gunanya agar membuat kodingan lebih terlihat rapi dan singkat ( tidak perlu mendeklarasikan variabel persen seteleh skor)
	if persen:= skor / 100; persen >= 100 {
		fmt.Println("Skor A");
	}else if persen >= 75 {
		fmt.Println("B")
	}else {
		fmt.Println("C")
	}

	/*
		SWITCH CASE.
		sama seperti bahasa pemerograman lainnya, golang juga mempunyai switch case dalam menentukan kondisi.
		switch case merupakan kondisi yang memfokuskan kepada satu variable. kemudian variable tersebut akan di cek nilainya satu persatu.
		contoh : dalam sebuah variable skor adalah A, B, C atau seterusnya.
	*/

	var huruf = "A";
	switch huruf{
	case "A":
		fmt.Println("Huruf A");
	case "B":
		fmt.Println("Huruf B");
	case "C":
		fmt.Println("Huruf C");
	}

	/*
		di atas merupakan contoh kodingan menggunakan switch case, di mana ketika variable huruf isinya "A" maka program akan berhenti 
		di case "A". di bahasa pemerograman golang, kita tidak perlu menggunakan keyword break setelah case yang di mana di miliki oleh bahasa 
		pemerograman lain, seperti java. golang akan menghentikan otomatis kondisi ketika nilai sudah terpenuhi di dalam case.
	*/


	/*
		sebuah case juga dapat menampung banyak kondisi. contoh kodingan nya seperti yang di bawah ini.
	*/

	angka := 5;
	switch angka{
	case 1,2,3:
		fmt.Println("Angka 1 2 3");
	case 4,5,6:
		fmt.Println("Angka 4 5 6");
	case 7,8,9:
		fmt.Println("Angka 7 8 9");
	}
	// secara otomatis case akan masuk ke dalam case yang kedua yang terdapat angka 5 di dalamnya dikarenakan isi dari variable angka adalah 5

	/*
		SWITCH IF ELSE
		di bahasa pemerograman golang, switch case dapat di gunakan dengan gaya if - else, keyword if else akan di tulis setelah keyword case.
		bisa liat contoh di bawah ini.
	*/

	heigth := 171;
	// setelah keyword switch, langsung buka kurung kurawal tanpa memanggil variable heigth.
	switch {
	case heigth >= 170:
		fmt.Println("Anda Sangat Tinggi.");
	case (heigth >= 165 && heigth < 170):
		fmt.Println("Lumayan Tinggi.")
	default:
		{
			// kode di dalam  default akan di jalankan ketika tidak ada satupun syarat yang memenuhi case yang di atas.
			fmt.Println("Semangat Berolahraga Guys.")
			fmt.Println("Sehat Selalu")
		}
	}


	/*
		FALLTHROUGH SWITCH
		agak sedikit berbeda dengan bahasa pemerograman lain, di golang ada yang di namakan fallthrough di dalam switch case.
		keyword fallthrough ini di gunakan untuk memaksa program di teruskan ke case selanjutnya tanpa menghiraukan kondisi case nya.
		artinya akan selalu di anggap benar 
	*/
	weight := 50;
	switch{
	case weight >= 60:
		fmt.Println("Berat anda 60 ke atas");
	case (weight > 40) && (weight < 60):
		fmt.Println("Anda Lumayan Berat");
		fallthrough
	case weight <= 40:
		// kode di dalam case ini akan tetap di jalankan, di karenakan ada keyword fallthrough di dalam case ke 2
		fmt.Println("Jaga Kesehatan...")
	}

}