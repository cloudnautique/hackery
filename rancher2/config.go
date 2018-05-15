package rancher2

type ClientConfig struct {
	URL       string
	AccessKey string
	SecretKey string
	CACerts string
}


func (c *ClientConfig) Management() *