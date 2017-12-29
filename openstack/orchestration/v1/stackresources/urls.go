package stackresources

import "github.com/huaweicloudsdk/golangsdk"

func findURL(c *golangsdk.ServiceClient, stackName string) string {
	return c.ServiceURL("stacks", stackName, "resources")
}

func listURL(c *golangsdk.ServiceClient, stackName, stackID string) string {
	return c.ServiceURL("stacks", stackName, stackID, "resources")
}

func getURL(c *golangsdk.ServiceClient, stackName, stackID, resourceName string) string {
	return c.ServiceURL("stacks", stackName, stackID, "resources", resourceName)
}

func metadataURL(c *golangsdk.ServiceClient, stackName, stackID, resourceName string) string {
	return c.ServiceURL("stacks", stackName, stackID, "resources", resourceName, "metadata")
}

func listTypesURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("resource_types")
}

func schemaURL(c *golangsdk.ServiceClient, typeName string) string {
	return c.ServiceURL("resource_types", typeName)
}

func templateURL(c *golangsdk.ServiceClient, typeName string) string {
	return c.ServiceURL("resource_types", typeName, "template")
}
