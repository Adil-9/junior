package main

import (
	"encoding/json"
	"fmt"
	"io"
	"junior/api/requests"
	"junior/api/structures"
	"junior/internal"
	"net/http"
)

func main() {
	internal.Init() //creating logger

	mux := http.NewServeMux()
	mux.HandleFunc("/", HandleRequest)
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {

}

func testLinks() {
	link := requests.Agify + "Adil"
	println(link)

	req, err := http.NewRequest(http.MethodGet, link, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var answer structures.PersonAgify
	json.Unmarshal(body, &answer)
	fmt.Println(answer)
}
