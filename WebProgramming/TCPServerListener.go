package main

import (
	"bufio"
	"net"
	"fmt"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	defer li.Close()
	if err != nil {
		panic(err)
	}

	for {
		conn, err := li.Accept()
		if err != nil {
			panic(err)
		}

		go handle(conn)

	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan(){
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %s\n", ln)
	}
	defer conn.Close()
	fmt.Println("Code got here.")

}
