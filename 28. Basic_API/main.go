package main;

import (
	"strconv"
	"time"
	"math/rand"
	"encoding/json"
	"net/http"
	"log"
	"fmt"
	"github.com/gorilla/mux"
)

type Book struct{
	Issn string `json:"issn"`
	BookName string `json:"bookName"`
	Price float32 `json:"price"`
	Author *Author
}

type Author struct{
	Fullname string `json:"fullname"`
	Enail string `json:"email"`
}
func (b *Book) isEmpty() bool {
	return b.Issn == "" && b.BookName == ""
}

var bookDB []Book;


func main() {
	fmt.Println("Building a Basid Back-end API Server");


	var book1 Book = Book{
		Issn: "24535345",
		BookName: "Half Of a Yellow Sun",
		Price: 2345,
		Author: &Author{
			Fullname: "Chimamanda Adichie",
			Enail: "amanda@gmail.com",
		},
	}
	bookDB = append(bookDB, book1);

	r := mux.NewRouter(); // Creating an instance of the router from the "gorilla/mux" package


	r.HandleFunc("/", ServeHome).Methods("GET").Host("localhost") // serving the "ServerHome" function on the home route when hit with a "GET" request

	r.HandleFunc("/books", getAllBooks).Methods("GET");
	r.HandleFunc("/book", getSingleBook).Methods("GET");
	r.HandleFunc("/book", addSingleBook).Methods("POST");
	// r.HandleFunc("/book/{issn}", updateSingleBook).Methods("PATCH");
	s := r.PathPrefix("/book").Subrouter();
		s.HandleFunc("/{issn}", updateSingleBook).Methods("PATCH");

	r.PathPrefix("/book").HandlerFunc(ServeHome); // Program execution would definitely not get here

	log.Fatal(http.ListenAndServe(":3001", r)); // Spinning up the server and listing on the poert 3001. Also providing provision for loggin of errors if any is encounterd.
	// Always ensure that the line that listens and spins up the server is always at the bottom of the file, else the routes specified after it would not be reached for execution.
}

/* ======================= */
/* Handler Method Base URL */
/* ======================= */
func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<h1>Welcome To Server Home Directory</h1>`));
}


/* ===================================================== */
/* Handler Method To Get All Books In The Slice Of Books */
/* ===================================================== */
func getAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json");

	// Sending the Data as JSON using the JSON encoder to encode the response Writer
	json.NewEncoder(w).Encode(bookDB) // A one-liner equivalent for the below.

	// // sending the data as JSON using Marshal. This approach is good when it is used with MarshalIndent for viewing of the raw JSON.
	// dat, _ := json.Marshal(bookDB)
	// w.Write([]byte(dat));
}


/* ================================================================ */
/* Handler Method For Getting A Single Book From The Slice Of Books */
/* ================================================================ */
func getSingleBook(w http.ResponseWriter, r *http.Request) {
	bookId := mux.Vars(r);

	for _, book := range(bookDB) {
		if book.Issn == bookId["issn"] {
			w.Header().Set("Content-Type","application/json");
			json.NewEncoder(w).Encode(book)
			return;
		}
	}
	json.NewEncoder(w).Encode("Book not found in database");
}

/* ============================================================= */
/* Handler Method For Adding A Single Book To The Slice Of Books */
/* ============================================================= */
func addSingleBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json");

	// Checking if response body is empty and returning an error message
	if r.Body == nil {
		json.NewEncoder(w).Encode("Request body is empty");
		return;
	}

	// Checking if the request body is not empty, but does not contain the required properties
	var book Book;
	err := json.NewDecoder(r.Body).Decode(&book);
	HandleError(err);

	// Returning an error message and exiting the handler if the book does not contain the specified fields;
	if book.isEmpty() {
		json.NewEncoder(w).Encode("Request body does not contain the required properties");
		return
	}
	// Generating a Random number for the ISSN
	rand.NewSource(time.Now().UnixNano()); // Generating a Seed to be used for the generation of the random number.
	book.Issn = strconv.Itoa(rand.Intn(900) + 100) // Converting the random number generated to string to be used as the ISSN value.
	bookDB := append(bookDB, book); // Appending the book attded to the slice of books after is ISSN has been generated

	json.NewEncoder(w).Encode(bookDB); // Sending back the response, of all the books stored, with the new book added.
}


/* ===================================================== */
/* Handler Method To Update A Book In The Slice Of Books */
/* ===================================================== */
func updateSingleBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json");

	updateBookId := mux.Vars(r);
	fmt.Println(updateBookId["issn"]);
	for index, book := range(bookDB) {
		if book.Issn == updateBookId["issn"] {
			var updatedBook Book;
			err := json.NewDecoder(r.Body).Decode(&updatedBook);
			HandleError(err);
			updatedBook.Issn = updateBookId["issn"];
			bookDB = append(bookDB[:index], bookDB[(index + 1):]...);
			bookDB = append(bookDB, updatedBook);
			json.NewEncoder(w).Encode(bookDB);
			return
		}
	}

	json.NewEncoder(w).Encode("Book with ISSN not found in the books slice");
	return;
}

/* Utility Function For Error Handling */
func HandleError(err error) {
	if err != nil {
		panic(err);
	}
}