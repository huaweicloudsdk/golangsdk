package tokens

import "github.com/huaweicloudsdk/golangsdk"

func tokenURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("auth", "tokens")
}
