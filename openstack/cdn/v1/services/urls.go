package services

import "github.com/huaweicloudsdk/golangsdk"

func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("services")
}

func createURL(c *golangsdk.ServiceClient) string {
	return listURL(c)
}

func getURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("services", id)
}

func updateURL(c *golangsdk.ServiceClient, id string) string {
	return getURL(c, id)
}

func deleteURL(c *golangsdk.ServiceClient, id string) string {
	return getURL(c, id)
}
