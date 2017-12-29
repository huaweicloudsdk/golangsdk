package testing

import (
	"testing"

	"github.com/huaweicloudsdk/golangsdk/openstack/compute/v2/extensions/lockunlock"
	th "github.com/huaweicloudsdk/golangsdk/testhelper"
	"github.com/huaweicloudsdk/golangsdk/testhelper/client"
)

const serverID = "{serverId}"

func TestLock(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	mockStartServerResponse(t, serverID)

	err := lockunlock.Lock(client.ServiceClient(), serverID).ExtractErr()
	th.AssertNoErr(t, err)
}

func TestUnlock(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	mockStopServerResponse(t, serverID)

	err := lockunlock.Unlock(client.ServiceClient(), serverID).ExtractErr()
	th.AssertNoErr(t, err)
}
