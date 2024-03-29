package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, er := net.Listen("tcp", ":8080")
	if er != nil {
		panic(er)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		io.WriteString(conn, "\nHello from TCP server\n")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "Well, I hope!")

		conn.Close()

	}
}
