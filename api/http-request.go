package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var baseUrl = "http://localhost:8080"

func ambilApi() ([]pelajar, error) {
	var err error
	var client = &http.Client{}
	var data []pelajar

	request, err := http.NewRequest("POST", baseUrl+"/data", nil)

	if err != nil {
		return nil, err
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func req() {
	var datapel, err = ambilApi()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}
	for _, each := range datapel {
		fmt.Println("ID : ", each.id, "NAMA : ", each.nama, "KELAS : ", each.kelas, "UMUR : ", each.umur)
	}
}
