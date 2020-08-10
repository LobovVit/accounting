package main

import (
	"fmt"
	"github.com/tarantool/go-tarantool"
	"time"
)

func main() {
	z := time.Now()
	fmt.Printf("Вот  %v \n", z)

	conn, err := tarantool.Connect("tarantool1:3301", tarantool.Opts{})

	if err == nil {
		defer conn.Close()
		resp, _ := conn.Ping()
		fmt.Printf("resp=%v \n", resp)
		a := conn.Addr()
		fmt.Printf("resp=%v \n", a)
	} else {
		fmt.Printf("Connection refused %v \n", conn)
	}
}
