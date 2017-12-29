package testing

import (
	"testing"

	"github.com/huaweicloudsdk/golangsdk/openstack/compute/v2/extensions/resetstate"
	th "github.com/huaweicloudsdk/golangsdk/testhelper"
	"github.com/huaweicloudsdk/golangsdk/testhelper/client"
)

const serverID = "b16ba811-199d-4ffd-8839-ba96c1185a67"

func TestResetState(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	mockResetStateResponse(t, serverID, "active")

	err := resetstate.ResetState(client.ServiceClient(), serverID, "active").ExtractErr()
	th.AssertNoErr(t, err)
}
