package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"library_management/services"
	"os"
	"strconv"
	"strings"
)

type LibraryController struct {
	Service services.LibraryManager
	Reader  *bufio.Reader
}

func NewLibraryController(service services.LibraryManager) *LibraryController {
	return &LibraryController{
		Service: service,
		Reader:  bufio.NewReader(os.Stdin),
	}
}

func (c *LibraryController) Run() {
	for {
		c.displayMenu()
		choice := c.readInput()
		switch choice {
		case "1":
			c.addBook()
		case "2":
			c.removeBook()
		case "3":
			c.borrowBook()
		case "4":
			c.returnBook()
		case "5":
			c.listAvailableBooks()
		case "6":
			c.listBorrowedBooks()
        case "7":
            c.addMember()
			fmt.Println("Exiting Library Management System. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func (c *LibraryController) displayMenu() {
	fmt.Println("\n--- Library Management System Menu ---")
	fmt.Println("1. Add a new book")
	fmt.Println("2. Remove an existing book")
	fmt.Println("3. Borrow a book")
	fmt.Println("4. Return a book")
	fmt.Println("5. List all available books")
	fmt.Println("6. List all borrowed books by a member")
    fmt.Println("7. Add a new member (Helper)")
	fmt.Println("q. Quit")
	fmt.Print("Enter your choice: ")
}

func (c *LibraryController) readInput() string {
	input, _ := c.Reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func (c *LibraryController) readUintInput(prompt string) (uint, error) {
    fmt.Print(prompt)
    inputStr := c.readInput()
    val, err := strconv.ParseUint(inputStr, 10, 64)
    if err != nil {
        return 0, err
    }
    return uint(val), nil
}

func (c *LibraryController) addBook() {
	id, err := c.readUintInput("Enter book ID: ")
	if err != nil {
		fmt.Println("Invalid ID.")
		return
	}
	fmt.Print("Enter book Title: ")
	title := c.readInput()
	fmt.Print("Enter book Author: ")
	author := c.readInput()

	book := models.Book{
		ID:     id,
		Title:  title,
		Author: author,
        Status: "Available",
	}
	c.Service.AddBook(book)
}

func (c *LibraryController) removeBook() {
	id, err := c.readUintInput("Enter book ID to remove: ")
	if err != nil {
		fmt.Println("Invalid ID.")
		return
	}
	c.Service.RemoveBook(id)
}

func (c *LibraryController) borrowBook() {
	bookID, err := c.readUintInput("Enter book ID to borrow: ")
	if err != nil {
		fmt.Println("Invalid Book ID.")
		return
	}
	memberID, err := c.readUintInput("Enter member ID: ")
	if err != nil {
		fmt.Println("Invalid Member ID.")
		return
	}

	if err := c.Service.BorrowBook(bookID, memberID); err != nil {
		fmt.Println("Error borrowing book:", err)
	}
}

func (c *LibraryController) returnBook() {
	bookID, err := c.readUintInput("Enter book ID to return: ")
	if err != nil {
		fmt.Println("Invalid Book ID.")
		return
	}
	memberID, err := c.readUintInput("Enter member ID: ")
	if err != nil {
		fmt.Println("Invalid Member ID.")
		return
	}

	if err := c.Service.ReturnBook(bookID, memberID); err != nil {
		fmt.Println("Error returning book:", err)
	}
}

func (c *LibraryController) listAvailableBooks() {
	books := c.Service.ListAvailableBooks()
	fmt.Println("\n--- Available Books ---")
	if len(books) == 0 {
		fmt.Println("No available books.")
		return
	}
	for _, book := range books {
		fmt.Printf("ID: %d, Title: '%s', Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func (c *LibraryController) listBorrowedBooks() {
	memberID, err := c.readUintInput("Enter member ID: ")
	if err != nil {
		fmt.Println("Invalid Member ID.")
		return
	}
    
    if _, exists := c.Service.(*services.Library).Members[memberID]; !exists {
        fmt.Printf("Member with ID %d not found.\n", memberID)
        return
    }

	books := c.Service.ListBorrowedBooks(memberID)
	fmt.Printf("\n--- Books Borrowed by Member ID %d ---\n", memberID)
	if len(books) == 0 {
		fmt.Println("No books borrowed.")
		return
	}
	for _, book := range books {
		fmt.Printf("ID: %d, Title: '%s', Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func (c *LibraryController) addMember() {
    id, err := c.readUintInput("Enter member ID: ")
    if err != nil {
        fmt.Println("Invalid ID.")
        return
    }
    fmt.Print("Enter member Name: ")
    name := c.readInput()

    member := models.Member{
        ID:   id,
        Name: name,
    }
    c.Service.(*services.Library).AddMember(member)
}
