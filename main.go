package main

import (
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/MikeyA-yo/movieapi/docs"
	"github.com/MikeyA-yo/movieapi/hubroutes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type serials struct {
	Search []struct {
		Title  string `json:"Title"`
		Year   string `json:"Year"`
		ImdbID string `json:"imdbID"`
		Type   string `json:"Type"`
		Poster string `json:"Poster"`
	} `json:"Search"`
	TotalResults string `json:"totalResults"`
	Response     string `json:"Response"`
}

type serial struct {
	Title    string `json:"Title"`
	Year     string `json:"Year"`
	Rated    string `json:"Rated"`
	Released string `json:"Released"`
	Runtime  string `json:"Runtime"`
	Genre    string `json:"Genre"`
	Director string `json:"Director"`
	Writer   string `json:"Writer"`
	Actors   string `json:"Actors"`
	Plot     string `json:"Plot"`
	Language string `json:"Language"`
	Country  string `json:"Country"`
	Awards   string `json:"Awards"`
	Poster   string `json:"Poster"`
	Ratings  []struct {
		Source string `json:"Source"`
		Value  string `json:"Value"`
	} `json:"Ratings"`
	Metascore    string `json:"Metascore"`
	ImdbRating   string `json:"imdbRating"`
	ImdbVotes    string `json:"imdbVotes"`
	ImdbID       string `json:"imdbID"`
	Type         string `json:"Type"`
	TotalSeasons string `json:"totalSeasons"`
	Response     string `json:"Response"`
}

// @title MovieAPI
// @version 1.0
// @description A MovieAPI sever to serve random movies.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

func main() {
	//Add API key to URL from .env
	godotenv.Load(".env")
	//real URL
	Url := fmt.Sprintf("https://www.omdbapi.com/?apikey=%v", os.Getenv("API_KEY"))
	//Start the server other method's to come
	r := gin.Default()

	// Load HTML templates from the templates directory
	r.LoadHTMLGlob("templates/*")

	// Swagger setup
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler, url))

	r.GET("/", func(c *gin.Context) {
		// Render the HTML template
		c.HTML(200, "index.html", gin.H{
			"title": "Welcome Page",
		})
	})

	r.GET("/series/:name", func(c *gin.Context) {
		var data *serial
		name := c.Param("name")
		body := hubroutes.GetSeries(Url, name)
		json.Unmarshal(body, &data)
		c.AsciiJSON(200, data)
	})
	r.GET("/movies/:name", func(c *gin.Context) {
		var data *serial
		name := c.Param("name")
		body := hubroutes.GetMovies(Url, name)
		json.Unmarshal(body, &data)
		c.AsciiJSON(200, data)
	})
	r.GET("/search/:term", func(c *gin.Context) {
		var data *serials
		term := c.Param("term")
		rand := c.Query("rand")
		body := hubroutes.GetSearch(Url, term)
		if rand == "true" {
			body = hubroutes.GetSearchRand(Url, term)
		}
		json.Unmarshal(body, &data)
		c.AsciiJSON(200, data)
	})
	r.GET("/recommend", func(c *gin.Context) {
		//var data serials
		genre := c.Query("genre")
		if len(genre) > 2 {
			body := hubroutes.GetDetailedRecommendation(genre)
			c.AsciiJSON(200, body)
		}
	})
	r.Run()
}
