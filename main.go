package main

import (
	"ginmvc/app/controllers"
)

func main() {
	router := controllers.GetRouter()
	router.Run(":8080")
}
