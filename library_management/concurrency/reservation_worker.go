package concurrency

import (
	"fmt"
	"library_management/services"
	"sync"
	"time"
)

type Reservation struct {
	BookID   uint
	MemberID uint
	Response chan error
}

type ReservationWorker struct {
	LibraryService services.LibraryManager
	Reservations   chan Reservation
	Mutex          *sync.Mutex
}

func NewReservationWorker(service services.LibraryManager) *ReservationWorker {
	return &ReservationWorker{
		LibraryService: service,
		Reservations:   make(chan Reservation),
		Mutex:          &sync.Mutex{},
	}
}

func (w *ReservationWorker) Start() {
	go func() {
		for reservation := range w.Reservations {
			go w.processReservation(reservation)
		}
	}()
}

func (w *ReservationWorker) processReservation(reservation Reservation) {
	w.Mutex.Lock()
	defer w.Mutex.Unlock()

	err := w.LibraryService.ReserveBook(reservation.BookID, reservation.MemberID)
	reservation.Response <- err

	if err == nil {
		go w.cancelReservationAfterTimeout(reservation.BookID)
	}
}

func (w *ReservationWorker) cancelReservationAfterTimeout(bookID uint) {
	time.Sleep(5 * time.Second)

	w.Mutex.Lock()
	defer w.Mutex.Unlock()

	book := w.LibraryService.GetBook(bookID)
	if book != nil && book.Status == "Reserved" {
		w.LibraryService.UnreserveBook(bookID)
		fmt.Printf("Reservation for book ID %d has been canceled due to timeout.\n", bookID)
	}
}
