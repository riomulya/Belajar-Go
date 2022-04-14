package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Nama  string `json:"Nama"`
	Kelas string `json:"Kelas"`
	Umur  int
}

func main() {
	jsonString := `{"Nama" : "Rio","Umur" : 18,"Kelas":"XII"}`
	Object := []user{{"Rio", "XII", 18}, {"Ilham", "XII", 18}}
	jsonData := []byte(jsonString)

	var data user
	err := json.Unmarshal(jsonData, &data)
	jsonObject, err2 := json.Marshal(Object)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Nama   \t:", data.Nama)
		fmt.Println("Kelas  \t:", data.Kelas)
		fmt.Println("Umur   \t:", data.Umur)
	}

	if err2 != nil {
		fmt.Println(err.Error())
	} else {
		jsonToString := string(jsonObject)
		fmt.Println(jsonToString)
	}

}
