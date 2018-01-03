package aggregates

import (
	"github.com/huaweicloudsdk/golangsdk"
	"github.com/huaweicloudsdk/golangsdk/pagination"
)

// List makes a request against the API to list aggregates.
func List(client *gophercloud.ServiceClient) pagination.Pager {
	return pagination.NewPager(client, aggregatesListURL(client), func(r pagination.PageResult) pagination.Page {
		return AggregatesPage{pagination.SinglePageBase(r)}
	})
}
