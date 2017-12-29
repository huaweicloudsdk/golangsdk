package instances

import "github.com/huaweicloudsdk/golangsdk"

func baseURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("instances")
}

func resourceURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("instances", id)
}

func userRootURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("instances", id, "root")
}

func actionURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("instances", id, "action")
}
