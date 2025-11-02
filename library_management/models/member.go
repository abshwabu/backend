package models

type Member struct {
	ID uint 
	Name string
	BorrowedBooks []Book
}
