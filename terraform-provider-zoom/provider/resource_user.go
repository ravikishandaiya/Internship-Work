package provider

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"zoom"
)


func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceCreateItem,
		Read: resourceReadItem,
		Update: resourceUpdateItem,
		Delete: resourceDeleteItem,

		Schema: map[string]*schema.Schema{
			"first_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			}
			"last_name": {
				Type: schema.TypeString,
				Required: true,
			}
			"email": {
				Type: schema.TypeString,
				Required: true,
			}
		}
	}
}


func resourceCreateItem(d *schema.ResourceData, m interface{}) error {
	address := m.address,
	token := m.token

	createUserRequest := CreateUserRequest{
		Action: "create",
		UserInfo: Userinfo{
			Email:     d.email,
			Type:      1,
			FirstName: d.first_name,
			LastName:  d.last_name,
		},
	}

	client.CreateUser(address, token, CreateUserRequest)
}



func resourceReadItem(d *schema.ResourceData, m interface{}) error {
	address := m.address,
	token := m.token
	
	client.ListUser(address, token)
}

