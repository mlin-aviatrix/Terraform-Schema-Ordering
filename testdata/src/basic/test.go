package basic

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceTestSchemaOrdering() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"vpc_id": {
				Description: "",
				Default: "",
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{},
				Type: schema.TypeString,
			},
		},
	}
}