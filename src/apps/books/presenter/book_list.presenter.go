package presenter

import "github.com/PickHD/bookstore-api/src/apps/books/model"

type BookListPresenter struct {
	UUID            string `json:"uuid"`
	Title           string `json:"title"`
	ISBN            string `json:"isbn"`
	Year            string `json:"year"`
	Price           int    `json:"price"`
	NoPages         int    `json:"no_pages"`
	BookDescription string `json:"book_description"`
}

func NewBookListPresenter() BookListPresenter {
	return BookListPresenter{}
}

func (bookList BookListPresenter) Build(model []model.Books, total int64) (list []BookListPresenter) {
	if total > 0 {
		for _, val := range model {
			listItems := BookListPresenter{
				UUID:            val.UUID,
				Title:           val.Title,
				ISBN:            val.ISBN,
				Year:            val.Year,
				Price:           val.Price,
				NoPages:         val.NoPages,
				BookDescription: val.BookDescription,
			}

			list = append(list, listItems)
		}
	} else {
		list = []BookListPresenter{}
	}

	return

}
