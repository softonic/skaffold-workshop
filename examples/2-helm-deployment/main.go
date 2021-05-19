package main

import (
	"fmt"
	"time"
)

func main() {
	for counter := 0; ; counter++ {
		fmt.Println("Hello world! by pod!", counter)

		time.Sleep(time.Second * 1)
	}
}
