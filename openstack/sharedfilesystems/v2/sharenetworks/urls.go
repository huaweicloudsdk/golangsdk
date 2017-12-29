package sharenetworks

import "github.com/huaweicloudsdk/golangsdk"

func createURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("share-networks")
}

func deleteURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("share-networks", id)
}

func listDetailURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("share-networks", "detail")
}

func getURL(c *golangsdk.ServiceClient, id string) string {
	return deleteURL(c, id)
}

func updateURL(c *golangsdk.ServiceClient, id string) string {
	return deleteURL(c, id)
}

func addSecurityServiceURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("share-networks", id, "action")
}

func removeSecurityServiceURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("share-networks", id, "action")
}
