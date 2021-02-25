package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// User is dummy JSON structures
type User struct {
	InventoryID int      `json:"inventory_id"`
	Name        string   `json:"name"`
	Type        string   `json:"type"`
	Tags        []string `json:"tags"`
	PurchasedAt int      `json:"purchased_at"`
	Placement   struct {
		RoomID int    `json:"room_id"`
		Name   string `json:"name"`
	} `json:"placement"`
}

// func parseJSON() {
// 	content, err := ioutil.ReadFile("data.json")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	var data []User
// 	err2 := json.Unmarshal(content, &data)
// 	if err2 != nil {
// 		fmt.Println("Failed to parse JSON data")
// 		fmt.Println(err2.Error())
// 	}
// 	fmt.Println("users:", data)
// 	return
// }

//Items in Meeting Room
// func (user User) getItems() []string {
// 	var mySlice [5]string
// 	for

// }

// All Electronic Devices

// All Furniture

// Item purchased on 16 Januari 2020

// All items with brown color

func main() {
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
	fmt.Println("Finished Executing")
}
