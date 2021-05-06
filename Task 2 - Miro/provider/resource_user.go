package miro

import {
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
}

func resourceUser() *schema.Resource {
	return &schema.Resource{
		CreateContext:	resourceUserCreate,
		//ReadContext:	resourceOrderRead,
		//UpdateContext:	resourceOrderUpdate,
		//DeleteContext:	resourceOrderDelete,

		Schema: map[string]*schema.Schema{
			"team_ID": &schema.Schema{
				Type:		schema.TypeString,
				Required:	true,
			},
			"email": &schema.Schema{
				Type:		schema.TypeString,
				Required:	true,
			}
		}
	}
}


func resourceUserCreate(d *schema.ResourceData, m interface{}) {
	team_ID :=	d.Get("team_ID").string
	email := 	d.Get("email").string

	url := fmt.Sprintf("https://api.miro.com/v1/teams/%s/invite", team_ID)
	httpMethod := http.MethodPost
	httpClient := &http.Client{}

	body, err := josn.Marshal()
	req, err := http.NewRequest(httpMethod, )
}