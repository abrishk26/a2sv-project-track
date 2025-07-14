package controllers

import (
	"errors"
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"os"
	
	"github.com/abrishk26/a2sv-project-track/task3/services"
	"github.com/abrishk26/a2sv-project-track/task3/models"
	"github.com/olekukonko/tablewriter"
)

type LibraryController struct {
	L *services.Library
	R *bufio.Reader
	T *tablewriter.Table
}

func (lc *LibraryController) HandleInput(input int) (string, error) {
	switch input {
		case 1:
			fmt.Print("Enter Book Name: ")
			input, err := lc.R.ReadString('\n')
			if err != nil {
				return "", err
			}
			bookName := strings.Trim(input, "\n")

			fmt.Print("Enter Author Name: ")
			input, err = lc.R.ReadString('\n')
			if err != nil {
				return "", err
			}

			authorName := strings.Trim(input, "\n")

			bookID := len(lc.L.Books)
			book := models.Book{bookID, bookName, authorName}
			lc.L.AddBook(book)
			return "Book Added successfully", nil
		case 2:
			fmt.Print("Enter Member Name: ")
			input, err := lc.R.ReadString('\n')
			if err != nil {
				return "", err
			}
			memberName := strings.Trim(input, "\n")

			memberID := len(lc.L.Members)
			member := models.Member{memberID, memberName, []models.Book{}}
			lc.L.AddMember(member)
			return "Member Added successfully", nil
		case 3:
			fmt.Print("Enter Book ID: ")
			input, err := lc.R.ReadString('\n')
			if err != nil {
				return "", err
			}

			bookID, err := strconv.Atoi(strings.Trim(input, "\n"))
			if err != nil {
				return "", err
			}

			if _, ok := lc.L.Books[bookID]; !ok && bookID < 0 {
				return "", errors.New("Book with the given ID does not exist")
			}

			lc.L.RemoveBook(bookID)

			return "Book Removed successfully", nil
		case 4:
			fmt.Print("Enter Book ID: ")
			input, err := lc.R.ReadString('\n')
			if err != nil {
				return "", err
			}

			bookID, err := strconv.Atoi(strings.Trim(input, "\n"))
			if err != nil {
				return "", err
			}

			if _, ok := lc.L.Books[bookID]; !ok && bookID < 0 {
				return "", errors.New("Book with the given ID does not exist")
			}
			
			fmt.Print("Enter Member ID: ")
			input, err = lc.R.ReadString('\n')
			if err != nil {
				return "", err
			}

			memberID, err := strconv.Atoi(strings.Trim(input, "\n"))
			if err != nil {
				return "", err
			}

			if _, err := lc.L.Members[memberID]; !err && memberID < 0 {
				return "", errors.New("Member with the given ID does not exist")
			}

			err = lc.L.BorrowBook(bookID, memberID)
			if err != nil {
				return "", err
			}

			return "Done", nil
		case 5:
			fmt.Print("Enter Book ID: ")
			input, err := lc.R.ReadString('\n')
			if err != nil {
				return "", err
			}

			bookID, err := strconv.Atoi(strings.Trim(input, "\n"))
			if err != nil {
				return "", err
			}

			if _, ok := lc.L.Books[bookID]; !ok && bookID < 0 {
				return "", errors.New("Book with the given ID does not exist")
			}

			
			fmt.Print("Enter Member ID: ")
			input, err = lc.R.ReadString('\n')
			if err != nil {
				return "", err
			}

			memberID, err := strconv.Atoi(strings.Trim(input, "\n"))
			if err != nil {
				return "", err
			}

			if _, err := lc.L.Members[memberID]; !err && memberID < 0 {
				return "", errors.New("Member with the given ID does not exist")
			}

			err = lc.L.ReturnBook(bookID, memberID)
			if err != nil {
				return "", err
			}

			return "Done", nil
		case 6:
			lc.T = tablewriter.NewWriter(os.Stdout)
			lc.T.Header([]string{"Book Name", "Author Name"})
			for _, book := range lc.L.Books {
				lc.T.Append([]string{book.Title, book.Author})
			}

			lc.T.Render()
			return "", nil
		case 7:
			fmt.Print("Enter Member ID: ")
			input, err := lc.R.ReadString('\n')
			if err != nil {
				return "", err
			}

			memberID, err := strconv.Atoi(strings.Trim(input, "\n"))
			if err != nil {
				return "", err
			}

			if _, err := lc.L.Members[memberID]; !err && memberID < 0 {
				return "", errors.New("Member with the given ID does not exist")
			}

			
			books := []models.Book{}
			for bookID, borrowMap := range lc.L.Borrow {
				if _, exist := borrowMap[memberID]; exist {
					books = append(books, lc.L.Books[bookID])
				}
			}
			
			lc.T = tablewriter.NewWriter(os.Stdout)
			lc.T.Header([]string{"Book Name", "Author Name"})
			for _, book := range books {
				lc.T.Append([]string{book.Title, book.Author})
			}

			lc.T.Render()
			return "", nil
		default:
			return "", errors.New("Unsupported input") 
	}
}
