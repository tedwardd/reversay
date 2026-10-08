package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Fprintln(conn, reverse(line))
	}
}

func main() {
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	fmt.Printf("Listening on %s\n", ln.Addr().String())

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("accept error:", err)
			continue
		}
		go handleConn(conn)
	}
}
