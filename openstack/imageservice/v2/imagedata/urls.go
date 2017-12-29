package imagedata

import "github.com/huaweicloudsdk/golangsdk"

// `imageDataURL(c,i)` is the URL for the binary image data for the
// image identified by ID `i` in the service `c`.
func uploadURL(c *golangsdk.ServiceClient, imageID string) string {
	return c.ServiceURL("images", imageID, "file")
}

func downloadURL(c *golangsdk.ServiceClient, imageID string) string {
	return uploadURL(c, imageID)
}
