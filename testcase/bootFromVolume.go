package main

import (
	"fmt"
	"os"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	createOpts "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/bootfromvolume"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/serverutil"
)

func main() {
	serverCreateOpts := servers.CreateOpts{
		Name:      "bootFromVolume-basic-sina-SAS",
		FlavorRef: "normal1",
		Networks: []servers.Network{servers.Network{UUID: "c29a24b7-3ca1-4bf8-95db-fde5faf6c512"}},
		Metadata:  map[string]string{"hello":"world"},
		SecurityGroups: []string{"default"},
		UserData: []byte("IyEvYmluL2Jhc2gKZWNobyAncm9vdDokNiRPRjMxdlo0cm1CWUpvZzBLJE1ldlVrS3dSYVI0SmM2QVRaSi9lT2s4Q0ZFWUo1NFVSOFlvc2xsZUd0RERIRHd4TWRuU3lJcUw0WS9jN0MvSFlRcmRVZG45WXJKQnlhRnlvZm5ybjYuJyB8IGNocGFzc3dkIC1l"),
		AvailabilityZone:"eu-de-01",
	}

	Opts := []createOpts.BlockDevice{
		createOpts.BlockDevice{
			BootIndex:           0,
			DeleteOnTermination: true,
			DestinationType:     "volume",
			SourceType:          "image",
			VolumeSize:          40,
			UUID:                "53b2fbb5-ef2c-412a-bb0a-571436fa78ad",
			VolumeType:          "SAS",
		},
		createOpts.BlockDevice{
			BootIndex:           -1,
			DeleteOnTermination: true,
			DestinationType:     "volume",
			SourceType:          "blank",
			VolumeSize:          10,
		},
	}

	volumeOpts := createOpts.CreateOptsExt{
		CreateOptsBuilder: serverCreateOpts,
		BlockDevice: Opts,
		
	}
	
	envOpts := gophercloud.AuthOptions{
		IdentityEndpoint: "https://iam.eu-de.otc.t-systems.com/v3",
		Username:         "DickInPussy",
		Password:         "cnp200@HW",
		DomainName:       "DickInPussy",
		TenantName:       "eu-de",
	}

	provider, err2 := openstack.AuthenticatedClient(envOpts)
	if err2 != nil {
		fmt.Println(err2)
	}
	computeClient, err3 := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if err3 != nil {
		fmt.Println(err3)
	}
	server, err1 := serverutil.CreateServer(computeClient, volumeOpts)
	if err1.Message == "" {
		fmt.Println(server.VolumeAttached)
	} else {
		fmt.Println(err1)
	}
}

