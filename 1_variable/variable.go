package main

import "fmt"

func main(){

	/* variable di golang memiliki aturan uniq, 
	yaitu tidak boleh ada satupun variabel di golang yang tidak boleh tidak terpakai, ( wajib di gunakan )

	
	/* 
	konsep variable dengan manifest typing : dengan menyebutkan variable tertentu yang di gunakan ( misal string), 
	maka perlu menuliskan kata string setelah nama variable
	*/ 

	var firstName string = "Agus";
	var lastName string;
	lastName = "Kurniawan";

	fmt.Println("Halo " + firstName + " " + lastName);

	/* 
	konsep variable dengan type inference : yaitu metode deklarasi variable dengan tipe datanya 
	akan di tentukan oleh nilai dari variabelnya dengan syarat langsung mengisi nilai variablenya menggunakan
	tanda :=
	*/ 

	age := 24;
	var city = "Makassar";

	fmt.Println("Age : ",age);
	fmt.Println("City :", city);

	/* 
	golang mendukung konsep multi variable secara bersamaan, caranya menggunakan tanda koma (,) setelah nama variable yang di buat
	*/ 

	var first, second, third string;

	/* 
	pengisian nama variable yang multi variable juga di dukung oleh golang, berikut caranya:
	*/ 

	first, second, third = "satu", "dua", "tiga";
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)

	// lebih ringkasnya untuk multi variable
	name, isMerried, height := "Agus Kurniawan", false, 165;
	fmt.Println("Name : " , name);
	fmt.Println("isMerried : " , isMerried);
	fmt.Println("Height : " , height);

	/* 
	variable underscore di golang.
	seperti yang di katakan di atas, bahwa variable di golang memiliki aturan uniq, 
	yaitu tidak boleh ada satupun variabel di golang yang tidak boleh tidak terpakai, ( wajib di gunakan ), 

	beda dengan variable underscore (_) di golang, variable ini adalah resrved variable yang bisa di gunakan untuk menampung
	nilai yang tidak di pakai. bisa di bilang variable keranjang sampah. konsep variable ini jarang di gunakan, walaupun di gunakan
	pastinya tidak akan pernah terpakai, jadi data apapun yang masuk ke dalam variable underscore ini akan terjebak di dalam variable tersebut
	*/ 
	_ = "Belajar Golang";
	_ = "Semangat Belajar GOLANG";


}