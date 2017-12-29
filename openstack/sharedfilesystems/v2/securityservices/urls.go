package securityservices

import "github.com/huaweicloudsdk/golangsdk"

func createURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("security-services")
}

func deleteURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("security-services", id)
}

func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("security-services", "detail")
}

func getURL(c *golangsdk.ServiceClient, id string) string {
	return deleteURL(c, id)
}

func updateURL(c *golangsdk.ServiceClient, id string) string {
	return deleteURL(c, id)
}
