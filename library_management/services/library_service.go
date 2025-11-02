package services

import (
	"errors"
	"fmt"
	"library_management/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID uint)
	BorrowBook(bookID uint, memberID uint) error
	ReturnBook(bookID uint, memberID uint) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID uint) []models.Book
}

type Library struct {
	Books   map[uint]models.Book
	Members map[uint]models.Member
}

func NewLibrary() *Library {
	return &Library{
		Books:   make(map[uint]models.Book),
		Members: make(map[uint]models.Member),
	}
}

func (l *Library) AddBook(book models.Book) {
	if _, exists := l.Books[book.ID]; exists {
		fmt.Printf("Book with ID %d already exists.\n", book.ID)
		return
	}
	book.Status = "Available"
	l.Books[book.ID] = book
	fmt.Printf("Book '%s' added to the library.\n", book.Title)
}

func (l *Library) RemoveBook(bookID uint) {
	if book, exists := l.Books[bookID]; exists {
		if book.Status == "Borrowed" {
			fmt.Printf("Cannot remove a borrowed book (ID: %d).\n", bookID)
			return
		}
		delete(l.Books, bookID)
		fmt.Printf("Book ID %d removed from the library.\n", bookID)
	} else {
		fmt.Printf("Book with ID %d not found.\n", bookID)
	}
}

func (l *Library) BorrowBook(bookID uint, memberID uint) error {
	book, bookExists := l.Books[bookID]
	if !bookExists {
		return errors.New("book not found")
	}

	member, memberExists := l.Members[memberID]
	if !memberExists {
		return errors.New("member not found")
	}

	if book.Status == "Borrowed" {
		return errors.New("book is already borrowed")
	}

	book.Status = "Borrowed"
	l.Books[bookID] = book

	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.Members[memberID] = member

	fmt.Printf("Member %s borrowed book '%s'.\n", member.Name, book.Title)
	return nil
}

func (l *Library) ReturnBook(bookID uint, memberID uint) error {
	book, bookExists := l.Books[bookID]
	if !bookExists {
		return errors.New("book not found")
	}

	member, memberExists := l.Members[memberID]
	if !memberExists {
		return errors.New("member not found")
	}

	if book.Status == "Available" {
		return errors.New("book was not borrowed")
	}

	book.Status = "Available"
	l.Books[bookID] = book

	for i, borrowedBook := range member.BorrowedBooks {
		if borrowedBook.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			break
		}
	}
	l.Members[memberID] = member

	fmt.Printf("Member %s returned book '%s'.\n", member.Name, book.Title)
	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	var available []models.Book
	for _, book := range l.Books {
		if book.Status == "Available" {
			available = append(available, book)
		}
	}
	return available
}

func (l *Library) ListBorrowedBooks(memberID uint) []models.Book {
	if member, exists := l.Members[memberID]; exists {
		return member.BorrowedBooks
	}
	return nil
}

func (l *Library) AddMember(member models.Member) {
    if _, exists := l.Members[member.ID]; exists {
        fmt.Printf("Member with ID %d already exists.\n", member.ID)
        return
    }
    l.Members[member.ID] = member
    fmt.Printf("Member %s added to the system.\n", member.Name)
}
