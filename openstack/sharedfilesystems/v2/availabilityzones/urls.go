package availabilityzones

import "github.com/huaweicloudsdk/golangsdk"

func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("os-availability-zone")
}
