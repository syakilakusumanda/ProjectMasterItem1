package Controller

import (
	"ProjectMasterItem/Model"
	"ProjectMasterItem/Node"
)

func InsertItem(id int, nama string, alamat string, point float32) bool {
	if id > 0 && nama != "" {
		Model.InsertItem(id, nama, alamat, point)
		return true
	} else {
		return false
	}
}

func ReadAllItem() *Node.TableItem {
	return Model.ReadAllItem()
}
