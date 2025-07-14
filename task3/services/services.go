package services

import (
	"errors"

	"github.com/abrishk26/a2sv-project-track/task3/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID, memberID int) error
	ReturnBook(bookID, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

func NewLibrary() *Library {
	return &Library{
		map[int]models.Book{},
		map[int]models.Member{},
		map[int]map[int]bool{},
	}
}

type Library struct {
	Books map[int]models.Book
	Members map[int]models.Member
	Borrow map[int]map[int]bool
}

func (l *Library) AddBook(book models.Book) {
	l.Books[book.ID] = book
}

func (l *Library) RemoveBook(bookID int) {
	delete(l.Books, bookID)
}

func (l *Library) BorrowBook(bookID, memberID int) error {
	book, exist := l.Books[bookID]
	if !exist {
		return errors.New("Book with the given ID does not exist")
	}

	member, exist := l.Members[memberID]
	if !exist {
		return errors.New("Member with the given ID does not exist")
	}

	l.Borrow[bookID][memberID] = true

	return nil
}

func (l *Library) ReturnBook(bookID, memberID int) error {
	book, exist := l.Borrow[bookID]
	if !exist {
		return errors.New("No has never been borrowed")
	}

	member, exist := l.Borrow[book][memberID]
	if !exist {
		return errors.New("member didn't borrow this book")
	}

	delete(l.Borrow[book][member])

	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	books := []models.Book{}

	for _, book := range l.Books {
		books = append(books, book)
	}

	return books
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	books := []models.Book{}

	for bookID := range l.Borrow {
		if _, exist := l.Borrow[bookID][memberID]; exist {
			books = append(books, l.Books[bookID])
		}
	}
}
