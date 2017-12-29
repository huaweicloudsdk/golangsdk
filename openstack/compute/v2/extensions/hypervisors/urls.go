package hypervisors

import "github.com/huaweicloudsdk/golangsdk"

func hypervisorsListDetailURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("os-hypervisors", "detail")
}
