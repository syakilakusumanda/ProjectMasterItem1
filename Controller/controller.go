package Controller

import (
	"ProjectMasterItem/Model"
	"ProjectMasterItem/Node"
)

func InsertItem(id int, nama string, stok int, harga int) bool {
	if id > 0 && nama != "" {
		Model.InsertItem(id, nama, stok, harga)
		return true
	} else {
		return false
	}
}

func ReadAllItem() *Node.TableItem {
	return Model.ReadAllItem()
}
