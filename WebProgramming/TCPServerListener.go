package main

import (
	"bufio"
	"fmt"
	"net"
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
		/*
			Below part of code is sending a content to Client un comment to send data to TCP client
		*/
		/*io.WriteString(conn, "\nHello How are you doing? This is TCP Server\n")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "Well, I hope!")
		conn.Close()*/
		// handle basically handle incoming data from client.
		go handle(conn)

	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %s\n", ln)
	}
	defer conn.Close()
	fmt.Println("Code got here.")

}
