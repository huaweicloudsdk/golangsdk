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
		fmt.Println(err2)
	}
	computeClient, err3 := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if err3 != nil {
		fmt.Println(err3)
	}
	server_id := "a265c2ac-ee28-4112-a08e-e07b1b52c008"
	//server_id := "6259fdb3-8878-41cd-9f4f-b9b80f5e12da"
	err1 := serverutil.DeleteServer(computeClient, server_id)
	if err1.Message == "" {
		fmt.Println("删除成功")
	} else {
		fmt.Println(err1)
	}
}
