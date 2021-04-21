package provider

import (
	"github.com.hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"zoom_api/client"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"address": {
				Type: schema.TypeString,
				Required: true,
			},
			"token": {
				Type: schema.TypeString,
				Required: true,
			},
		},
		ResourceMap: map[string]*schema.Resource{
			"provider_user": resourceUser(),
		},
		ConfigureFunc: providerConfigure,
	}
} 

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	address := d.Get("address").(string)
	token := d.Get("token").(string)
	return client.NewClient(address, token), nil
}
