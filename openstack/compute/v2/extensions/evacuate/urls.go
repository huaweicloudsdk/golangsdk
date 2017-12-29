package evacuate

import (
	"github.com/huaweicloudsdk/golangsdk"
)

func actionURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "action")
}
