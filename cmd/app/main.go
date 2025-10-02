package main

import (
	"fmt"

	internal "github.com/RifkyA911/management-inventory/internal/database"
)

func main() {
	internal.InitMongo()

	// contoh ambil collection
	usersCol := internal.MongoDB.Collection("barang")

	fmt.Println("Collection:", usersCol.Name())
}
