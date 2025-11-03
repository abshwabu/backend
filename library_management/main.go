package main

import (
	"fmt"
	"library_management/concurrency"
	"library_management/controllers"
	"library_management/models"
	"library_management/services"
	"sync"
)

func main() {
	libraryService := services.NewLibrary()
	reservationWorker := concurrency.NewReservationWorker(libraryService)
	reservationWorker.Start()

	libraryController := controllers.NewLibraryController(libraryService)

	libraryService.AddBook(models.Book{ID: 1, Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams"})
	libraryService.AddBook(models.Book{ID: 2, Title: "1984", Author: "George Orwell"})
	libraryService.AddBook(models.Book{ID: 3, Title: "Dune", Author: "Frank Herbert"})
	libraryService.AddMember(models.Member{ID: 101, Name: "Alice"})
	libraryService.AddMember(models.Member{ID: 102, Name: "Bob"})

	// Simulate concurrent reservations
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		response := make(chan error)
		reservationWorker.Reservations <- concurrency.Reservation{
			BookID:   1,
			MemberID: 101,
			Response: response,
		}
		err := <-response
		if err != nil {
			fmt.Println("Reservation failed for member 101:", err)
		} else {
			fmt.Println("Reservation successful for member 101")
		}
	}()

	go func() {
		defer wg.Done()
		response := make(chan error)
		reservationWorker.Reservations <- concurrency.Reservation{
			BookID:   1,
			MemberID: 102,
			Response: response,
		}
		err := <-response
		if err != nil {
			fmt.Println("Reservation failed for member 102:", err)
		} else {
			fmt.Println("Reservation successful for member 102")
		}
	}()

	wg.Wait()

	libraryController.Run()
}