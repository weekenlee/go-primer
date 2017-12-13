package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"

	chat "liweijian.com/chat"
)

// templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates",
			t.filename)))
	})
	t.templ.Execute(w, nil)
}

// newRoom makes a new Room.
func newRoom() *chat.Room {
	return &chat.Room{
		Forward: make(chan []byte),
		Join:    make(chan *chat.Client),
		Leave:   make(chan *chat.Client),
		Clients: make(map[*chat.Client]bool),
	}
}

func main() {
	r := newRoom()
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/Room", r)
	// get the Room going
	go r.Run()
	// start the web server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}