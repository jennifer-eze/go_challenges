package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// fmt.Println()
	url := "https://open-weather13.p.rapidapi.com/city/landon/EN --header 'x-rapidapi-host: open-weather13.p.rapidapi.com"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	req.Header.Add("x-rapidapi-key", "YOU_API_KEY")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer res.Body.Close()
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Print(err.Error())
	}
	fmt.Println(string(body))
}
