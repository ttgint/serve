package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// FileSystem custom file system handler
type FileSystem struct {
	fs   http.FileSystem
	root string
}

// Open opens file
func (fs FileSystem) Open(path string) (http.File, error) {
	f, err := fs.fs.Open(path)
	if err != nil {
		f, _ = fs.fs.Open(fmt.Sprintf("/%s", fs.root))
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := fmt.Sprintf("%s/%s", strings.TrimSuffix(path, "/"), fs.root)
		if _, err := fs.fs.Open(index); err != nil {
			return nil, err
		}
	}

	return f, nil
}

func main() {
	port := flag.String("p", "3000", "port to serve on")
	directory := flag.String("d", ".", "the directory to host")
	root := flag.String("r", "index.html", "the root (index) file")
	flag.Parse()

	fileServer := http.FileServer(FileSystem{
		fs:   http.Dir(*directory),
		root: *root,
	})

	r := mux.NewRouter()
	r.PathPrefix("/").Handler(fileServer)
	http.Handle("/", r)
	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
