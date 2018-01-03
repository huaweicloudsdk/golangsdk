package aggregates

import "github.com/huaweicloudsdk/golangsdk"

func aggregatesListURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-aggregates")
}
