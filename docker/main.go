package main

import (
	"fmt"
	"github.com/tarantool/go-tarantool"
	"log"
	"time"
)

func main() {
	z := time.Now()
	fmt.Printf("Вот  %v \n", z)

	conn, err := tarantool.Connect("localhost:3305", tarantool.Opts{})

	if err == nil {
		defer conn.Close()
		resp, _ := conn.Ping()
		fmt.Printf("resp=%v \n", resp)
		a := conn.Addr()
		fmt.Printf("resp=%v \n", a)
	} else {
		fmt.Printf("Connection refused %v \n", conn)
	}

	rr, err := conn.Select("customer", "customer_id", 0, 1000, tarantool.IterAll, []interface{}{uint(1)})
	log.Println("Select")
	log.Println("Error", err)
	log.Println("Code", rr.Code)
	log.Println("Data", rr.Data)
}
