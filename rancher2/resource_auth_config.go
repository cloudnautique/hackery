package rancher2

import "github.com/hashicorp/terraform/helper/schema"

func resourceAuthConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceAuthConfigCreate,
		Read:   resourceAuthConfigRead,
		Update: resourceAuthConfigUpdate,
		Delete: resourceAuthConfigDelete,

		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceAuthConfigCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAuthConfigRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAuthConfigUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAuthConfigDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
