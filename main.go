package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

// FileSystem custom file system handler
type FileSystem struct {
	fs    http.FileSystem
	index string
}

// Open opens file
func (fs FileSystem) Open(path string) (http.File, error) {
	f, err := fs.fs.Open(path)
	if err != nil {
		// Rewrite all not-found requests to index (assume a virtual path)
		return fs.fs.Open(fmt.Sprintf("/%s", fs.index))
	}

	return f, nil
}

func main() {
	addr := flag.String("l", ":3000", "Specify a URI endpoint on which to listen")
	directory := flag.String("d", ".", "The root directory to host")
	index := flag.String("i", "index.html", "The index file")

	flag.Parse()

	fileServer := http.FileServer(FileSystem{
		fs:    http.Dir(*directory),
		index: *index,
	})

	http.Handle("/", fileServer)

	log.Printf("Serving %s on HTTP addr: %s\n", *directory, *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
