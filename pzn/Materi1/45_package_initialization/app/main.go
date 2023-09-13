package main

/**
Blank Identifier -> supaya dapat melakukan import package walaupun package tersebut tidak digunakan
karena biasanya package yang tidak digunakan akan otomatis hilang dari import
*/

import (
	"Materi1/45_package_initialization/database"
	_ "Materi1/45_package_initialization/tes" // blank identifier
	"fmt"
)

func main() {
	connection := database.GetDatabase()

	fmt.Println(connection)
}
