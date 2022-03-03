package main

import (
	"fmt"
	"homework-2-hilalbalci/myfunctions"
	"os"
	"strconv"
	"strings"
)

var myBookList []string

type author struct {
	name string
}
type Deletable interface {
	deleteBook() bool
}

type book struct {
	id         int
	pageNumber int
	stockCount int
	bookName   string
	stockCode  string
	isbnNumber string
	author
	price     float64
	isDeleted bool
	Deletable bool
}

var myBookStructList []*book

func init() {
	myBookList = []string{"In Search of Lost Time",
		"Ulysses",
		"Ul",
		"Don Quixote",
		"Don Quixote Second",
		"The Great Gatsby",
		"Moby Dick",
		"Hamlet",
		"The Odyssey",
		"Madame Bovary"}
	//for index, v := range myBookList {
	//	mybook := newBook(v, index)
	//	myBookStructList = append(myBookStructList, mybook)
	//}

}
func main() {
	myBookStructList = []*book{
		{
			id:         1,
			pageNumber: 30,
			stockCount: 30,
			bookName:   "hello",
			stockCode:  "dsjadas",
			isbnNumber: "dshajkdsa",
			author:     author{name: "hello author"},
			price:      26.04,
			isDeleted:  false,
			Deletable:  true},
	}
	if (len(os.Args[1:]) == 0) || myfunctions.CheckArgs(strings.ToLower(os.Args[1])) {
		fmt.Print("Commands that you can use in this program : \nsearch \nlist \nbuy \ndelete\n")
	} else {
		operations(os.Args[2:], strings.ToLower(os.Args[1]), myBookStructList)
	}

}

func operations(keyword []string, key string, list []*book) {
	if key == "search" {
		fmt.Print(contains(list, keyword[0]))
	} else if key == "list" {
		for _, v := range list {
			fmt.Printf("Id: %v \nBook Name: %s \nAuthor Name: %s \nPage Number: %v \nStock Count: %v \nStock Code: %s \nISBN Number: %s \nPrice: %v \nIs Deleted: %t\n\n", v.id, v.bookName, v.author.name, v.pageNumber, v.stockCount, v.stockCode, v.isbnNumber, v.price, v.isDeleted)
		}
	} else if key == "buy" {
		if len(keyword) != 2 {
			fmt.Println("You must type the id and the copies of the book you want to buy. E.g : buy 1 3")
			return
		}
		firstArg, _ := strconv.Atoi(keyword[0])
		book, err := detectBookFromId(firstArg)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			secondArg, error := strconv.Atoi(keyword[1])
			if error != nil {
				fmt.Println("You must type the id and the copies of the book you want to buy. E.g : buy 1 3")
				return
			}
			err := buyBook(book, secondArg)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("Congrats! You've bought %v books. Current stock for this book is now %v", firstArg, book.stockCount)
		}

	} else if key == "delete" {
		firstArg, _ := strconv.Atoi(keyword[0])
		book, err := detectBookFromId(firstArg)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			err := deleteBook(book)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("Congrats! You've deleted the book '%s' from the list. IsDeleted and Deletable fields for this book are now %t and %t", book.bookName, book.isDeleted, book.Deletable)
		}
	}
}

func contains(s []*book, str string) []book {
	result := []book{}
	for _, v := range s {
		if strings.Contains(strings.ToLower(v.author.name), strings.ToLower(str)) || strings.Contains(strings.ToLower(v.bookName), strings.ToLower(str)) || strings.Contains(strings.ToLower(v.isbnNumber), strings.ToLower(str)) {
			result = append(result, *v)
		}
	}
	return result
}

func newBook(bookName string, id int) *book {
	p := new(book)
	p.id = id
	p.pageNumber = myfunctions.RandomNumberGenerator()
	p.stockCount = myfunctions.RandomNumberGenerator()
	p.bookName = bookName
	p.stockCode = myfunctions.GenerateRandomString(20)
	p.isbnNumber = myfunctions.GenerateRandomString(20)
	p.author = author{"Author Name"}
	p.price = float64(myfunctions.RandomNumberGenerator())
	p.isDeleted = false
	return p
}

func detectBookFromId(bId int) (*book, error) {
	for _, v := range myBookStructList {
		if bId == v.id {
			return v, nil
		}
	}
	fmtErr := fmt.Errorf("Seems like this book doesn't exist.. ")
	return nil, fmtErr
}

func buyBook(b *book, count int) error {
	if b.stockCount >= count {
		b.stockCount -= count
		return nil
	} else {
		fmtErr := fmt.Errorf("There's no enough stock for this book.. ")
		return fmtErr
	}

}

func deleteBook(b *book) error {
	if b.Deletable && (b.isDeleted == false) {
		b.isDeleted = true
		b.isDeleted = false
		return nil
	} else {
		fmtErr := fmt.Errorf("Either this book is not deletable or is already deleted.. ")
		return fmtErr
	}
}
