package buildinfo

import "github.com/huaweicloudsdk/golangsdk"

// Get retreives data for the given stack template.
func Get(c *golangsdk.ServiceClient) (r GetResult) {
	_, r.Err = c.Get(getURL(c), &r.Body, nil)
	return
}
