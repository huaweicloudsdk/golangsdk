package sharetypes

import "github.com/huaweicloudsdk/golangsdk"

func createURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("types")
}

func deleteURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("types", id)
}

func listURL(c *golangsdk.ServiceClient) string {
	return createURL(c)
}

func getDefaultURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("types", "default")
}

func getExtraSpecsURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("types", id, "extra_specs")
}

func setExtraSpecsURL(c *golangsdk.ServiceClient, id string) string {
	return getExtraSpecsURL(c, id)
}

func unsetExtraSpecsURL(c *golangsdk.ServiceClient, id string, key string) string {
	return c.ServiceURL("types", id, "extra_specs", key)
}

func showAccessURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("types", id, "share_type_access")
}

func addAccessURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("types", id, "action")
}

func removeAccessURL(c *golangsdk.ServiceClient, id string) string {
	return addAccessURL(c, id)
}
