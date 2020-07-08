package main

import (
	"fmt"
	"github.com/LobovVit/accounting/accModule/app/domain"
)

func main() {
	fmt.Printf(domain.Document(domain.NewDocument()))
}


