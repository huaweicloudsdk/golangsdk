package serverutil

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/bootfromvolume"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/startstop"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
)

func CreateServer(client *gophercloud.ServiceClient, Opts bootfromvolume.CreateOptsExt) (*servers.Server, gophercloud.ErrorMessage) {
	var ecsErrMessage gophercloud.ErrorMessage
	if len(Opts.BlockDevice) > 0 {
		server, err := bootfromvolume.Create(client, Opts).Extract()
		if err != nil {
			TransErrorMessage(err, &ecsErrMessage)
		}
		return server, ecsErrMessage
	} else {
		server, err := servers.Create(client, Opts.CreateOptsBuilder).Extract()
		if err != nil {
			TransErrorMessage(err, &ecsErrMessage)
		}
		return server, ecsErrMessage
	}
}

func GetServerDetail(client *gophercloud.ServiceClient, server_id string) (*servers.Server, gophercloud.ErrorMessage) {
	var ecsErrMessage gophercloud.ErrorMessage
	server, err := servers.Get(client, server_id).Extract()
	if err != nil {
		TransErrorMessage(err, &ecsErrMessage)
	}
	return server, ecsErrMessage
}

func DeleteServer(client *gophercloud.ServiceClient, server_id string) gophercloud.ErrorMessage {
	var ecsErrMessage gophercloud.ErrorMessage
	err := servers.Delete(client, server_id).ExtractErr()
	if err != nil {
		TransErrorMessage(err, &ecsErrMessage)
	}
	return ecsErrMessage
}

func RebootServer(client *gophercloud.ServiceClient, server_id string, rtype string) gophercloud.ErrorMessage {
	var ecsErrMessage gophercloud.ErrorMessage
	var rebootMethod servers.RebootMethod

	if rtype == "SOFT" {
		rebootMethod = servers.SoftReboot
	} else {
		rebootMethod = servers.HardReboot
	}

	rebootOpts := &servers.RebootOpts{
		Type: rebootMethod,
	}
	err := servers.Reboot(client, server_id, rebootOpts).ExtractErr()
	if err != nil {
		TransErrorMessage(err, &ecsErrMessage)
	}
	return ecsErrMessage
}

func StartServer(client *gophercloud.ServiceClient, server_id string) gophercloud.ErrorMessage {
	var ecsErrMessage gophercloud.ErrorMessage
	err := startstop.Start(client, server_id).ExtractErr()
	if err != nil {
		TransErrorMessage(err, &ecsErrMessage)
	}
	return ecsErrMessage
}

func StopServer(client *gophercloud.ServiceClient, server_id string) gophercloud.ErrorMessage {
	var ecsErrMessage gophercloud.ErrorMessage
	err := startstop.Stop(client, server_id).ExtractErr()
	if err != nil {
		TransErrorMessage(err, &ecsErrMessage)
	}
	return ecsErrMessage
}
