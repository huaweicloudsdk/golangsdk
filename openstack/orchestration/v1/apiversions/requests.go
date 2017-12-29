package apiversions

import (
	"github.com/huaweicloudsdk/golangsdk"
	"github.com/huaweicloudsdk/golangsdk/pagination"
)

// ListVersions lists all the Neutron API versions available to end-users
func ListVersions(c *golangsdk.ServiceClient) pagination.Pager {
	return pagination.NewPager(c, apiVersionsURL(c), func(r pagination.PageResult) pagination.Page {
		return APIVersionPage{pagination.SinglePageBase(r)}
	})
}
