package buildinfo

import "github.com/huaweicloudsdk/golangsdk"

func getURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("build_info")
}
