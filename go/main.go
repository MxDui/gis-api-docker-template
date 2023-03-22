package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
)

func main() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./public", true)))
	r.Run(":8080")
}
