package miro

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

//Provider 

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourceMap: map[string]*schema.Resource{
			"miro_user": resourceUser(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

