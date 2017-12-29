package stacks

import "github.com/huaweicloudsdk/golangsdk"

func createURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("stacks")
}

func adoptURL(c *golangsdk.ServiceClient) string {
	return createURL(c)
}

func listURL(c *golangsdk.ServiceClient) string {
	return createURL(c)
}

func getURL(c *golangsdk.ServiceClient, name, id string) string {
	return c.ServiceURL("stacks", name, id)
}

func updateURL(c *golangsdk.ServiceClient, name, id string) string {
	return getURL(c, name, id)
}

func deleteURL(c *golangsdk.ServiceClient, name, id string) string {
	return getURL(c, name, id)
}

func previewURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("stacks", "preview")
}

func abandonURL(c *golangsdk.ServiceClient, name, id string) string {
	return c.ServiceURL("stacks", name, id, "abandon")
}
