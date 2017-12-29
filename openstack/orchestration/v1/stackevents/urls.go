package stackevents

import "github.com/huaweicloudsdk/golangsdk"

func findURL(c *golangsdk.ServiceClient, stackName string) string {
	return c.ServiceURL("stacks", stackName, "events")
}

func listURL(c *golangsdk.ServiceClient, stackName, stackID string) string {
	return c.ServiceURL("stacks", stackName, stackID, "events")
}

func listResourceEventsURL(c *golangsdk.ServiceClient, stackName, stackID, resourceName string) string {
	return c.ServiceURL("stacks", stackName, stackID, "resources", resourceName, "events")
}

func getURL(c *golangsdk.ServiceClient, stackName, stackID, resourceName, eventID string) string {
	return c.ServiceURL("stacks", stackName, stackID, "resources", resourceName, "events", eventID)
}
