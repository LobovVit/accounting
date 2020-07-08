package main

import (
	"../app/domain"
	"fmt"
)

func main() {
	fmt.Printf(domain.Document(domain.NewDocument()))
}


