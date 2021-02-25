package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// User is dummy JSON structures
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Profile  struct {
		FullName string   `json:"full_name"`
		Birthday string   `json:"birthday"`
		Phones   []string `json:"phones"`
	} `json:"profile"`
	Articles []struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		PublishedAt string `json:"published_at"`
	} `json:"articles:"`
}

func parseJSON() {
	fmt.Println("Start Executing")
	content, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	var data []User
	err2 := json.Unmarshal(content, &data)
	if err2 != nil {
		fmt.Println("Failed to parse JSON data")
		fmt.Println(err2.Error())
	}
	fmt.Println("users:", data)
}

func main() {
	fmt.Println("Start Executing")
	parseJSON()
	fmt.Println("Finished Executing")
}
