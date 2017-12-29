package shares

import "github.com/huaweicloudsdk/golangsdk"

func createURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("shares")
}

func deleteURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("shares", id)
}

func getURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("shares", id)
}

func getExportLocationsURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("shares", id, "export_locations")
}

func grantAccessURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("shares", id, "action")
}
