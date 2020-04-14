package main

import (
	"fmt"
)

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("works")
} 