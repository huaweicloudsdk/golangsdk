package testing

import (
	"testing"

	"github.com/huaweicloudsdk/golangsdk/openstack/compute/v2/extensions/migrate"
	th "github.com/huaweicloudsdk/golangsdk/testhelper"
	"github.com/huaweicloudsdk/golangsdk/testhelper/client"
)

const serverID = "b16ba811-199d-4ffd-8839-ba96c1185a67"

func TestMigrate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	mockMigrateResponse(t, serverID)

	err := migrate.Migrate(client.ServiceClient(), serverID).ExtractErr()
	th.AssertNoErr(t, err)
}
