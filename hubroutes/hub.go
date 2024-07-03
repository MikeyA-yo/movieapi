package hubroutes

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
	"slices"
	"strconv"
	"strings"
)

var Num = rand.Intn(100)

/*
 Implementation Details:
  takes to arguments:
  1.) the global omdb API Url with your api key
  2.) Name of search query
  ensures to return only type series (omdbapi.com)
  combines them into a url and fetch with http
  then return the response body
*/

// Function to identify a series based on it's name, returns an array of bytes, takes a url and name as parameters
func GetSeries(url, name string) []byte {
	var errR []byte
	combineUrl := fmt.Sprintf("%v&t=%v&type=series", url, name)
	res, err := http.Get(combineUrl)
	if err != nil {
		fmt.Print("Error\n")
		fmt.Println(err.Error())
		return errR
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return body
}

/*
 Implementation Details:
  takes to arguments:
  1.) the global omdb API Url with your api key
  2.) Name of search query
  ensures to return only type movie (omdbapi.com)
  combines them into a url and fetch with http
  then return the response body
*/

// Function to identify a movie based on it's name, returns an array of bytes, takes a url and name as parameters
func GetMovies(url, name string) []byte {
	combineUrl := fmt.Sprintf("%v&t=%v&type=movie", url, name)
	res, err := http.Get(combineUrl)
	if err != nil {
		fmt.Print("Error\n")
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return body
}

/*
 Implementation Details:
  takes to arguments:
  1.) the global omdb API Url with your api key
  2.) Name of search query
  combines them into a url and fetch with http
  then return the response body
*/

// function to Search for movies and series generally
func GetSearch(url, name string) []byte {
	combineUrl := fmt.Sprintf("%v&s=%v", url, name)
	res, err := http.Get(combineUrl)
	if err != nil {
		fmt.Print("Error\n")
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return body
}

/*
 Implementation Details:
 Create a struct type based of the json response for search
 check out omdbapi.com for search results
*/

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

/*
 Implementation Details:
 use the GetSearch funtion to get the response body
 then store the count of possible page numbers from TotalResults property (read omdbapi.com docs)
 generate a random page number based  on the max possible page
 use that random number to fetch a random page and return the response body
*/

// Function to get a search page pseudo randomly
func GetSearchRand(url, name string) []byte {
	var dataJson serials
	randData := GetSearch(url, name)
	json.Unmarshal(randData, &dataJson)
	count, err := strconv.Atoi(dataJson.TotalResults)
	if err != nil {
		fmt.Println("Error")
		return randData
	}
	maxPageNumber := math.Round(float64(count / 10))
	num := rand.Intn(int(maxPageNumber)) + 1
	if num > 100 {
		num = 100
	}
	combineUrl := fmt.Sprintf("%v&s=%v&page=%v", url, name, num)
	res, err := http.Get(combineUrl)
	if err != nil {
		fmt.Print("Error\n")
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return body
}

/*
 Implementation Details:
 Get large list's of word's basically from an api and return it as a slice of strings
*/

// function to simplify process
func Words() []string {
	result := []string{}
	res, err := http.Get("https://random-word-api.herokuapp.com/all")
	if err != nil {
		fmt.Print("Error\n")
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &result)
	return result
}

// store all the words in this variable
var WordsArray = Words()

/*
 Implementation Details:
  use a random number to get an indiviual word from the slice
*/
//https://wordgenerator-api.herokuapp.com/api/v1/resources/words?lang=EN&amount=5
// function For a random word 178186
func GetWord() string {
	var WordNum = rand.Intn(len(WordsArray))
	result := WordsArray
	return result[WordNum]
}

/*
 Implementation Details:
   Below i'd be working on genre not much to be said on implementation details yet
*/

// Type for working with genre
type searchResult struct {
	Title string `json:"Title"`
	Type  string `json:"Type"`
}

// actually get a title, and type to search based on GetWord() function
func GetTitle() []searchResult {
	var title serials
	randomWord := GetWord()
	URL := fmt.Sprintf("https://movieapi-gcve.onrender.com/search/%v", randomWord)
	res, err := http.Get(URL)
	if err != nil {
		fmt.Print("Error\n")
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &title)
	if title.Response == "False" {
		return GetTitle()
	}
	arrLength := len(title.Search)
	titlesRes := []searchResult{}
	//titles := []string{}
	for i := 0; i < arrLength; i++ {
		results := searchResult{
			Title: title.Search[i].Title,
			Type:  title.Search[i].Type,
		}
		titlesRes = append(titlesRes, results)
		//titles = append(titles, title.Search[i].Title)
	}
	return titlesRes
}

type genreLike struct {
	Title string `json:"Title"`
	Genre string `json:"Genre"`
}

// huge slice of movies and genre for fallback cases (todo)
var breakingBad genreLike = genreLike{
	Title: "Breaking Bad",
	Genre: "Crime, Drama, Thriller",
}

var gameOfThrones genreLike = genreLike{
	Title: "Game of Thrones",
	Genre: "Action, Adventure, Drama",
}
var chernobyl genreLike = genreLike{
	Title: "Chernobyl",
	Genre: "Drama, History, Thriller",
}

var frieren genreLike = genreLike{
	Title: "Frieren: Beyond Journey's End",
	Genre: "Animation, Adventure, Drama",
}

var gintama genreLike = genreLike{
	Title: "Gintama",
	Genre: "Animation, Action, Comedy",
}

var attackOnTitan genreLike = genreLike{
	Title: "Attack on Titan",
	Genre: "Animation, Action, Adventure",
}

var hxh genreLike = genreLike{
	Title: "Hunter x Hunter",
	Genre: "Animation, Action, Adventure",
}

var demonSlayer genreLike = genreLike{
	Title: "Demon Slayer: Kimetsu no Yaiba",
	Genre: "Animation, Action, Adventure",
}

var rickNMorty genreLike = genreLike{
	Title: "Rick and Morty",
	Genre: "Animation, Adventure, Comedy",
}

var theOffice genreLike = genreLike{
	Title: "The Office",
	Genre: "Comedy",
}

var theWalkingDead genreLike = genreLike{
	Title: "The Walking Dead",
	Genre: "Drama, Horror, Thriller",
}

var theLastOfUs genreLike = genreLike{
	Title: "The Last of Us",
	Genre: "Action, Adventure, Drama",
}

var recommendations = []genreLike{breakingBad, gameOfThrones, gintama, frieren, chernobyl, attackOnTitan, hxh, theLastOfUs, theOffice, theWalkingDead, demonSlayer, rickNMorty}

func GetRandomRec() genreLike {
	n := rand.Intn(len(recommendations))
	return recommendations[n]
}

// function to check if a slice is in a slice
func ContainsSlice(a, b []string) bool {
	for _, v := range b {
		if !slices.Contains(a, v) {
			return false
		}
	}
	return true
}

type serialP struct {
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

// function to find and keep genres
func GetDetailedRecommendation(genre string) serialP {
	var jsonData serialP
	URL := "https://movieapi-gcve.onrender.com/"
	lists := GetTitle()
	length := len(lists)
	genreLower := strings.ToLower(genre)
	for i := 0; i < length; i++ {
		name := fmt.Sprintf("%v", lists[i].Title)
		fmt.Println(lists)
		var genreData genreLike
		if lists[i].Type == "movie" {
			urlContent := fmt.Sprintf("%vmovies/%v", URL, name)
			res, e := http.Get(urlContent)
			if e != nil {
				fmt.Println(e)
			}
			defer res.Body.Close()
			body, _ := io.ReadAll(res.Body)
			json.Unmarshal(body, &genreData)

			genreSlice := strings.Split(strings.ToLower(genreData.Genre), ", ")
			if strings.Contains(genreLower, ",") {

				genLowSlice := strings.Split(genreLower, ",")
				if ContainsSlice(genreSlice, genLowSlice) {
					json.Unmarshal(body, &jsonData)
					break
				} else {
					res, err := http.Get("https://movieapi-gcve.onrender.com/series/frieren")
					if err != nil {
						fmt.Println(err)
					}
					defer res.Body.Close()
					body, _ := io.ReadAll(res.Body)
					json.Unmarshal(body, &jsonData)
				}
			} else if slices.Contains(genreSlice, genreLower) {
				json.Unmarshal(body, &jsonData)
				break
			} else {
				res, err := http.Get("https://movieapi-gcve.onrender.com/series/frieren")
				if err != nil {
					fmt.Println(err)
				}
				defer res.Body.Close()
				body, _ := io.ReadAll(res.Body)
				json.Unmarshal(body, &jsonData)
			}
		} else {
			urlContent := fmt.Sprintf("%vseries/%v", URL, name)
			res, e := http.Get(urlContent)
			if e != nil {
				fmt.Println(e)
			}
			defer res.Body.Close()
			body, _ := io.ReadAll(res.Body)
			json.Unmarshal(body, &genreData)
			genreSlice := strings.Split(genreData.Genre, ", ")
			if strings.Contains(genreLower, ",") {
				genLowSlice := strings.Split(genreLower, ",")
				if ContainsSlice(genreSlice, genLowSlice) {
					json.Unmarshal(body, &jsonData)
					break
				} else {
					res, err := http.Get("https://movieapi-gcve.onrender.com/series/frieren")
					if err != nil {
						fmt.Println(err)
					}
					defer res.Body.Close()
					body, _ := io.ReadAll(res.Body)
					json.Unmarshal(body, &jsonData)
				}
			} else if slices.Contains(genreSlice, genreLower) {
				json.Unmarshal(body, &jsonData)
				break
			} else {
				res, err := http.Get("https://movieapi-gcve.onrender.com/series/frieren")
				if err != nil {
					fmt.Println(err)
				}
				defer res.Body.Close()
				body, _ := io.ReadAll(res.Body)
				json.Unmarshal(body, &jsonData)
			}
		}
	}
	return jsonData
}

// Fuction to get random character
//
//	func RandomCharacter() string {
//		aZ := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
//		return aZ[Num]
//	}
