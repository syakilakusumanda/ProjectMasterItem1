package View

import (
	"ProjectMasterItem/Controller"
	"fmt"
)

func InsertItem() {
	var id int
	var nama, alamat string
	var point float32
	fmt.Print("masukkan id item : ")
	fmt.Scan(&id)
	fmt.Print("masukkan nama item : ")
	fmt.Scan(&nama)
	fmt.Print("masukkan alamat item : ")
	fmt.Scan(&alamat)
	fmt.Print("masukkan point item : ")
	fmt.Scan(&point)
	cek := Controller.InsertItem(id, nama, alamat, point)
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
			fmt.Println("Alamat Item : ", items.Next.Data.Alamat)
			items = items.Next
		}
	}
}
