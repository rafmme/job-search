package main

import (
	"fmt"
	"os"
)

func init() {
	fmt.Println(os.Getenv("PORT"))
}

func main() {
	fmt.Println("Go forth!")
}
