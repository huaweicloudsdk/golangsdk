package flavors

import "github.com/huaweicloudsdk/golangsdk"

func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("flavors")
}

func getURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("flavors", id)
}
