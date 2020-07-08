package main

import (
	"fmt"
	"github.com/LobovVit/accounting/docModule/app/domain"
)

func main() {
	fmt.Printf(domain.Document(domain.NewDocument("qqq")))
}
