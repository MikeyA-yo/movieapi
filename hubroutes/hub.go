package hubroutes

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

var Num = rand.Intn(100)

// Function to identify a series based on it's name, returns an array of bytes, takes a url and name as parameters
func GetSeries(url, name string) []byte {
	combineUrl := fmt.Sprintf("%v&t=%v&type=series", url, name)
	res, err := http.Get(combineUrl)
	if err != nil {
		fmt.Print("Error\n")
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return body
}

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

// Function to get a search page pseudo randomly
func GetSearchRand(url, name string) []byte {

	combineUrl := fmt.Sprintf("%v&s=%v&page=%v", url, name, Num)
	res, err := http.Get(combineUrl)
	if err != nil {
		fmt.Print("Error\n")
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return body
}

// Fuction to get random character
// func RandomCharacter() string {
// 	aZ := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
// 	return aZ[Num]
// }

// function For a random word 178186
func GetWord() string {
	var WordNum = rand.Intn(178186)
	result := []string{}
	res, err := http.Get("https://random-word-api.herokuapp.com/all")
	if err != nil {
		fmt.Print("Error\n")
		return ""
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &result)
	return result[WordNum]
}

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

type searchResult struct {
	Title string `json:"Title"`
	Type  string `json:"Type"`
}

// actually get a title, and type to search based on GetWord() function
func GetTitle() []searchResult {
	var title serials
	randomWord := GetWord()
	URL := fmt.Sprintf("http://localhost:8080/search/%v", randomWord)
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

// function to find and keep genres
func GetDetailedRecommendation() {
	lists := GetTitle()
	length := len(lists)
	for i := 0; i < length; i++ {
		if lists[i].Type == "movie" {
			return
			//http.Get("http://localhost:8080/")
		}
	}
}
