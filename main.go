package main

import (
	"github.com/arthurflame/terraform-provider-spinnaker/spinnaker"
	"github.com/hashicorp/terraform/tree/main/internal/plugin"
	"github.com/hashicorp/terraform/tree/main/internal/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return spinnaker.Provider()
		},
	})
}
