package hubroutes

import (
	"fmt"
	"io"
	"net/http"
)

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

var Message = "Hello"
