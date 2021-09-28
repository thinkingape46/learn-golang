package main

import (
	"fmt"
	"path"
)


func main() {
	var dir, file string

	dir, file = path.Split("src/index.html")
	fmt.Println("line 13:", dir, file)

	// if only one variable is needed, use _ to discard other values
	_, fileTwo := path.Split("styles/main.scss")
	fmt.Println("filename: ", fileTwo)
}
