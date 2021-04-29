package client

import (
	//"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type client struct {
	address string
	token string
}

func NewClient(address string, token string) *client {
	return &client{
		address: address,
		token: token,
	}
}

// struct to map responce data 

type ListUsers struct {
	PageCount     int    `json:"page_count"`
	PageNumber    int    `json:"page_number"`
	PageSize      int    `json:"page_size"`
	TotalRecords  int    `json:"total_records"`
	NextPageToken string `json:"next_page_token"`
	Users         []User `json:"users"`
}

type User struct {
	Id            string `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	Type          int    `json:"type"`
	PMI           int    `json:"pmi"`
	Verified      int    `json:"verified"`
	CreatedAt     string `json:"created_at"`
	LastLoginTime string `json:"last_login_time"`
	PicUrl        string `json:"pic_url"`
	Language      string `json:"language"`
	Status        string `json:"status"`
	RoleId        string `json:"role_id"`
}

// creating new request and autherizing  & returning responce

func handleReadRequest(url string, httpMethod string, JWT_Token string) (response []byte, err error) {
	httpClient := &http.Client{}
	
	var req *http.Request
	req, err = http.NewRequest(httpMethod, url, nil)
	if err != nil {
		return 
	}

	req.Header.Add("Authorization", "Bearer "+JWT_Token)
	req.Header.Add("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return
	}

	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return

}

// get list for all user

func ListUser(url, token string) {
	res, err := handleReadRequest(url, http.MethodGet, token)
	if err != nil {
		fmt.Println("Error\n")
	}

	var apiResponse ListUsers
	err = json.Unmarshal(res, &apiResponse)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("List of users:\n")
	for i:=0; i < len(apiResponse.Users); i++ {
		fmt.Print(apiResponse.Users[i],"\n")
	}
}



