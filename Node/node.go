package Node

type Item struct {
	Id     int
	Nama   string
	Stok   int
	Harga  int
}

type TableItem struct {
	Data Item
	Next *TableItem
}
