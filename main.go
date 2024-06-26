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
	//Add API key to URL from .env
	godotenv.Load(".env")
	api_url := fmt.Sprintf("https://www.omdbapi.com/?apikey=%v&t=frieren", os.Getenv("API_KEY"))

	//fetch response with http.Get
	res, err := http.Get(api_url)
	//check for errors
	if err != nil {
		fmt.Println("error")
		return
	}

	//create result variable which is of type *fiction struct
	var Result *fiction
	// close response at any time the function is returned
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	// idk what i'm doing
	// DecodeJson(res.Body, &Result)
	// fmt.Printf("%v \n", *Result)

	//Decode JSON from body into Result Variable
	json.Unmarshal(body, &Result)
	fmt.Printf("%v \n", *Result)
	//Just print the whole actual result
	fmt.Println(string(body))
	fmt.Println("Welcome!", hubroutes.Message)
	//Start the server other method's to come
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome")
	})
	r.Run()
}
