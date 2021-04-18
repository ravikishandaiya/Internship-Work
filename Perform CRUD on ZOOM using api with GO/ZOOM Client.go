package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const JWT_Token string = "JWT Token"
const base_URL string = "https://api.zoom.us/v2/users/"

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

func handleRequest(userID string) (response []byte, err error) {
	httpClient := &http.Client{}
	
	httpMethod := http.MethodGet
	
	req, err := http.NewRequest(httpMethod, base_URL+userID, nil)
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


func ListUser() {
	res, err := handleRequest("")
	if err != nil {
		fmt.Println("Error\n")
	}

	var apiResponse ListUsers
	err = json.Unmarshal(res, &apiResponse)

	if err != nil {
		fmt.Println(err)
	}
	
	for i:=0; i < len(apiResponse.Users); i++ {
		fmt.Print(apiResponse.Users[i],"\n")
	}
}


func Read(userID string) {
	res, err := handleRequest(userID)
	if err != nil {
		fmt.Print("Error\n")
	}

	var apiResponse User
	err = json.Unmarshal(res, &apiResponse)

	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(apiResponse, "\n")
}


func main() {
	fmt.Print("List of Users:\n")
	ListUser()

	fmt.Print("Requested user:\n")
	Read("{USER_ID}")
	/*CreateUser
	/*DelteUser
	/*UpdateUser
	*/
}



