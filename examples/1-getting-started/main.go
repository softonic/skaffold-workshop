package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Hello skaffold! Manual")

		time.Sleep(time.Second * 1)
	}
}
