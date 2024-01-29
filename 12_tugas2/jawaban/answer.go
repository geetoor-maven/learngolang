package main

import "fmt"

func main(){
	/*
		JANGAN LANGSUNG KE SINI KALAU BELUM CARI TAU ATAU BERFIKIR CARA PENYELESAIAN TUGASNYA HEHE.



		BERIKUT JAWABAN Tugas 2.
	*/
	fmt.Println("PROGRAM ToDo - List");

	fmt.Println("--------------------");

	// inisialisasi array buat menyimpan data todo list
	var todoLists = make([]string, 10);

	for{

		fmt.Println("1. Menambah toDo List")
		fmt.Println("2. Melihat toDo List")
		fmt.Println("3. Menghapus toDo List")
		fmt.Println("4. Mengupdate toDo List")
		fmt.Println("5. Keluar applikasi")

		var scanInput int8;
		fmt.Print("Masukan Angka yang dipilih : ");
		fmt.Scanln(&scanInput);

		if scanInput == 5 {
			break;
		}

		if scanInput == 1 {
			fmt.Println("--------------------");
			fmt.Println("Menu menambah toDo List");
			fmt.Print("Masukan To-Do : ")
			var todo string;
			fmt.Scanln(&todo);

			for i := 0; i < len(todoLists); i++ {
				if todoLists[i] == "" {
					todoLists[i] = todo;
					break;
				}
			}
			fmt.Println("--------------------");
		}else if scanInput == 2 {
			fmt.Println("--------------------");
			fmt.Println("Menu melihat toDo List");

			for i := 0; i < len(todoLists); i++ {
				if todoLists[i] != "" {
					fmt.Println((i+1), ":", todoLists[i])	
				}
			}
			fmt.Println("--------------------");
		}
		
	}
}