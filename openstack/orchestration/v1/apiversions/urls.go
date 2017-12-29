package apiversions

import "github.com/huaweicloudsdk/golangsdk"

func apiVersionsURL(c *golangsdk.ServiceClient) string {
	return c.Endpoint
}
