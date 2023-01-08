package handler

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type BookList struct {
	Books []string
}

func fileRead(fileName string) []string {
	var bookList []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bookList = append(bookList, scanner.Text())
	}
	return bookList
}

func New(books []string) *BookList {
	return &BookList{Books: books}
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	bookList := fileRead("reading.txt")
	html, err := template.ParseFiles("view.html")
	if err != nil {
		log.Fatal(err)
	}
	getBooks := New(bookList)
	if err := html.Execute(w, getBooks); err != nil {
		log.Fatal(err)
	}
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	formValue := r.FormValue("value")
	voidJudge := strings.ReplaceAll(formValue, " ", "")
	if voidJudge != "" {
		file, err := os.OpenFile("reading.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.FileMode(0600))
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		_, err = fmt.Fprintln(file, formValue)
		if err != nil {
			log.Fatal(err)
		}
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	// deleteformValue := r.PostFormValue("delete")
	// formValue := r.FormValue("delete")
	// voidJudge := strings.ReplaceAll(formValue, " ", "")
	// if voidJudge != "" {
	// 	file, err := os.OpenFile("reading.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.FileMode(0600))
	// 	defer file.Close()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	_, err = fmt.Fprintln(file, formValue)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	http.Redirect(w, r, "/", http.StatusFound)
}
