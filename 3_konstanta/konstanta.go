package main

import "fmt"

func main(){
	/*
		konstanta merupakan jenis variable yang nilainya tidak akan bisa di ubah. 
		pengisian nilai dari variable konstanta harus di lakukan di awal, dan setelah itu nilainya tidak boleh di ubah kembali.
	*/

	// nilai dari variable ipAddress tidak boleh di ubah kembali. karena merupakan variabel yang jenis konstanta
	const ipAddress string = "192.1.1.1";
	fmt.Println("connect to ip :", ipAddress)

	
	/*
		multi konstata, konstanta juga dapat di deklarasikan secara bersamaan.
	*/
	const(
		// variable di deklarasikan dengan secara metode type inference
		name = "golang"
		years = 2009
		authors = "google"
		isRecomend = true
	)
	fmt.Println("Bahasa pemerograman :", name);
	fmt.Println("Tahun di buat :", years);
	fmt.Println("Pencipta :", authors);
	fmt.Println("is Recomend :", isRecomend);

	// contoh deklarasi konstanta yang lain.
	const(
		today = "minggu"
		/*
			isi nilai variable hariIni adalah minggu, karena akan mengambil nilai today sama seperti di atasnya
		*/
		hariIni
	)
	fmt.Println("Today :", today);
	fmt.Println("Hari ini :", hariIni);

	/*
		contoh inisialisasi nilai konstanta dengan satu baris
	*/
	const one, two, three = 1, 2, 3;
	fmt.Println("Satu :", one);
	fmt.Println("Dua :", two);
	fmt.Println("Tiga :", three);

}