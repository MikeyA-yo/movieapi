package hubroutes

import (
	"fmt"
	"io"
	"net/http"
)

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

var Message = "Hello"
