package View

import (
	"ProjectMasterItem/Controller"
	"fmt"
)

func InsertItem() {
	var id, stok, harga int
	var nama string
	fmt.Print("input id item : ")
	fmt.Scan(&id)
	fmt.Print("input nama item : ")
	fmt.Scan(&nama)
	fmt.Print("input stok item : ")
	fmt.Scan(&alamat)
	fmt.Print("input harga item : ")
	fmt.Scan(&point)
	cek := Controller.InsertItem(id, nama, stok, harga)
	if cek == true {
		fmt.Println("data berhasil di input")
	} else {
		fmt.Println("data gagal diinput")
	}
}

func ReadAllItem() {
	items := Controller.ReadAllItem()
	if items != nil {
		for items.Next != nil {
			fmt.Println("Nama Item : ", items.Next.Data.Nama)
			fmt.Println("Id Item : ", items.Next.Data.Id)
			fmt.Println("Stok Item : ", items.Next.Data.Alamat)
			items = items.Next
		}
	}
}
