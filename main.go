package main

import (
	"github.com/akshrana12/Go-Assignment/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	routes.UserRoute(router)
	router.Run()
}
