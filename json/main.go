package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Datas struct {
	InventoryID int       `json:"inventory_id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Tags        []string  `json:"tags"`
	PurchasedAt int       `json:"purchased_at"`
	Placement   Placement `json:"placement"`
}

type Placement struct {
	RoomID int    `json:"room_id"`
	Name   string `json:"name"`
}

//var jsonFile *os.File

func main() {

	// Open jsonFile
	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	datas := []Datas{}

	err = json.Unmarshal(byteValue, &datas)
	if err != nil {
		fmt.Println("errorr : ", err.Error())
	}

	for i := 0; i < len(datas); i++ {
		fmt.Println("Inventory ID: ", datas[i].InventoryID)
		fmt.Println("Name: ", datas[i].Name)
	}

}
