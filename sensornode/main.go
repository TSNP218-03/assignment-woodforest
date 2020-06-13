package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

func main() {
	c, err := net.Dial("tcp", "localhost:55555")
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Println("connected.")

	for {
		// text := strconv.Itoa(rand.Intn(100))
		// c.Write([]byte(text))
		// log.Println(text)

		floatnum := rand.Float64() 	num)))
		log.Println(floatnum)
		time.Sleep(2 * time.Second) // Timeout for 2 second
	}
}
