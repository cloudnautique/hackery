package rancher2

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var (
	descriptions map[string]string
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"rancher_url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    false,
				DefaultFunc: schema.EnvDefaultFunc("RANCHER_URL", ""),
				Description: descriptions["rancher_url"],
			},
			"access_key": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    false,
				DefaultFunc: schema.EnvDefaultFunc("RANCHER_ACCESS_KEY", ""),
				Description: descriptions["access_key"],
			},
			"secret_key": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    false,
				DefaultFunc: schema.EnvDefaultFunc("RANCHER_SECRET_KEY", ""),
				Description: descriptions["secret_key"],
			},
			"CACerts": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    false,
				DefaultFunc: schema.EnvDefaultFunc("RANCHER_CA_CERTS", ""),
				Description: descriptions["ca_certs"],
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"rancher2_cluster": resourceCluster(),
			"rancher2_project": resourceProject(),
		},

		ConfigureFunc: configureProvider,
	}
}

func init() {
	descriptions = map[string]string{
		"rancher_url": "The URL for the Rancher Server Installation",
		"access_key":  "Rancher Access Key",
		"secret_key":  "Rancher Secret Key",
		"ca_certs":    "Base64 encoded certs",
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	url := d.Get("rancher_url").(string)

	config := &ClientConfig{
		URL:       d.Get("rancher_url").(string),
		AccessKey: d.Get("access_key").(string),
		SecretKey: d.Get("secret_key").(string),
		CACerts:   d.Get("ca_certs").(string),
	}

	return config, nil
}
