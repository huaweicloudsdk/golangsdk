package stacktemplates

import "github.com/huaweicloudsdk/golangsdk"

func getURL(c *golangsdk.ServiceClient, stackName, stackID string) string {
	return c.ServiceURL("stacks", stackName, stackID, "template")
}

func validateURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("validate")
}
