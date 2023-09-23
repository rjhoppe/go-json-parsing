package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

type Author struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Developer bool   `json:"is_developer"`
}

func main() {
	fmt.Println("Hello World")

	author := Author{Name: "Elliot Forbes", Age: 25, Developer: true}
	book := Book{Title: "Learning Concurrency in Python", Author: author}

	fmt.Printf("%+v\n", book)

	byte, err := json.MarshalIndent(book, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byte))

}
