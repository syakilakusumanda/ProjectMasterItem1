package main

import (
	"ProjectMasterItem/View"
	"fmt"
)

func menu() {
	for {
		fmt.Println("Menu Program")
		fmt.Println("1. Insert Item")
		fmt.Println("2. Update Item")
		fmt.Println("3. Delete Item")
		fmt.Println("4. Read All Item")
		fmt.Println("5. Search Item")
		fmt.Println("6. Exit")
		fmt.Println("----------------")
		fmt.Print("Input menu pilihan anda: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			View.InsertItem()
		case 2:
			//updateItem()
		case 3:
			//deleteItem()
		case 4:
			View.ReadAllItem()
		case 5:
			//searchItem()
		case 6:
			fmt.Println("exit program")
			return
		}
	}
}

func main() {
	menu()
}
