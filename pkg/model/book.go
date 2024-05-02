package model

type Books []Book

type Book struct {
	Bookname string  `json:"bookname"`
	Author   string  `json:"author,omitempty"`
	Genre    string  `json:"genre,omitempty"`
	Year     int     `json:"year,omitempty"`
	Count    int     `json:"count,omitempty"`
	Price    float32 `json:"price,omitempty"`
}
