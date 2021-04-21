package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"Zoom/provider"
)
 
func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider(),
	})
}
