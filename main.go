package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func loadBooks(filename string) []Book {
	file, err := os.Open(filename)
	if err != nil {
		return []Book{} // No file yet? Start empty
	}
	defer file.Close()

	var books []Book
	err = json.NewDecoder(file).Decode(&books)
	if err != nil {
		fmt.Println("Error decoding file:", err)
		return []Book{}
	}
	return books
}
func saveBooks(filename string, books []Book) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error saving file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(books)
	if err != nil {
		fmt.Println("Error writing JSON:", err)
	}
}

func main() {
	books := []Book{}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("--pick an option--")
	fmt.Println("1-Add a book")
	fmt.Println("2-View the list of books")
	fmt.Println("3-Search by title")
	fmt.Println("4-delete by title")
	fmt.Println("5-save to a file")
	fmt.Println("6-load to a file")
	for {
		scanner.Scan()
		option := scanner.Text()
		switch option {
		case "1":
			fmt.Println("add the title of the book")
			scanner.Scan()
			title := scanner.Text()
			fmt.Println("Enter the name of the author")
			scanner.Scan()
			author := scanner.Text()
			fmt.Println("Enter the pages:")
			scanner.Scan()
			pagesStr := scanner.Text()
			pages, err := strconv.Atoi(pagesStr)
			if err != nil {
				fmt.Println("Invalid number of pages. Please enter a valid integer.")
				break
			}
			book := Book{Title: title, Author: author, Pages: pages}
			books = append(books, book)
			fmt.Println("the book was added successfully!")
		case "2":
			if len(books) == 0 {
				fmt.Println("No books to view yet")
				break
			}
			for i, e := range books {
				fmt.Printf("%d - Title: %s, Author: %s\n", i+1, e.Title, e.Author)
			}
		case "3":
			fmt.Println("Enter title to search")
			scanner.Scan()
			searchTitle := scanner.Text()
			found := false
			for _, book := range books {
				if strings.EqualFold(book.Title, searchTitle) {
					fmt.Printf("Book found! Title: %s, Author: %s\n", book.Title, book.Author)
					found = true
					break
				}
			}
			if !found {
				fmt.Println("The book is not found")
			}
		case "4":
			fmt.Println("Enter the title to delete")
			scanner.Scan()
			deleteTitle := scanner.Text()
			found := false
			for i, book := range books {
				if strings.EqualFold(book.Title, deleteTitle) {
					books = append(books[:i], books[i+1:]...)
					found = true
					fmt.Println("Book deleted!")
					break
				}
			}
			if !found {
				fmt.Println("Not found")
			}
		case "5":
			saveBooks("books.json", books)
			fmt.Println("Books saved successfully!")
		case "6":
			books = loadBooks("books.json")
			fmt.Println("Books loaded successfully!")
		default:
			fmt.Println("Invalid option. Please try again.")
		}

	}
}
