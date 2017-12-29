package zones

import "github.com/huaweicloudsdk/golangsdk"

func baseURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("zones")
}

func zoneURL(c *golangsdk.ServiceClient, zoneID string) string {
	return c.ServiceURL("zones", zoneID)
}
