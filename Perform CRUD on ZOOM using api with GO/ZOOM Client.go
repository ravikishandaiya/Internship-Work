package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"bytes"
)

// login credentials

const JWT_Token string = "JWT TOken"
const base_URL string = "https://api.zoom.us/v2/users/"

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

func handleReadRequest(url string, httpMethod string) (response []byte, err error) {
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

func ListUser() {
	res, err := handleReadRequest(base_URL, http.MethodGet)
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

// read a specific user

func Read(userID string) {
	res, err := handleReadRequest(base_URL+userID, http.MethodGet)
	if err != nil {
		fmt.Print("Error\n")
	}

	var apiResponse User
	err = json.Unmarshal(res, &apiResponse)

	if err != nil {
		fmt.Print(err)
	}

	fmt.Print("User:/n", apiResponse, "\n")
}


// struct to send data to html body and for mapping data from responce 

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

//Handle create request

func handleCreateRequest(url string, httpMethod string, body []byte) (response []byte, err error) {
	httpClient := &http.Client{}
	
	var req *http.Request
	req, err = http.NewRequest(httpMethod, url, bytes.NewBuffer(body))
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



// function to add new user


func createUser() {
	createUserRequest := CreateUserRequest{
		Action: "create",
		UserInfo: Userinfo{
			Email:     "abc@gmail.com",
			Type:      1,
			FirstName: "Ravi",
			LastName:  "Daiya",
		},
	}

	var reqBody []byte
	var err error
	reqBody, err = json.Marshal(createUserRequest)
	if err != nil {
		return
	}

	httpMethod := http.MethodPost

	var b []byte
	b, err = handleCreateRequest(base_URL, httpMethod, reqBody)

	if err != nil {
		return
	}

	var createUserResponse CreateUserResponse
	err = json.Unmarshal(b, &createUserResponse)
	if err != nil {
		return
	}

	fmt.Println(createUserResponse)
}


func main() {
	fmt.Print("List of Users:\n")
	ListUser() 
	fmt.Print("\nRequested user:\n")
	Read({User_ID})
	fmt.Print("\ncreate a user:\n")
	createUser()
	
	/*DelteUser
	/*UpdateUser
	*/
}



