package main

import "fmt"

func main(){
	/*
		OPERATOR ARITMATIKA.
		operator aritmatika di golang merupakan operator yang bersifat perhitungan matematika.
		golang mendukung beberapa operasi aritmatika di antaranya : 
		* : perkalian
		- : pengurangan
		+ : penjumlahan
		/ : pembagian
		% : modulus / sisa bagi
		hampir semua bahasa pemerograman mendukung operasi aritmatika ini.
	*/
	fmt.Println("OPERATOR ARITMATIKA");

	// masih ingat tipe deklarasi seperti ini ? 
	const(
		angka1 = 3
		angka2
	)

	var penjumlahan int = angka1 + angka2;
	var pengurangan  = angka1 - angka2;
	var perkalian  = angka1 * angka2;
	var pembagian  = angka1 / angka2;
	var modulus  = angka1 % angka2;
	fmt.Println("Hasi Penjumlahan 3 + 3 :", penjumlahan);
	fmt.Println("Hasi Pengurangan 3 - 3 :", pengurangan);
	fmt.Println("Hasi Perkalian 3 * 3 :", perkalian);
	fmt.Println("Hasi Pembagian 4 / 4 :", pembagian);
	fmt.Println("Hasi Sisa Bagi 4 % 4 :", modulus);


	/*
		OPERATOR PERBANDINGAN.
		operator perbandingan di gunakan untuk menentukan nilai benar atau salah  ( true or false )
		beberapa operator perbandingan di golang :
		== : menentukan apakah nilai kiri dan kanan itu sama 
		!= : menentukan apakah nilai kiri dan kanan itu tidak sama
		> : menentukan apakah nilai kiri lebih besar dari nilai kanan
		< : menentukan apakah nilai kiri lebih kecil dari nilai kanan 
		>= : menentukan apakah nilai kiri lebih besar sama dengan dari nilai kanan
		<= : menentukan apakah nilai kiri lebih kecil sama dengan dari nilai kanan
	*/
	fmt.Println("OPERATOR PERBANDINGAN");
	var valueOne int = 5;
	var valueTwo int = 10;

	// coba ubah aja operator logika == menjadi > atau < , dan liat hasilnya true atau false
	var isSame = valueOne == valueTwo;

	fmt.Println("Apakah 3 sama dengan 3 ? ", isSame)


	/*
		OPERATOR LOGIKA.
		operator logika di gunakan untuk menentukan hasil benar tidaknya dari operator perbandingan.
		beberapa operator logika yang ada di golang : 
		&& : and : operator dan dan atau kadang di sebut AND, akan bernilai true apabila dua operator perbandingan nilanya true
		|| : atau : operator atau, kadang di sebut OR AKAN bernilai tru jika salah satu operator perbandingannya nilainya true
		! : negasi atau membalikkan nilai true ke false, begitupun sebaliknya.
	*/
	fmt.Println("OPERATOR LOGIKA");

	var isTrue = true;
	var isFalse = false;

	// hasilnya akan false ( di karenakan nilai keduanya tidak true )
	fmt.Println(isTrue && isFalse);
	// hasilnya akan true ( di karenakan nilai keduanya salah satunya true )
	fmt.Println(isTrue || isFalse);
	// hasilnya akan false ( karena membalikkan nilai true jadi false)
	fmt.Println(!isTrue);


}