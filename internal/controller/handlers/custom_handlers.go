package handlers

import "github.com/yehormironenko/reseller/internal/service"

// /search?book=book-name&author=x-author&genre=fiction
type HandlerWithService struct {
	BookService service.BookService
}
