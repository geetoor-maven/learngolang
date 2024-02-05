package main

import "fmt"

func main(){
	/*
		VARIADIC FUNCTION.

		Variadic Function Merupakan kemampuan sebuah function untuk memiliki parameter yang di jadikan sebuah var args. yang artinya 
		bisa menerima lebih dari satu inputan. hampir mirip dengan tipe data array yaa..

		note : parameter nya harus berada di posisi terakhir.

		sebelumnya saya sebutkan kalau var args ini hampir mirip dengan tipe data array.
		jadi apa bedanya ya dengan array ? 
		- jika sebuah function menggunakan parameter yang bertipe array, maka wajib membuat array terlebih dahulu sebelum mengirim datanya ke dalam function tersebut
		- tapi jika sebuah function menggunakan parameter var args, maka kita dapat langsung mengirim datanya tersebut dan bisa lebih dari satu, cukup menggunakan tanda koma

	*/


	// memanggil variadic function
	numbers := sumNumbers(10, 5, 5, 10)
	fmt.Println(numbers)

	theNames("Budi", "Ninja", "Go")

	theNamesAndProgrammingLenguange("Gee", "golang", "java", "Dart")
	fmt.Println()

	// namun jika anda sudah memiliki sebuah variable yang bertipe slice, maka bisa di kirimkan ke variadic function
	// contohnya seperti ini 
	theNumbers := []int{12, 12, 12}
	total := sumNumbers(theNumbers...)
	fmt.Println(total)
}


// contoh sebuah variadic function 
// memiliki parameter var args ( cirinya memiliki titik tiga ...)
func sumNumbers(numbers ...int)int{
	total := 0;

	for _, number := range numbers{
		total += number;
	} 

	return total;
}

func theNames(names ...string){

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

}

// ingat, ketika ingin menggunakan variadic function, parameter var args nya harus yang paling akhir.
func theNamesAndProgrammingLenguange(name string, lengunges ...string){
	fmt.Println("Your name,", name)
	fmt.Print("Your lenguange : ")
	for _, lenguange := range lengunges{
		fmt.Print(lenguange, ",")
	}
}