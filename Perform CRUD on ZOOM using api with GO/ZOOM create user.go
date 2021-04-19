package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Userinfo struct {
	Email     string `json:"email"`
	Type      int    `json:"type"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type CreateUserRequest struct {
	Action   string   `json:"action"`
	UserInfo Userinfo `json:"user_info"`
}

type CreateUserResponse struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Type      int    `json:"type"`
}

func handleRequest(url string, httpMethod string, body []byte) (response []byte, err error) {
	httpClient := &http.Client{}

	var req *http.Request
	req, err = http.NewRequest(httpMethod, url, bytes.NewBuffer(body))
	if err != nil {
		return
	}

	authToken := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MTkzNTg3MzcsImlhdCI6MTYxODc1MzkzN30.0w1BOrihMQ2DIay703RY7ZfihqHXfMQqBA1m3vDSzl0"

	req.Header.Add("Authorization", "Bearer "+authToken)
	req.Header.Add("Content-Type", "application/json")

	var resp *http.Response
	resp, err = httpClient.Do(req)
	if err != nil {
		return
	}

	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}

func createUser() (createUserResponse CreateUserResponse, err error) {
	createUserRequest := CreateUserRequest{
		Action: "create",
		UserInfo: Userinfo{
			Email:     "ravikishandaiaya@gmail.com",
			Type:      1,
			FirstName: "Ravi",
			LastName:  "Daiya",
		},
	}

	var reqBody []byte
	reqBody, err = json.Marshal(createUserRequest)
	if err != nil {
		return
	}

	base_url := "https://api.zoom.us/v2/users/"
	httpMethod := http.MethodPost

	var b []byte
	b, err = handleRequest(base_url, httpMethod, reqBody)

	if err != nil {
		return
	}

	err = json.Unmarshal(b, &createUserResponse)
	if err != nil {
		return
	}

	return
}

func main() {
	response, err := createUser()
	if err != nil {
		fmt.Println("error")
	}

	fmt.Println(response)
}
