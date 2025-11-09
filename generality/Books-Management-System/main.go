package main
import (
	"fmt"
	"encoding/json"
	"os"
	"strings"
)

type Book struct {
	ID       int `json:"id"`
	Title    string  `json:"title"`
	Author   string  `json:"author"`
	Year     int   `json:"year"`
	IsBorrowed bool  `json:"isBorrowed"`
}

func LoadBooks() (map[int]Book, error) {
	fileName := "books.json"
    data, err := os.ReadFile(fileName)
	if err !=nil || len(data) == 0 {
		if os.IsNotExist(err) {
			file, err := os.Create(fileName)
			if err != nil {
				return nil, err
			}
			defer file.Close()
			books := []Book{}
			booksJson, err := json.MarshalIndent(books, "", "  ")
			if err != nil {
				return nil, err
			}
			_, err = file.Write(booksJson)
			if err != nil {
				return nil, err
			}
			return nil, nil
		}
		return nil, err
	}
	var books []Book
	err = json.Unmarshal(data, &books)
	if err != nil {
		return nil, err
	}
	var booksMap =make(map[int]Book)
	//convert books slice to map
	for _,book :=range books {
       booksMap[book.ID]=book
	}
   return booksMap,nil
}

func SaveBooks(books []Book) error {
	fileName := "books.json"
	booksJson, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(fileName, booksJson, 0644)
	if err != nil {
		return err
	}
	return nil
}

func AddBooks() ([]Book,error){
   booksMap, error:=LoadBooks()	
   if error!=nil{
	fmt.Println("Error Loading Books ❌")
	return nil,error
   }
   newBooks := make([]Book, 0)
   var choice string
   var Id int
   
   for{
	 fmt.Print("Do you want to add book y/n :")
	 fmt.Scan(&choice)
	 switch choice {
case "y", "yes", "Y", "Yes", "YES":
         fmt.Print("Enter the Book Id :")
		 fmt.Scan(&Id)
		 if booksMap[Id].ID==Id{
			fmt.Println("Book ID already exists. Please enter a unique ID.")
			fmt.Scan(&Id)
			if booksMap[Id].ID==Id{
				fmt.Println("Book ID already exists. Please enter a unique ID.")
				return nil,fmt.Errorf("Book ID already exists. Please enter a unique ID.")
			}
		 }
		 newBook := Book{ID: Id}
		 fmt.Print("Enter the Book title :")
		 fmt.Scan(&newBook.Title)
		 fmt.Print("Enter the Book author :")
		 fmt.Scan(&newBook.Author)
		 fmt.Print("Enter the edition year of the book :")
		 fmt.Scan(&newBook.Year)
		 newBook.IsBorrowed=false

		 newBooks = append(newBooks, newBook)
	 case "n", "no", "N", "No", "NO":
		SaveBooks(newBooks)
		return newBooks,nil
	 default:
		fmt.Println("Invalid choice. Please try again.")
	 }
   }
}

func BorrowBook(bookId int) error{
	booksMap, error:=LoadBooks()	
	if error!=nil{
		fmt.Println("Error Loading Books ❌")
		return error
	}
	book, exists := booksMap[bookId]
	if !exists {
		return fmt.Errorf("Book with ID %d does not exist", bookId)
	}

	book.IsBorrowed = true
	booksMap[bookId] = book
	//convaert map to slice
	books := make([]Book, 0)
	for _, book := range booksMap {
		books = append(books, book)
	}
	return SaveBooks(books)
}

func ReturnBook(bookId int) error{
	booksMap, error:=LoadBooks()	
	if error!=nil{
		fmt.Println("Error Loading Books ❌")
		return error
	}
	book, exists := booksMap[bookId]
	if !exists {
		return fmt.Errorf("Book with ID %d does not exist", bookId)
	}

	book.IsBorrowed = false
	booksMap[bookId] = book
	//convaert map to slice
	books := make([]Book, 0)
	for _, book := range booksMap {
		books = append(books, book)
	}
	return SaveBooks(books)
}

func displayBooks(books map[int]Book) {
	fmt.Println("===== Books List =====")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s, Year: %d, Is Borrowed: %t\n", book.ID, book.Title, book.Author, book.Year, book.IsBorrowed)
	}
}

func searchBook(query string, books map[int]Book) {
	fmt.Println("===== Search Results =====")
	for _, book := range books {
		if strings.Contains(strings.ToLower(book.Title), strings.ToLower(query)) || strings.Contains(strings.ToLower(book.Author), strings.ToLower(query)) {
			fmt.Printf("ID: %d, Title: %s, Author: %s, Year: %d, Is Borrowed: %t\n", book.ID, book.Title, book.Author, book.Year, book.IsBorrowed)
		}
	}
}


func main() {
	fmt.Println("===== Books Management System =====")
	booksMap, err := LoadBooks()
	if err != nil {
		fmt.Println("Error loading books:", err)
		return
	}
	for {
		fmt.Println("1. Add a book")
		fmt.Println("2. Borrow a book")
		fmt.Println("3. Return for a book")
		fmt.Println("4. Search for a book")
		fmt.Println("5. Show all books")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			_, err := AddBooks()
			if err != nil {
				fmt.Println("Error adding book:", err)
			}
		case 2:
			var bookId int
			fmt.Print("Enter the book ID to borrow: ")
			fmt.Scanln(&bookId)
			if err := BorrowBook(bookId); err != nil {
				fmt.Println("Error borrowing book:", err)
			}
		case 3:
			var bookId int
			fmt.Print("Enter the book ID to return: ")
			fmt.Scanln(&bookId)
			if err := ReturnBook(bookId); err != nil {
				fmt.Println("Error returning book:", err)
			}
		case 4:
			var query string
			fmt.Print("Enter the search query: ")
			fmt.Scanln(&query)
			searchBook(query, booksMap)
		case 5:
			displayBooks(booksMap)
		case 6:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}	
	}
}
