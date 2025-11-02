package main

import (
	"library_management/controllers"
	"library_management/services"
	"library_management/models"
)

func main() {
	libraryService := services.NewLibrary()

	libraryController := controllers.NewLibraryController(libraryService)

    libraryService.AddBook(models.Book{ID: 1, Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams"})
    libraryService.AddBook(models.Book{ID: 2, Title: "1984", Author: "George Orwell"})
    libraryService.AddBook(models.Book{ID: 3, Title: "Dune", Author: "Frank Herbert"})
    libraryService.AddMember(models.Member{ID: 101, Name: "Alice"})
    libraryService.AddMember(models.Member{ID: 102, Name: "Bob"})

	libraryController.Run()
}

