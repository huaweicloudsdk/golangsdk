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
	var serverList = [...]string{
              	"2ff03277-3448-45bd-aaf3-7d05536b7519",
		}
	for _, server_id := range serverList {
		err1 := serverutil.DeleteServer(computeClient, server_id)
		if err1.Message == "" {
			fmt.Println("删除成功")
		} else {
			fmt.Println(err1)
		}
	}
}
