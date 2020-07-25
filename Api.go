package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type responseStruct struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles     []struct {
		Title string `json:"title"`
	} `json:"articles"`
}

type recieve struct {
	answer string
}

func main() {
	// Get it

	var f string

	apiKey := os.Getenv("API_KEY")

	httpResponse, error := http.Get("https://newsapi.org/v2/top-headlines?country=us&pageSize=10&page=1" + "&apiKey=" + apiKey)
	if error != nil {
		log.Fatalln(error)
	}

	responseData, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer httpResponse.Body.Close()

	var responseObject responseStruct
	json.Unmarshal(responseData, &responseObject)
	fmt.Println(responseObject.Articles)

	a := &recieve{}

	out, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(out))

	f = string(out)
	fmt.Print(f)

}
