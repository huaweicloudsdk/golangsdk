package serviceassets

import "github.com/huaweicloudsdk/golangsdk"

func deleteURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("services", id, "assets")
}
