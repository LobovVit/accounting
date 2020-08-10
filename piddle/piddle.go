package old_main

import (
	"fmt"
	"github.com/tarantool/go-tarantool"
	_ "gopkg.in/goracle.v2"
)

func old_main() {
	conn, err := tarantool.Connect("127.0.0.1:3301", tarantool.Opts{})

	if err != nil {
		fmt.Println("Connection refused")
	}
	defer conn.Close()
	resp, err := conn.Ping()
	fmt.Printf("resp=%v", resp)
	a := conn.Addr()
	fmt.Printf("resp=%v \n", a)
	//
	//db, err := sql.Open("goracle", "user/pass@(DESCRIPTION=(ADDRESS_LIST=(ADDRESS=(PROTOCOL=tcp)(HOST=hostname)(PORT=port)))(CONNECT_DATA=(SERVICE_NAME=sn)))")
	//if err != nil {
	//	panic(err)
	//}
	//value, err := db.Query("select doc_type from xxt_bs_doc_reristry where rownum =1")
	//
	//var table_name string
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer value.Close()
	//for value.Next() {
	//	err := value.Scan(&table_name)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	log.Println(table_name)
	//}
	//err = value.Err()
	//if err != nil {
	//	log.Fatal(err)
	//}
}
