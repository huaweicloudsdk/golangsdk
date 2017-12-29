package base

import "github.com/huaweicloudsdk/golangsdk"

// Get retrieves the home document, allowing the user to discover the
// entire API.
func Get(c *golangsdk.ServiceClient) (r GetResult) {
	_, r.Err = c.Get(getURL(c), &r.Body, nil)
	return
}

// Ping retrieves a ping to the server.
func Ping(c *golangsdk.ServiceClient) (r PingResult) {
	_, r.Err = c.Get(pingURL(c), nil, &golangsdk.RequestOpts{
		OkCodes:     []int{204},
		MoreHeaders: map[string]string{"Accept": ""},
	})
	return
}
