package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Book struct {
	Id      int        `json:"id"`
	Title   string     `json:"title"`
	Author  string     `json:"author"`
	Pages   int        `json:"pages"`
	IsRead  bool       `json:"is_read"`
	AddedAt time.Time  `json:"added_at"`
	ReadAt  *time.Time `json:"read_at"`
}

var (
	library = make(map[int]Book)
	mu      sync.RWMutex
	nextID  = 1
)

func AddBookHandler(w http.ResponseWriter, r *http.Request) {
	var b Book

	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	b.Id = nextID
	b.AddedAt = time.Now()
	library[b.Id] = b
	nextID++

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(b)
}

func ReadBookHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	if idStr == "" {
		http.Error(w, "id parameter is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	b, ok := library[id]

	if !ok {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	b.IsRead = true
	now := time.Now()
	b.ReadAt = &now
	library[id] = b

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(b)
}

func GetBookInfo(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	if idStr == "" {
		http.Error(w, "id parameter is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	mu.RLock()
	defer mu.RUnlock()

	b, ok := library[id]

	if !ok {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(b)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	defer mu.RUnlock()

	books := make([]Book, 0, len(library))

	for _, book := range library {
		books = append(books, book)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	if idStr == "" {
		http.Error(w, "id parameter is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if _, exists := library[id]; !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	delete(library, id)
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	http.HandleFunc("/api/v1/book", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetBookInfo(w, r)
		case http.MethodDelete:
			DeleteBook(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/v1/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetAllBooks(w, r)
		case http.MethodPost:
			AddBookHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/v1/book/read", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			ReadBookHandler(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
