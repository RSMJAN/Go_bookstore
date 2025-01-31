package controllers

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "github.com/RSMJAN/go-bookstore/pkg/utils"
    "github.com/RSMJAN/go-bookstore/pkg/models"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
    newBooks := models.GetAllBooks()
    res, _ := json.Marshal(newBooks)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    bookId := vars["bookId"]
    ID, err := strconv.ParseInt(bookId, 0, 0)
    if err != nil {
        fmt.Println("error while parsing")
    }
    bookDetails, _ := models.GetBookById(ID)
    res, _ := json.Marshal(bookDetails)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
    createBook := &models.Book{}
    utils.ParseBody(r, createBook)
    b := createBook.CreateBook() // This line assumes CreateBook method exists in models.Book
    res, _ := json.Marshal(b)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    bookId := vars["bookId"]
    ID, err := strconv.ParseInt(bookId, 0, 0)
    if err != nil {
        fmt.Println("error while parsing")
    }
    book := models.DeleteBook(ID)
    res, _ := json.Marshal(book)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
    updateBook := &models.Book{}
    utils.ParseBody(r, updateBook)
    vars := mux.Vars(r)
    bookId := vars["bookId"]
    ID, err := strconv.ParseInt(bookId, 0, 0)
    if err != nil {
        fmt.Println("error while parsing")
    }
    bookDetails, db := models.GetBookById(ID)
    if updateBook.Name != "" {
        bookDetails.Name = updateBook.Name
    }
    if updateBook.Author != "" {
        bookDetails.Author = updateBook.Author
    }
    if updateBook.Publication != "" {
        bookDetails.Publication = updateBook.Publication
    }
    db.Save(&bookDetails)
    res, _ := json.Marshal(bookDetails)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}
