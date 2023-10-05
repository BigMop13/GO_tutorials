package main

import (
	"fmt"
	"io"
	"os"
)

// 2 solutions

// func main() {
// 	filePath := os.Args[1]

// 	file, err := os.ReadFile(filePath)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(string(file))
// }

func main() {
	f, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error: ", err)
	}

	io.Copy(os.Stdout, f)
}
