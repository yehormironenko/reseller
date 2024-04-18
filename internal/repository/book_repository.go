package repository

import "reseller/internal/repository/entities"

type BookRepositoryInterface interface {
	GetBookByParams(...string) (entities.Book, error)
	//buy
}
