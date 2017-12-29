// +build acceptance compute hypervisors

package v2

import (
	"testing"

	"github.com/huaweicloudsdk/golangsdk/acceptance/clients"
	"github.com/huaweicloudsdk/golangsdk/acceptance/tools"
	"github.com/huaweicloudsdk/golangsdk/openstack/compute/v2/extensions/hypervisors"
)

func TestHypervisorsList(t *testing.T) {
	client, err := clients.NewComputeV2Client()
	if err != nil {
		t.Fatalf("Unable to create a compute client: %v", err)
	}

	allPages, err := hypervisors.List(client).AllPages()
	if err != nil {
		t.Fatalf("Unable to list hypervisors: %v", err)
	}

	allHypervisors, err := hypervisors.ExtractHypervisors(allPages)
	if err != nil {
		t.Fatalf("Unable to extract hypervisors")
	}

	for _, h := range allHypervisors {
		tools.PrintResource(t, h)
	}
}
