package presenter

import "github.com/PickHD/bookstore-api/src/apps/books/model"

type BookDetailPresenter struct {
	UUID            string `json:"uuid"`
	Title           string `json:"title"`
	ISBN            string `json:"isbn"`
	Year            string `json:"year"`
	Price           int    `json:"price"`
	NoPages         int    `json:"no_pages"`
	BookDescription string `json:"book_description"`
}

func NewBookDetailPresenter() BookDetailPresenter {
	return BookDetailPresenter{}
}

func (bookDetail BookDetailPresenter) Build(model model.Books) (detail BookDetailPresenter) {
	newBookDetail := BookDetailPresenter{
		UUID:            model.UUID,
		Title:           model.Title,
		ISBN:            model.ISBN,
		Year:            model.Year,
		Price:           model.Price,
		NoPages:         model.NoPages,
		BookDescription: model.BookDescription,
	}

	return newBookDetail
}
