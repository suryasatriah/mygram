package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/suryasatriah/mygram/database"
	"github.com/suryasatriah/mygram/middleware"
	"github.com/suryasatriah/mygram/router"
)

const portnumber = 3000

func main() {

	database.ConnectDatabase()
	r := gin.Default()

	public := r.Group("/")
	router.PublicRoutes(public)

	private := r.Group("/")
	private.Use(middleware.Authentication())
	router.PrivateRoutes(private)

	r.Run(fmt.Sprintf("localhost:%d", portnumber))

}
