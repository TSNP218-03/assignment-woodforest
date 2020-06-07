package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	ln, err := net.Listen("tcp", "localhost:55555")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer ln.Close()

	c, err := ln.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Println("connected.")
	buf := make([]byte, 1024) // create a buffer (buf) of size 1024 bytes

	for {
		n, _ := c.Read(buf)
		buf = buf[:n]    // trim byte slices
		s := string(buf) // convert a slice of bytes to a string
		log.Printf("Received %s", s)

		// i, _ := strconv.Atoi(s)       // string parse to Int
		// log.Printf("Add 1 = %d", i+1) // perform addition (+1)
	}
}
