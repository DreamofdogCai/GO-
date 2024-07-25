package main

import (
	"bubble_practice2/dao"
	"bubble_practice2/models"
	"bubble_practice2/routers"
	"fmt"
)

func main() {

	if err := dao.Connection(); err != nil { //database connection
		fmt.Println("The database connection failed")
		panic(err)
	}

	if err := models.CreateATable(); err != nil { //table creation
		fmt.Println("The table creation failed")
		panic(err)
	}

	defer dao.Close() //The connection Close

	r := routers.SetupRouters()

	r.Run("127.0.0.1:8080")
}
