package Node

type Item struct {
	Id     int
	Nama   string
	Alamat string
	Point  float32
}

type TableItem struct {
	Data Item
	Next *TableItem
}
