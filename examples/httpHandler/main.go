package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func infoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Info\n")
}

// func mapHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println(pathsToUrls)
// }

func main() {
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	router := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	// Per route middleware, you can add as many as you desire.
	router.GET("/benchmark", MyBenchLogger(), benchEndpoint)

}
