package main

import (
	"fmt"
	"junior/internal/db"
	"junior/internal/handlers"
	"junior/internal/logger"
	"net/http"
	"strconv"
)

func main() {
	// test()
	logger.Init() //creating logger
	db := db.InitDB()
	handler := handlers.CreateHandler(db)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.HandleRequest)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		logger.DebugLog.Fatal(err)
	}
}

func test() {
	_, err := strconv.Atoi("")
	fmt.Println(err)
}

// func testLinks() {
// 	link := requests.Agify + "Adil"
// 	println(link)

// 	req, err := http.NewRequest(http.MethodGet, link, nil)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	var answer structures.PersonAgify
// 	json.Unmarshal(body, &answer)
// 	fmt.Println(answer)
// }
