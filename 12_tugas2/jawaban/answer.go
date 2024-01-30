package main

import "fmt"

func main(){
	/*
		JANGAN LANGSUNG KE SINI KALAU BELUM CARI TAU ATAU BERFIKIR CARA PENYELESAIAN TUGASNYA HEHE.



		BERIKUT JAWABAN Tugas 2.
	*/
	fmt.Println("PROGRAM ToDo - List");

	fmt.Println("--------------------");

	// inisialisasi array slice buat menyimpan data todo list
	 todoLists := []string{};

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

			// menambahkan data todo ke todoLists menggunakan append function
			todoLists = append(todoLists, todo);
			
			fmt.Println("--------------------");
		}else if scanInput == 2 {
			fmt.Println("--------------------");
			fmt.Println("Menu melihat toDo List");
			
			// looping data slice untuk melihat semua data to do list
			for i := 0; i < len(todoLists); i++ {
				if todoLists[i] != "" {
					fmt.Println("- ", ":", todoLists[i]);
				}
			}

			
			fmt.Println("--------------------");
		}else if scanInput == 3 {
			fmt.Println("--------------------");
			fmt.Println("Menu menghapus toDo List");

			fmt.Print("Masukan To-Do yang ingin di hapus : ")
			var todo string;
			fmt.Scanln(&todo);

			isEmpty := true;

			for i := 0; i < len(todoLists); i++ {
				if todoLists[i] == todo {
					todoLists[i] = "";
					fmt.Println("Berhasil menghapus todo list")
					fmt.Println("--------------------");
					isEmpty = false;
					break	
				}
			}

			if isEmpty {
				fmt.Println("Todo List tidak di temukan.")
				fmt.Println("--------------------");
			}
			
		}else if scanInput == 4{
			fmt.Println("--------------------");
			fmt.Println("Menu mengupdate toDo List");

			fmt.Print("Masukan To-Do yang ingin di update : ")
			var todo string;
			fmt.Scanln(&todo);

			isEmpty := true;

			for i := 0; i < len(todoLists); i++ {
				if todoLists[i] == todo {
					fmt.Print("Masukan To-Do baru : ")
					var todoUpdate string;
					fmt.Scanln(&todoUpdate);
					todoLists[i] = todoUpdate;

					isEmpty = false;
					break	
				}
			}

			if isEmpty {
				fmt.Println("Todo List tidak di temukan.")
				fmt.Println("--------------------");
			}

		}else {
			fmt.Println("Feature tidak di temukan.")
		}
		
	}
}