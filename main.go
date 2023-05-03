package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	id = flag.String("id", "00", "starting device id")
	n  = flag.Int("number", 10, "number of query to generate")
)

func generateInsertQuery(id string, n int) string {

	// 0x206A94E74320
	hexStr := "0x"
	hexStr += id

	num, err := strconv.ParseInt(hexStr, 0, 64)
	if err != nil {
		panic(err)
	}

	var query string
	for i := 0; i < n; i++ {
		id := num + int64(i)

		query += fmt.Sprintf("('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s'), ", strings.ToUpper(strconv.FormatInt(id, 16)), "8fc766ce-7320-44b6-8b19-84f178dcb5bc", "3.2.4.0-0-nogit-dev", "ssl:manager-kenshi.tpd.hitroncloud.cn:443", "OS2210", "ssl:redirector-kenshi.tpd.hitroncloud.cn:443", "2023-03-29 05:05:28.897007+00", "2023-03-29 05:05:28.897007+00")
	}
	return fmt.Sprintf("INSERT INTO awlan_nodes (id, _uuid, firmware_version, manager_addr, model, redirector_addr, created_at, updated_at) VALUES %s;", query[:len(query)-2])
}

func main() {
	flag.Parse()

	// create file
	file, err := os.Create("query.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	query := generateInsertQuery(*id, *n)
	// fmt.Println(query)

	_, err = file.WriteString(query)
	if err != nil {
		log.Fatal(err)
	}
}
