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

	
	

}