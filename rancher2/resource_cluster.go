package rancher2

import "github.com/hashicorp/terraform/helper/schema"

func resourceCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceClusterCreate,
		Read:   resourceClusterRead,
		Update: resourceClusterUpdate,
		Delete: resourceClusterDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceClusterCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceClusterRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceClusterUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceClusterDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
