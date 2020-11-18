package main

import (
	"fmt"
	"todo/db"
	"todo/routers"
)

func main() {

	db.InitDB()
	r := routers.SetupRouter()
	if err := r.Run(":8090"); err != nil {
		fmt.Printf("Server lift failed with error: %v\n", err)
	}
}
