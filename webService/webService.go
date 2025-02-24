package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type status struct {
	Status string `json:"status"`
}

var stat = []status{
	{Status: "okay"},
}

func init() {
	m, err := migrate.New(
		"file://migrations",
		"postgres://postgres:postgres@localhost:5432/webservice?sslmode=disable")
	CheckError(err)
	m.Up()

}

func main() {
	// connect to postgres database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"127.0.0.1", 5432, "postgres", "postgres", "webservice")
	db, err := sql.Open("postgres", psqlInfo)

	CheckError(err)
	err = db.Ping()
	CheckError(err)
	defer db.Close()

	router := gin.Default()
	router.GET("/", getRoot)
	router.GET("/health", getHealth)
	router.GET("/date", getDate)
	// add a route for the day of the week
	router.GET("/day", getDay)
	// add a route to receive a json value and return the value in the response
	router.POST("/json", getUpper)

	// add a route to write json to the database
	router.POST("/users", postUsers)
	router.Run("0.0.0.0:8080")
}

// write json to the database
func postUsers(c *gin.Context) {
	var json struct {
		user_id  int
		Username string `json:"username"`
		Email    string `json:"email"`
	}
	if c.BindJSON(&json) == nil {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			"127.0.0.1", 5432, "postgres", "postgres", "webservice")
		db, err := sql.Open("postgres", psqlInfo)
		CheckError(err)
		sqlStatement := `
	INSERT INTO users (username, email)
	VALUES ($1, $2)`
		user_id := 0
		_, err = db.Exec(sqlStatement, json.Username, json.Email)
		CheckError(err)
		fmt.Println(user_id)
	}
}

// convert value of json key to uppercase and return as json
func getUpper(c *gin.Context) {
	var json struct {
		Value string `json:"value"`
	}
	if c.BindJSON(&json) == nil {
		c.IndentedJSON(http.StatusOK, strings.ToUpper(json.Value))
	}
}

func getDay(c *gin.Context) {
	t := time.Now()
	c.IndentedJSON(http.StatusOK, t.Weekday())
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

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
