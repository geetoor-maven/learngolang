package main

import "fmt"

func main(){
	/*
		JANGAN LANGSUNG KE SINI KALAU BELUM CARI TAU ATAU BERFIKIR CARA PENYELESAIAN TUGASNYA HEHE.



		BERIKUT JAWABAN Tugas 3.
	*/
	viewOutput("PROGRAM ToDo - List")

	line()

	// inisialisasi array slice buat menyimpan data todo list
	todoLists := []string{};

	for{

		viewFeature()

		var scanInput int8;
		fmt.Print("Masukan Angka yang dipilih : ");

		fmt.Scanln(&scanInput);

		if scanInput == 5 {
			break;
		}

		if scanInput == 1 {
			line()
			viewOutput("Menu menambah toDo List")
			var todo = inputUser("Masukan To-Do : ");
			todoLists = append(todoLists, todo);
			line()
		}else if scanInput == 2 {
			line()
			viewOutput("Menu melihat toDo List")
			viewDataList(todoLists)
			line()
		}else if scanInput == 3 {
			line()
			viewOutput("Menu menghapus toDo List")
			var todo = inputUser("Masukan To-Do yang ingin di hapus : ");
			isEmpty := updateOrDeleteDataList(todoLists, todo, "delete");
			if isEmpty {
				viewOutput("Todo List tidak di temukan.")
				line()
			}	
		}else if scanInput == 4{
			line()
			viewOutput("Menu mengupdate toDo List")
			var todo = inputUser("Masukan To-Do yang ingin di update : ");
			isEmpty := updateOrDeleteDataList(todoLists, todo, "update");
			if isEmpty {
				viewOutput("Todo List tidak di temukan.")
				line()
			}
		}else {
			viewOutput("Feature tidak di temukan.")
		}
		
	}
	
}

func viewDataList(todoLists []string){
	for i := 0; i < len(todoLists); i++ {
		if todoLists[i] != "" {
			fmt.Println("- ", ":", todoLists[i]);
		}
	}

}

func updateOrDeleteDataList(todoLists []string,  todo string, code string)bool{
	isEmpty := true;

	if code == "delete" {
		for i := 0; i < len(todoLists); i++ {
			if todoLists[i] == todo {
				todoLists[i] = "";
				viewOutput("Berhasil menghapus todo list")
				line()
				isEmpty = false;
				break	
			}
		}
	}else if code == "update" {
		for i := 0; i < len(todoLists); i++ {
			if todoLists[i] == todo {
				var todoUpdate = inputUser("Masukan To-Do baru :");
				todoLists[i] = todoUpdate;

				isEmpty = false;
				break	
			}
		}
	}

	return isEmpty
}

func line(){
	fmt.Println("--------------------");
}

func viewFeature(){
	fmt.Println("1. Menambah toDo List")
	fmt.Println("2. Melihat toDo List")
	fmt.Println("3. Menghapus toDo List")
	fmt.Println("4. Mengupdate toDo List")
	fmt.Println("5. Keluar applikasi")
}

func viewOutput(output string){
	fmt.Println(output)
}

func inputUser(message string)string{
	fmt.Print(message)
	var todo string;
	fmt.Scanln(&todo);
	return todo;
}

