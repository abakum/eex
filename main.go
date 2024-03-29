// go get internal/tool
package main

import (
	"io/fs"
	"log"

	"github.com/abakum/embed-encrypt/encryptedfs"
)

//go:generate go run internal/main.go

//go:generate go run github.com/abakum/embed-encrypt

//encrypted:embed hello.txt
var hello string

//encrypted:embed gopher.png
var gopher []byte

// encrypted:embed gopher.png
//
//encrypted:embed hello.txt "another.txt" "with spaces .txt"
var multiplefiles encryptedfs.FS

//encrypted:embed *.txt
var glob encryptedfs.FS

func main() {
	c, err := fs.ReadFile(glob, "hello.txt")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(hello, len(gopher), multiplefiles)

	log.Println(string(c))
}
