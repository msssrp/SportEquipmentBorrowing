package main

import (
	"fmt"
	"github.com/msssrp/SportEquipmentBorrowing/pkg/database"
)

func main() {
	_, err := database.ConnectMongo()
	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Printf("connected")
}
