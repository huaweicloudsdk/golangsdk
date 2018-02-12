package main

import (
	"fmt"
	"os"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/serverutil"
)

func main() {
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: "https://iam.eu-de.otc.t-systems.com/v3",
		Username:         "DickInPussy",
		Password:         "cnp200@HW",
		DomainName:       "DickInPussy",
		TenantName:       "eu-de",
	}

	provider, err2 := openstack.AuthenticatedClient(opts)
	if err2 != nil {
		fmt.Println("====================", err2)
	}
	computeClient, err3 := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if err3 != nil {
		fmt.Println(err3)
	}

	server_id := "d9aae046-1734-41b8-852e-bdac422a38c4"
	server, err1 := serverutil.GetServerDetail(computeClient, server_id)
	if err1.Message == "" {
		fmt.Println(server)
	} else {
		fmt.Println(err1)
	}
}

