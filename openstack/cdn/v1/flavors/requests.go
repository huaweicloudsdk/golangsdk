package flavors

import (
	"github.com/huaweicloudsdk/golangsdk"
	"github.com/huaweicloudsdk/golangsdk/pagination"
)

// List returns a single page of CDN flavors.
func List(c *golangsdk.ServiceClient) pagination.Pager {
	return pagination.NewPager(c, listURL(c), func(r pagination.PageResult) pagination.Page {
		return FlavorPage{pagination.SinglePageBase(r)}
	})
}

// Get retrieves a specific flavor based on its unique ID.
func Get(c *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(getURL(c, id), &r.Body, nil)
	return
}
