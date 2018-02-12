package serverutil

import (
	"encoding/json"
	"fmt"

	"github.com/gophercloud/gophercloud"
)

func TransErrorMessage(err error, ecsErrMessage *gophercloud.ErrorMessage) {
	errString := fmt.Sprintf("%v", err)
	isJson := json.Unmarshal([]byte(errString), &ecsErrMessage)
	if isJson != nil {
		ecsErrMessage.ErrCode = "UnknowCode"
		ecsErrMessage.Message = errString
	}
}
