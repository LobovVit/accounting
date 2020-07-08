package main

import (
	"../../accModule/app/domain"
	"fmt"
)

func main() {
	fmt.Printf(domain.Document(domain.NewDocument()))
}


