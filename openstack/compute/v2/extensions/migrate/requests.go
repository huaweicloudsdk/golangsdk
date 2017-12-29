package migrate

import (
	"github.com/huaweicloudsdk/golangsdk"
)

// Migrate will initiate a migration of the instance to another host.
func Migrate(client *golangsdk.ServiceClient, id string) (r MigrateResult) {
	_, r.Err = client.Post(actionURL(client, id), map[string]interface{}{"migrate": nil}, nil, nil)
	return
}
