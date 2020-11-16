package main

import (
	"fmt"
	"todo/routers"
)

func main() {
	r := routers.SetupRouter()
	if err := r.Run(":8090"); err != nil {
		fmt.Printf("Server lift failed with error: %v\n", err)
	}
}
