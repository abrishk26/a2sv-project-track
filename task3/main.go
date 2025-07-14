package main

import (	
	"os"
	"bufio"
	"strconv"
	"fmt"
	"strings"

	"github.com/abrishk26/a2sv-project-track/task3/services"
	"github.com/abrishk26/a2sv-project-track/task3/controllers"
	"github.com/olekukonko/tablewriter"
)

func main() {
	data := [][]string{
		{"Add a new Book", "1"},
		{"Add a new Member", "2"},
		{"Remove an existing Book", "3"},
		{"Borrow a book", "4"},
		{"Return a book", "5"},
		{"List all available books", "6"},
		{"List all borrowed books by a member", "7"},
		{"Exit", "8"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Action", "Input"})
	table.Bulk(data[0:])
	table.Render()

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	choice, err := strconv.Atoi(strings.Trim(input, "\n"))
	if err != nil {
		fmt.Println(err)
		return
	}
	
	controller := &controllers.LibraryController{services.NewLibrary(), reader, table}
	for choice != 8 {

		res, err := controller.HandleInput(choice)
		if err != nil {
			fmt.Println(err)
			return
		}

		if res != "" {
			fmt.Println(res)
		}

	
		table := tablewriter.NewWriter(os.Stdout)
		table.Header([]string{"Action", "Input"})
		table.Bulk(data[0:])
		table.Render()

		
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		choice, err = strconv.Atoi(strings.Trim(input, "\n"))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
