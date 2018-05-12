package main

import (
	"culfi/auth"
	"fmt"
)

func main() {
	// TODO start server
	// for now test code
	content, _ := auth.ReadAuth()
	fmt.Println(content)
}
