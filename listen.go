package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":53")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func(conn net.Conn) {
			name, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			name = name[:len(name)-1]
			fmt.Println(name)
			ips, err := resolve(name)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(ips)
			conn.Write([]byte(ips[0] + "\n"))
			conn.Close()
		}(conn)
	}
}
