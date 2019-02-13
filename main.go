package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
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
		// Rewrite all not-found requests to root (assume a virtual path)
		return fs.fs.Open(fmt.Sprintf("/%s", fs.root))
	}

	return f, nil
}

func main() {
	port := flag.String("l", "3000", "Specify a URI endpoint on which to listen")
	directory := flag.String("d", ".", "The directory to host")
	root := flag.String("r", "index.html", "The root (index) file")

	flag.Parse()

	fileServer := http.FileServer(FileSystem{
		fs:   http.Dir(*directory),
		root: *root,
	})

	http.Handle("/", fileServer)

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
