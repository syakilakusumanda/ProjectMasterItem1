package Model

import (
	"ProjectMasterItem/Database"
	"ProjectMasterItem/Node"
)

func InsertItem(id int, nama string, alamat string, point float32) {
	newDataItem := Node.TableItem{
		Data: Node.Item{id, nama, alamat, point},
		Next: nil,
	}
	var tempLL *Node.TableItem
	tempLL = &Database.DBitem
	if Database.DBitem.Next == nil {
		Database.DBitem.Next = &newDataItem
	} else 
		for tempLL.Next != nil {
			tempLL = tempLL.Next
		}
		tempLL.Next = &newDataItem
	}
}

func ReadAllItem() *Node.TableItem {
	var tempLL *Node.TableItem
	tempLL = &Database.DBitem
	if Database.DBitem.Next == nil {
		return nil
	} else {
		return tempLL
	}
}

func DeleteItem(id int) bool {
	var tempLL *Node.TableItem
	tempLL = &Database.DBitem
	if tempLL.Next == nil {
		return false
	} else {
		for tempLL.Next != nil {
			if tempLL.Next.Data.Id == id {
				tempLL.Next = tempLL.Next.Next
				return true
			}
			tempLL = tempLL.Next
		}
		return false
	}
}

func SearchItem(id int) *Node.TableItem {
	var tempLL *Node.TableItem
	tempLL = Database.DBitem.Next
	cek := false
	if Database.DBitem.Next == nil {
		return nil
	} else {
		for tempLL != nil {
			if id == tempLL.Data.Id {
				cek = true
				return tempLL
			}
			tempLL = tempLL.Next
		}
		if cek != true {
			return nil
		}
	}
	return nil
}

func UpdateItem(item Node.Item) bool {
	addr := SearchItem(item.Id)
	if addr == nil {
		return false
	} else {
		addr.Data.Nama = item.Nama
		addr.Data.Alamat = item.Alamat
		return true
	}
}
