package images

import (
	"net/url"

	"github.com/huaweicloudsdk/golangsdk"
)

// `listURL` is a pure function. `listURL(c)` is a URL for which a GET
// request will respond with a list of images in the service `c`.
func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("images")
}

func createURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("images")
}

// `imageURL(c,i)` is the URL for the image identified by ID `i` in
// the service `c`.
func imageURL(c *golangsdk.ServiceClient, imageID string) string {
	return c.ServiceURL("images", imageID)
}

// `getURL(c,i)` is a URL for which a GET request will respond with
// information about the image identified by ID `i` in the service
// `c`.
func getURL(c *golangsdk.ServiceClient, imageID string) string {
	return imageURL(c, imageID)
}

func updateURL(c *golangsdk.ServiceClient, imageID string) string {
	return imageURL(c, imageID)
}

func deleteURL(c *golangsdk.ServiceClient, imageID string) string {
	return imageURL(c, imageID)
}

// builds next page full url based on current url
func nextPageURL(currentURL string, next string) (string, error) {
	base, err := url.Parse(currentURL)
	if err != nil {
		return "", err
	}
	rel, err := url.Parse(next)
	if err != nil {
		return "", err
	}
	return base.ResolveReference(rel).String(), nil
}
