package schedulerstats

import "github.com/huaweicloudsdk/golangsdk"

func storagePoolsListURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("scheduler-stats", "get_pools")
}
