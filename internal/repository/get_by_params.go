package repository

import (
	"fmt"

	"github.com/uptrace/bun"

	"reseller/internal/repository/entities"
)

type BookRepository struct {
	db *bun.DB
}

func NewBookRepository(db *bun.DB) BookRepository {
	return BookRepository{
		db: db,
	}
}

func (br BookRepository) GetBookByParams(args ...string) (entities.Book, error) {
	fmt.Println("I'm in repos")
	fmt.Println(args)
	return entities.Book{}, nil
}
