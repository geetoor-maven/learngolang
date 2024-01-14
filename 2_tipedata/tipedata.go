package main

import "fmt"

func main(){
	/*
		Tipe data numerik non-decimal atau non floating point di golang ada beberapa jenis.
		ada 2 tipe data kategori non-decimal yang perlu di ketahui yaitu:
		uint : tipe data untuk bilangan cacah. (bilangan positif)
		int : tipe data untuk bilangan bulan ( positif dan negatif)

		banyak tipe data yang berjenis uint dan int yang dapat menyimpan maksimal kapasitas data, boleh di lihat pada link di bawah ini : 
		https://www.w3schools.com/go/go_integer_data_type.php
	*/
	var angkaPositif uint8 = 120;
	var angkaNegatif int = -134343453;
	fmt.Println("Angka Positif :", angkaPositif);
	fmt.Println("Angka Negatif :", angkaNegatif);

	
	/*
		tipe data numerik decimal digolang ada 2 jenis, float32 dan float 64.
		jenisnya cuma berada di lebar cakupan nilai yang bisa di tampung.
		lebih lanjut penjelasannya ada di sini : 
		https://www.w3schools.com/go/go_float_data_type.php
	*/
	var angkaFloat32 float32 = 2.00;
	var angkaFloat64 float64 = 23.23432;
	fmt.Println("Angka Float 32 :", angkaFloat32);
	fmt.Println("Angka Float 64 :", angkaFloat64);



	/*
		tipe data boolean, sama halnya di bahasa pemerograman lain. golang juga dapat menyimpan variable dengan nilai benar (true)
		dan salah (false). BOOLEAN. tipe data ini akan sangat di pakai pada bagian kondisi dan perulangan
	*/
	var empty bool = true;
	fmt.Println("Kosong :", empty);


	/*
		tipe data String, sama dengan di bahasa pemerograman lainnya, golang juga mempunyain tipe data yang dapat menampung karakter ( string ).
		ciri khas dari tipe data ini ialah, nilainya di apit oleh 2 tanda petik( " " )
	*/
	var message string = "ini data string";
	fmt.Println(message);



}