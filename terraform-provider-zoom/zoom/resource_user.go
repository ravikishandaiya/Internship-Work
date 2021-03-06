package zoom

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var authToken string = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MjAzNjI4MjEsImlhdCI6MTYxOTc1ODAyMn0.vbRsXIQDly5la1Phi80IsVtkBoHMxTVr5EOChIEwsmI"

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

type GetUserResponse struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Type      int    `json:"type"`
}

type UpdateRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func resourceUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserCreate,
		ReadContext:   resourceOrderRead,
		UpdateContext: resourceOrderUpdate,
		DeleteContext: resourceOrderDelete,
		Schema: map[string]*schema.Schema{
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"firstname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lastname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func handlePostRequest(url string, httpMethod string, body []byte) (response []byte, err error) {
	httpClient := &http.Client{}

	var req *http.Request
	req, err = http.NewRequest(httpMethod, url, bytes.NewBuffer(body))
	if err != nil {
		return
	}

	req.Header.Add("Authorization", "Bearer "+authToken)
	req.Header.Add("Content-Type", "application/json")

	var resp *http.Response
	resp, err = httpClient.Do(req)
	if err != nil {
		return
	}

	if resp.StatusCode == 409 {
		er := errors.New("User Already Exists")
		err = er
		return
	}

	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	email := d.Get("email").(string)
	fn := d.Get("firstname").(string)
	ln := d.Get("lastname").(string)

	if fn == "" && ln == "" {
		clienturl := "https://api.zoom.us/v2/users/"
		userID := email
		url := fmt.Sprintf("%s%s", clienturl, userID)

		var err error

		httpMethod := http.MethodGet
		httpClient := &http.Client{}

		var req *http.Request
		req, err = http.NewRequest(httpMethod, url, nil)
		if err != nil {
			return diag.FromErr(err)
		}

		req.Header.Add("Authorization", "Bearer "+authToken)
		req.Header.Add("Content-Type", "application/json")

		var resp *http.Response
		resp, err = httpClient.Do(req)
		if err != nil {
			return diag.FromErr(err)
		}

		if resp.StatusCode == 404 {
			er := errors.New("User Doesn't Exist")
			err = er
			return diag.FromErr(err)
		}

		var response GetUserResponse
		response, _ = handleReadRequest(userID)

		eml := response.Email
		frstnme := response.FirstName
		id := response.Id
		lstnme := response.LastName

		d.SetId(id)

		d.Set("email", &eml)
		d.Set("firstname", &frstnme)
		d.Set("lastname", &lstnme)

		return diags
	}

	createUserRequest := CreateUserRequest{
		Action: "create",
		UserInfo: Userinfo{
			Email:     email,
			Type:      1,
			FirstName: fn,
			LastName:  ln,
		},
	}

	url := "https://api.zoom.us/v2/users/"
	httpMethod := http.MethodPost

	reqBody, err := json.Marshal(createUserRequest)

	if err != nil {
		return diag.FromErr(err)
	}

	var b []byte
	b, err = handlePostRequest(url, httpMethod, reqBody)

	if err != nil {
		return diag.FromErr(err)
	}

	var createUserResponse CreateUserResponse
	err = json.Unmarshal(b, &createUserResponse)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(createUserResponse.Id)

	resourceOrderRead(ctx, d, m)

	return diags
}

func handleReadRequest(userID string) (getUserResponse GetUserResponse, err error) {
	clienturl := "https://api.zoom.us/v2/users/"
	url := fmt.Sprintf("%s%s", clienturl, userID)

	httpMethod := http.MethodGet
	authToken := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MTk3NjQxNzgsImlhdCI6MTYxOTc1ODc3M30.0om8lKBPT-3FyB1_R1EH3a6uvDBDKrLhD0GOPLepeu4"

	httpClient := &http.Client{}

	var req *http.Request
	req, err = http.NewRequest(httpMethod, url, nil)
	if err != nil {
		return
	}

	req.Header.Add("Authorization", "Bearer "+authToken)
	req.Header.Add("Content-Type", "application/json")

	var resp *http.Response
	resp, err = httpClient.Do(req)
	if err != nil {
		return
	}

	var b []byte
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &getUserResponse)
	if err != nil {
		return
	}

	return
}

func resourceOrderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func UpdateUser(userID string, ois UpdateRequest) (err error) {
	var reqBody []byte
	reqBody, err = json.Marshal(ois)
	if err != nil {
		return
	}

	clienturl := "https://api.zoom.us/v2/users/"
	url := fmt.Sprintf("%s%s", clienturl, userID)
	httpMethod := http.MethodPatch

	httpClient := &http.Client{}
	var req *http.Request
	req, err = http.NewRequest(httpMethod, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return
	}

	req.Header.Add("Authorization", "Bearer "+authToken)
	req.Header.Add("Content-Type", "application/json")

	_, err = httpClient.Do(req)
	if err != nil {
		return
	}

	return
}

func resourceOrderUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	userID := d.Id()

	if d.HasChange("lastname") {
		ln := d.Get("lastname").(string)
		fn := d.Get("firstname").(string)
		ois := UpdateRequest{
			FirstName: fn,
			LastName:  ln,
		}

		err := UpdateUser(userID, ois)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceOrderRead(ctx, d, m)
}

func deleteUser(userID string) (err error) {
	clienturl := "https://api.zoom.us/v2/users/"
	url := fmt.Sprintf("%s%s", clienturl, userID)
	httpMethod := http.MethodDelete

	httpClient := &http.Client{}
	var req *http.Request
	req, err = http.NewRequest(httpMethod, url, nil)
	if err != nil {
		return
	}

	authToken := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MTk3NjQxNzgsImlhdCI6MTYxOTc1ODc3M30.0om8lKBPT-3FyB1_R1EH3a6uvDBDKrLhD0GOPLepeu4"
	req.Header.Add("Authorization", "Bearer "+authToken)
	req.Header.Add("Content-Type", "application/json")

	_, err = httpClient.Do(req)
	if err != nil {
		return
	}

	return
}

func resourceOrderDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	userID := d.Id()

	err := deleteUser(userID)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
