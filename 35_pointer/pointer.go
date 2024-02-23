package main

import "fmt"

func main(){
	/*
		POINTER GOLANG

		secara default di bahasa pemerograman golang semua variable di passing by value.
		jadi ketika ada 2 buah variabel yang di mana variabel ke 2 akan menduplikasi data dari variabel ke 1, maka ketika variabel ke2 di ubah
		data variabel ke 1 nya tetap ke nilai awal dan tak berubah. jadi seakan akan 2 variable tersebut di simpan di lokasi data yang berbeda.
	*/
	
	// contoh pass by value di golang
	// inisialisasi data struct profile1
	profile1 := Profile{"Agus", "Sinjai", 24}
	// inisialisasi data struct profile2 dengan nilainya dari profile1
	profile2 := profile1

	// profile2.Name di ubah, akan tetapi struct profile1 tidak akan berubah di karenakan profile2 hanya mengcopy data dari profile1
	profile2.Name = "Ikbal"

	fmt.Println(profile1)
	fmt.Println(profile2)


	// jadi gimana cara agar profile2 di ubah maka data profile1 akan ikut keubah ? variable akan di passing by reference ( data akan di simpan di satu lokasi yang sama )
	// berikut cara pass by reference di golang
	profileRef1 := Profile{
		Name: "Geetoor",
		Address: "Sinjai",
		Age: 22,
	}

	// menggunakan keyword & maka akan terjadi passing by reference
	profileRef2 := &profileRef1
	profileRef3 := &profileRef1

	// profileRef1 akan ikut berubah
	profileRef2.Name = "Riswan"

	fmt.Println(profileRef1)
	fmt.Println(profileRef2)


	*profileRef2 = Profile{"Fathur", "Sinjai", 18}
	fmt.Println(profileRef1)
	fmt.Println(profileRef2)
	fmt.Println(profileRef3)


	var profile *Profile = new(Profile)
	profile.Name = "Profile"
	fmt.Println(profile)
}

type Profile struct{
	Name, Address string
	Age int
}