package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"test.com/hub/hubroutes"
)

type fiction struct {
	Title string
	Genre string
}

func main() {
	godotenv.Load(".env")
	api_url := fmt.Sprintf("https://www.omdbapi.com/?apikey=%v&t=frieren", os.Getenv("API_KEY"))
	// req, _ := http.NewRequest("GET", api_url, nil)
	// req.Header.Add("accept", "application/json")

	res, err := http.Get(api_url)
	if err != nil {
		fmt.Println("error")
		return
	}
	var Result *fiction
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	// idk what i'm doing
	// DecodeJson(res.Body, &Result)

	// fmt.Printf("%v \n", *Result)
	json.Unmarshal(body, &Result)
	fmt.Printf("%v \n", *Result)
	fmt.Println(string(body))
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome")
	})
	fmt.Println("Welcome!", hubroutes.Message)
	// r.Run()
}
