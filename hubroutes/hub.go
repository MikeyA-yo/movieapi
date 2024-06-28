package hubroutes

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

var Num = rand.Intn(26)

// Function to identify a series based on it's name, returns an array of bytes, takes a url and name as parameters
func GetSeries(url string, name string) []byte {
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
func GetMovies(url string, name string) []byte {
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
func GetSearch(url string, name string) []byte {
	combineUrl := fmt.Sprintf("%v&s=%v", url, name)
	res, err := http.Get(combineUrl)
	if err != nil {
		fmt.Print("Error\n")
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return body
}

// Fuction to get random character
func RandomCharacter() string {
	aZ := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	return aZ[Num]
}

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
