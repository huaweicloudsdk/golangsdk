package hypervisors

import (
	"github.com/huaweicloudsdk/golangsdk"
	"github.com/huaweicloudsdk/golangsdk/pagination"
)

// List makes a request against the API to list hypervisors.
func List(client *golangsdk.ServiceClient) pagination.Pager {
	return pagination.NewPager(client, hypervisorsListDetailURL(client), func(r pagination.PageResult) pagination.Page {
		return HypervisorPage{pagination.SinglePageBase(r)}
	})
}
