package testing

import (
	"testing"

	"github.com/huaweicloudsdk/golangsdk/openstack/compute/v2/extensions/aggregates"
	"github.com/huaweicloudsdk/golangsdk/pagination"
	"github.com/huaweicloudsdk/golangsdk/testhelper"
	"github.com/huaweicloudsdk/golangsdk/testhelper/client"
)

func TestListAggregates(t *testing.T) {
	testhelper.SetupHTTP()
	defer testhelper.TeardownHTTP()
	HandleListSuccessfully(t)

	pages := 0
	err := aggregates.List(client.ServiceClient()).EachPage(func(page pagination.Page) (bool, error) {
		pages++

		actual, err := aggregates.ExtractAggregates(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 aggregates, got %d", len(actual))
		}
		testhelper.CheckDeepEquals(t, FirstFakeAggregate, actual[0])
		testhelper.CheckDeepEquals(t, SecondFakeAggregate, actual[1])

		return true, nil
	})

	testhelper.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}
