package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("MemoryGuard Collector iniciado")

	for {
		fmt.Println("Collector ativo:", time.Now())
		time.Sleep(10 * time.Second)
	}
}