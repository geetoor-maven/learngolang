package main

import "fmt"

func main(){
	/*
		PERULANGAN ATAU FOR LOOP.
		selain golang, semua bahasa pemerograman mempunyai metode perulangan atau biasa di katakan for loop.
		perulangan ini adalah prilaku di mana program kita akan berulang terus tanpa henti, selama kodisi yang di deteksi tidak 
		terpenuhi.

	*/
	
	// contoh perulangan yang tidak akan berhenti ( kondisi akan terus menerus false)
	// for i := 0; i == i; i++ {
	// 	fmt.Println(i);
	// }

	// contoh perulangan menggunakan kondisi yang akan berhenti ketika kodisinya sudah false
	value := 5;
	for i := 0; i < value; i++ {
	  fmt.Println("Angka ke : ", i);
	  // kode program akan berhenti ketika variable i sudah lebih besar dari nilai variable value
	}
	fmt.Println("---------------")

	/*
		berikut cara yang berbeda menggunakan keyword looping di golang.
		mirip seperti while di bahasa pemerograman lain.
		golang sendiri tidak memiliki keyword while untuk perulangan, jadi ketika ingin menggunakan konsep while di dalam golang
		makan berikut contoh kode yang sama metode nya dengan while.
	*/
	var i = 0;
	for i < 5{
		fmt.Println("Nomor ke :", i);
		// variable i akan di counter sebelum naik ke pengecekan for kembali.
		i++;
	}

	fmt.Println("---------------")


	/*
		ada juga konsep perulangan tanpa kondisi di dalam golang.
		jadi keyword for nya akan terus true. lalu bagaimana menghentikan programmnya ? dengan keyword break, maka program akan berhenti.
		berikut contoh kodenya.
	*/

	counter := 1;
	for{
		fmt.Println("Angka ke :", counter);
		counter++;

		// disini kita berikan kodisi ketika counter sudah sama dengan 5, maka break program ( hentikan perulangan )
		if counter == 6 {
			break;
		}
	}

	fmt.Println("---------------")

	/*
		Keyword continue dalam golang.
	*/

	valueOdd := 10;
	for i := 0; i < valueOdd; i++ {
		// ketika nilai i dibagi 2 sama dengan 0, maka program akan melanjutkan ke pengecekan selanjutnya yaitu for yang di atas
		if i % 2 == 0 {
			continue;
		}

		// code di sini tidak akan di jalankan ketika kondisi if modulus terpenuhi, karena di dalam kondisi tersebut terdapat keyword
		// continue yang akan melanjutkan kembali ke pengecekan looping selanjutnya.
		fmt.Println("Angka Ganjil,", i);
	}

	fmt.Println("---------------")

	/*
		Perulangan di dalam perulangan ( for dalam for ).
		bukan cuma kondisi yang dapat di tulis kembali di dalam kondisi. tapi perulangan (for) dapat juga melakukan seperti itu.

		tapi perlu di perhatikan, jangan sampai code kita melakukan perulangan terus menerus yang dapat mengakibatkan applikasi dapat crush
	*/
	
	// contoh perulangan bersarang yang menampilkan segitiga pada console
	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}