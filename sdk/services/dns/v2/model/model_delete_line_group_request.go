package model

import (
	"github.com/libdns/huaweicloud/sdk/core/utils"

	"strings"
)

// DeleteLineGroupRequest Request Object
type DeleteLineGroupRequest struct {

	// 线路分组ID。
	LinegroupId string `json:"linegroup_id"`
}

func (o DeleteLineGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLineGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteLineGroupRequest", string(data)}, " ")
}
