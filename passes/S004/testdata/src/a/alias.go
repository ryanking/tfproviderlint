package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func falias() {
	_ = s.Schema{ // want "schema should not enable Required and configure Default"
		Required: true,
		Default:  true,
	}

	_ = map[string]*s.Schema{
		"name": { // want "schema should not enable Required and configure Default"
			Required: true,
			Default:  true,
		},
	}
}
