package entities

import "github.com/uptrace/bun"

type Book struct {
	bun.BaseModel `bun:"table:books,alias:b"`

	ID       int64  `bun:",pk,autoincrement"`
	Bookname string `bun:"book_name"`
	Author   string
	Genre    string
	Year     int
	Count    int
	Price    float32
}
