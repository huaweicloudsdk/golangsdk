package testing

import (
	"testing"

	"github.com/huaweicloudsdk/golangsdk/openstack/compute/v2/extensions/hypervisors"
	"github.com/huaweicloudsdk/golangsdk/pagination"
	"github.com/huaweicloudsdk/golangsdk/testhelper"
	"github.com/huaweicloudsdk/golangsdk/testhelper/client"
)

func TestListHypervisors(t *testing.T) {
	testhelper.SetupHTTP()
	defer testhelper.TeardownHTTP()
	HandleHypervisorListSuccessfully(t)

	pages := 0
	err := hypervisors.List(client.ServiceClient()).EachPage(func(page pagination.Page) (bool, error) {
		pages++

		actual, err := hypervisors.ExtractHypervisors(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 hypervisors, got %d", len(actual))
		}
		testhelper.CheckDeepEquals(t, HypervisorFake, actual[0])
		testhelper.CheckDeepEquals(t, HypervisorFake, actual[1])

		return true, nil
	})

	testhelper.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllHypervisors(t *testing.T) {
	testhelper.SetupHTTP()
	defer testhelper.TeardownHTTP()
	HandleHypervisorListSuccessfully(t)

	allPages, err := hypervisors.List(client.ServiceClient()).AllPages()
	testhelper.AssertNoErr(t, err)
	actual, err := hypervisors.ExtractHypervisors(allPages)
	testhelper.AssertNoErr(t, err)
	testhelper.CheckDeepEquals(t, HypervisorFake, actual[0])
	testhelper.CheckDeepEquals(t, HypervisorFake, actual[1])
}
