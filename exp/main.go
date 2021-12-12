package main

import (
	"html/template"
	"os"
)

type User struct {
	Name  string
	Dog   string
	Int   int
	Float float32
	Slice []string
	Map   map[string]string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := User{
		Name:  "John Smith",
		Dog:   "Morty",
		Int:   176,
		Float: 3.928,
		Slice: []string{"1", "Two", "3"},
		Map: map[string]string{
			"key1": "value1",
			"key2": "value2",
		},
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}
