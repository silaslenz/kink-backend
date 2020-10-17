package main

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"kink_api/classifier"
	"kink_api/dataparsers"
	"time"
)

const testFilePath = "/home/silenz/downloads/kontohandelser2020-10-17.csv"

func main() {
	classifier.Train()
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/transactions", func(c *gin.Context) {
		c.JSON(200, dataparsers.ReadIcaBanken(testFilePath))
	})
	r.GET("/stats/:from/:to", func(c *gin.Context) {
		from, _ := time.Parse("2006-1-2", c.Param("from"))
		to, _ := time.Parse("2006-1-2", c.Param("to"))
		c.JSON(200, Sum(dataparsers.ReadIcaBanken(testFilePath), from, to))
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
