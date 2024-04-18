package model

type Book struct {
	Bookname string  `json:"bookname"`
	Author   string  `json:"author"`
	Genre    string  `json:"genre"`
	Year     int     `json:"year"`
	Count    int     `json:"count"`
	Price    float32 `json:"price"`
}
