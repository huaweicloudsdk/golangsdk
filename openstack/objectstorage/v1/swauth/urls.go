package swauth

import "github.com/huaweicloudsdk/golangsdk"

func getURL(c *golangsdk.ProviderClient) string {
	return c.IdentityBase + "auth/v1.0"
}
