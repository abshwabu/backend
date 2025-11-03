# Library Management System - Concurrency

## Concurrent Book Reservation System

The Library Management System has been updated to support concurrent book reservations using Goroutines, Channels, and Mutexes.

### Concurrency Implementation

- **Goroutines**: Each reservation request is processed in a separate Goroutine, allowing multiple requests to be handled simultaneously.
- **Channels**: A channel is used to queue incoming reservation requests, ensuring that they are processed in a controlled manner.
- **Mutex**: A `sync.Mutex` is used to prevent race conditions when accessing and modifying the book availability status.

### Reservation Workflow

1. A reservation request is sent to the `Reservations` channel.
2. The `ReservationWorker` processes the request in a new Goroutine.
3. The `ReserveBook` method is called, which locks the mutex to ensure exclusive access to the book's status.
4. If the book is available, its status is changed to "Reserved".
5. A timer-based Goroutine is started to automatically un-reserve the book after 5 seconds if it is not borrowed.

### Auto-Cancellation

If a reserved book is not borrowed within 5 seconds, it is automatically unreserved. This is handled by a separate Goroutine that sleeps for 5 seconds and then checks if the book is still in the "Reserved" state. If it is, the book's status is changed back to "Available".