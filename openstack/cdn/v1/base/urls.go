package base

import "github.com/huaweicloudsdk/golangsdk"

func getURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL()
}

func pingURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("ping")
}
