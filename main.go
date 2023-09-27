package main

import (
	"bufio"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"raphaelluethy/go-go-br/services/database"
)

func main() {
	// Create an empty slice
	//var dynamicArray []int

	//// Add elements to the dynamic array
	//dynamicArray = append(dynamicArray, 10)
	//dynamicArray = append(dynamicArray, 20)
	//dynamicArray = append(dynamicArray, 30)
	//
	//// Access and modify elements
	//fmt.Println(dynamicArray[0]) // Output: 10
	//
	//dynamicArray[1] = 50
	//fmt.Println(dynamicArray) // Output: [10 50 30]
	//
	//// Length and capacity of the dynamic array
	//fmt.Println(len(dynamicArray)) // Output: 3 (number of elements)
	//fmt.Println(cap(dynamicArray)) // Output: 4 (capacity of the underlying array)
	//
	//// Iterate over the dynamic array
	//for _, element := range dynamicArray {
	//	fmt.Println(element)
	//}
	HTTPServer()
}

func DoDBThings() {
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your name:")
	input.Scan()
	name := input.Text()
	fmt.Println("Enter your email:")
	input.Scan()
	email := input.Text()
	fmt.Println("Enter your password:")
	input.Scan()
	password := input.Text()

	db := database.GetDatabase()
	res, err := db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", name, email, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("result is:", res)

	err = db.Close()
	if err != nil {
		panic(err)
	}
}

func HTTPServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		if err != nil {
			return
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
