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
	num := rand.Intn(int(maxPageNumber))
	combineUrl := fmt.Sprintf("%v&s=%v&page=%v", url, name, num)
	res, err := http.Get(combineUrl)
	if err != nil {
		fmt.Print("Error\n")
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return body
}

// Fuction to get random character
//
//	func RandomCharacter() string {
//		aZ := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
//		return aZ[Num]
//	}

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

var WordsArray = Words()

// function For a random word 178186
func GetWord() string {
	var WordNum = rand.Intn(178186)
	result := WordsArray
	return result[WordNum]
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

type genreLike struct {
	Title string `json:"Title"`
	Genre string `json:"Genre"`
}

// function to find and keep genres
func GetDetailedRecommendation(genre string) []byte {
	URL := "http://localhost:8080/"
	lists := GetTitle()
	length := len(lists)
	var data []byte
	for i := 0; i < length; i++ {
		name := fmt.Sprintf("%v", lists[i].Title)
		var genreData genreLike
		if lists[i].Type == "movie" {
			urlContent := fmt.Sprintf("%v/movie/%v", URL, name)
			res, e := http.Get(urlContent)
			if e != nil {
				fmt.Println(e)
			}
			defer res.Body.Close()
			body, _ := io.ReadAll(res.Body)
			data = body
			json.Unmarshal(body, &genreData)
			genreSlice := strings.Split(genreData.Genre, ", ")
			if slices.Contains(genreSlice, genre) {
				return body
			}
		} else {
			urlContent := fmt.Sprintf("%v/series/%v", URL, name)
			res, e := http.Get(urlContent)
			if e != nil {
				fmt.Println(e)
			}
			defer res.Body.Close()
			body, _ := io.ReadAll(res.Body)
			data = body
			json.Unmarshal(body, &genreData)
			genreSlice := strings.Split(genreData.Genre, ", ")
			if slices.Contains(genreSlice, genre) {
				return body
			}
		}
	}
	return data
}
