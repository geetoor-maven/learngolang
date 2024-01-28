package main

import "fmt"

func main(){
	/*
		TIPE DATA SLICE.
		tipe data slice merupakan tipe data yang sangat berbeda, dan mungkin baru pertama kali anda temui di golang. karena sangat
		jarang tipe data ini di miliki oleh bahasa pemerograman lain.

		tipe data slice merupakan tipe data potongan dari array. ( mirip dengan tipe data array ) tapi yang membedakan dari tipe data ini
		adalah ukuran datanya yang bisa berubah.

		seperti yang saya katakan bahwa slice ini hampir mirip dengan tipe data array, yang dapat menampung banyak data di dalamnya. 
		cara kerja di belakang slice ini adalah menggunakan tipe data array, yang di mana tipe data slice ini akan menyimpan datanya di dalam array ( referensi datanya dari array).


		berbeda dengan array yang cuma memiliki fungsi length yang dapat mengambil seluruh total element di dalamnya, di slice ada 3 hal yang perlu di ketahui : 
		- length : adalah fungsi untuk mengambil nilai atau panjang data dari slice ( sama juga dengan length di array ).
		- pointer : data pertama di array dari slice
		- capacity : kapasitas dari tipe data slice
	*/
	

	// contoh membuat tipe data slice
	// buat tipe data array terlebih dahulu
	names := [...]string{"Agus", "Fikran", "Candra", "Aidhil", "Hardi", "Dadan", "Noval", "Abdillah"};
	
	// kemudian buat tipe data slice, dan mengisi datanya dari tipe data array names 
	// names[low : high]
	// low merupakan data yang akan di ambil dari array di mulai dari data yang ke dua : 0, 1, 2 ( "Candra" )
	// high merupakan data yang akan di ambil dari array di mulai dari data yang ke lima kurang satu : ("Hardi")
	// jadi isi data sliceNames di bawah ini adalah {"Candra", "Aidhil", "Hardi"}
	slicesNames := names[2:5];
	fmt.Println(slicesNames)

	

	// coba sebutkan isi data yang ada di dalam slice oldName di bawah ini 
	// oldNames := names[0 : 3];

	/*
	  	tipe data slice merupakan tipe data reference, yang artinya jika ada slice baru yang terbentuk dari slice yang lama, maka alamat size data slice
		yang baru akan sama dengan slice yang lama. yang apabila jika terjadi perubahan pada element slice baru, maka akan berdampak juga pada element slice yang lama.
		di karenakan 2 data slice tersebut memiliki refensi data array yang sama dan ketika salah satu element array tersebut di ubah, maka data slice pun akan berubah.

		contohnya seperti di bawah ini : 
	*/

	animals := [...]string{"Hourse", "Cat", "Cow", "Ant", "Zebra", "Crocodile", "Fish", "Snake"};
	fmt.Println("the Animals:", animals);

	// buat slice theAnimals
	theAnimals := animals[0:4];
	waterAnimals := animals[5:7]

	//[Hourse Cat Cow Ant]
	fmt.Println(theAnimals);

	//[Crocodile Fish]
	fmt.Println(waterAnimals)

	// jika salah satu element data di dalam array animals di ubah, maka data slice nya pun akan berubah. contohnya : 
	animals[5] = "Plankton";
	animals[0] = "Dog";

	//[Dog Cat Cow Ant]
	fmt.Println(theAnimals);

	//[Plankton Fish]
	fmt.Println(waterAnimals)

	// jika data slice di ubah, maka variable array yang menjadi referensi slice akan berubah juga.
	waterAnimals[1] = "BabyCrab";

	fmt.Println(animals);

	fmt.Println("-------------")

	// beberapa function yang dapat di gunakan di tipe data slice

	// fungsi len(), untuk mendapatkan panjang dari sebuah slice 
	fmt.Println(len(waterAnimals)); 

	// fungsi cup(), untuk mendapatkan kapasitas dari sebuah slice  
	// cap dari waterAnimals ada 3 , di karenakan di hitung dari indeks 5, 6, 7 (3 kapasitas)
	fmt.Println(cap(waterAnimals));

	// fungsi append(slice, data) membuat slice baru, tapi jika kapasitas sudah penuh maka akan membuat array baru.
	var days = [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}

	theDays := days[5 : 7];
	fmt.Println(theDays)

	// menggunakan fungsi append()
	var newDays = append(theDays, "Libur");
	newDays[0] = "Hari";
	// output yang muncul ialah [Sabtu Minggu Libur]
	fmt.Println(newDays)

	// data days tidak berubah walaupun newDays indeks ke 0 di ubah, karena variable newDays akan membuat array baru jika kapasitas theDays sudah penuh
	fmt.Println(days);

	// cara yang lebih produktif membuat data slice. menggunakan fungsi make()
	myNames := make([]string, 2, 6)
	myNames[0] = "Agus"
	myNames[1] = "Kurniawan"
	
	fmt.Println(myNames);
	// len() dari slice myNames adalah 2, dan cap() nya adalah 6
	fmt.Println(len(myNames));
	fmt.Println(cap(myNames));

	// fungsi copy()
	copyNames := make([]string, len(myNames), cap(myNames));
	copy(copyNames, myNames);
	fmt.Println(copyNames);
}