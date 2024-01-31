package main

import "fmt"

func main(){
	/*
		NAMED RETURN VALUE

		di bahasa pemerograman golang, kita dapat mengubah tipe data multiple value menjadi nama variable.
		mungkin agak membingungkan, di karenakan sepengetahuan penulis, cuma bahasa pemerograman golang yang mempunyai feature seperti ini.

		lebih jelasnya mungkin dapat di lihat di link di bawah ini :
		https://www.w3schools.com/go/go_function_returns.php


	*/
	
	facebook, instagram, twitter := getWebsite()
	fmt.Println(facebook)
	fmt.Println(instagram)
	fmt.Println(twitter)


	name, _ := getName()
	fmt.Println(name)
}

// contoh named return value di golang
// menggunakan nama variable, kemudian tipe data nya, dan variablenya dapat di assign nilai di dalam blok kode function
func getWebsite()(facebook, instagram, twitter string){
	facebook = "facebook.com"
	instagram = "instagram.com"
	twitter = "twitter.com"

	// cukup menggunakan return tanpa menyebut nama value, golang akan mengembalikan 3 nilai yang di atas.
	// sama saja seperti ini : return facebook, instagram, twitter
	return
}

// sebelumnya kalau return multiple value tanpa named return value
// hanya menggunakan tipe data (string, string)
func getName()(string, string){
	return "Ikbal", "Riswan";
}