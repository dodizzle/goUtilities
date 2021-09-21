package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type status struct {
	Status string `json:"status"`
}

var stat = []status{
	{Status: "okay"},
}

func main() {

	router := gin.Default()
	router.GET("/", getRoot)
	router.GET("/health", getHealth)
	router.GET("/date", getDate)
	router.Run("0.0.0.0:8080")
}

func getRoot(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "ok")
}
func getDate(c *gin.Context) {
	input := "2017-08-31"
	layout := "2006-01-02"
	t, _ := time.Parse(layout, input)
	mapD := map[string]string{"Date": t.String()}
	c.IndentedJSON(http.StatusOK, map[string]string(mapD))
}

func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, stat)
}
