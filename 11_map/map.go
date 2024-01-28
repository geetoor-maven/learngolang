package main

import (
	"fmt"
)

func main(){
	/*
		MAP adalah tipe data yang kumpulan key-value ("key" : "Value") yang di mana kata kunci key harus bersifat uniq (tidak boleh sama)
		sangat berbeda dengan slice atau array, di MAP kita dapat menyimpan data yang sebanyak banyaknya asal key ( kata kunci ) nya harus beda ( karena bersifat uniq )
		yang apabila jika kita membuat kata kunci (key) yang sama, maka data yang sebelumnya dengan key yang sama akan berubah. 

		cara membuat map yang pertama dengan menyebutkan langsung tipe data key dan valuenya : 
	*/
	
	profile := map[string] string{
		// name adalah key
		// Agus adalah value
		"name" : "Agus",
		"address" : "Makassar",
	}

	fmt.Println(profile)

	// cara memunculkan nilai value dari map adalah dengan menyebutkan key nya langsung.
	fmt.Println(profile["name"])
	fmt.Println(profile["address"])

	// menambah data di map atau merubah data di dalam map, cukup memanggil key nya dan isi data baru.
	profile["ipk"] = "3.32";
	fmt.Println(profile["ipk"]);

	// beberapa fungsi yang ada di dalam MAP

	// len() mendapatkan jumlah data di dalam map
	fmt.Println(len(profile))

	// delete(map, key) menghapus data di dalam map
	delete(profile, "ipk");
	fmt.Println(profile)

	// cara membuat map yang kedua :
	lenguangeProgramming := make(map[string]string);
	lenguangeProgramming["1"] = "Golang"
	lenguangeProgramming["2"] = "Java"
	lenguangeProgramming["3"] = "C++"

	fmt.Println(lenguangeProgramming)
}