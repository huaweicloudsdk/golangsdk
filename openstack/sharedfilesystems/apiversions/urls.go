package apiversions

import (
	"net/url"
	"strings"

	"github.com/huaweicloudsdk/golangsdk"
)

func getURL(c *golangsdk.ServiceClient, version string) string {
	u, _ := url.Parse(c.ServiceURL(""))
	u.Path = "/" + strings.TrimRight(version, "/") + "/"
	return u.String()
}

func listURL(c *golangsdk.ServiceClient) string {
	u, _ := url.Parse(c.ServiceURL(""))
	u.Path = "/"
	return u.String()
}
