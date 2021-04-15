package provider

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceItem() schema.Resource {
	return &schema..Resource{
		Create: resourceCreateItem,
		Read: resourceReadItem,
		Update: resourceUpdateItem,
		Delete: resourceDeleteItem,

		Schema: map[string]*schema.Schema{
			"first name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			}
			"last name": {
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
	apiClient := m.(*client.Client)
}
