package model

import (
	"github.com/libdns/huaweicloud/sdk/core/utils"

	"strings"
)

// DeletePrivateZoneRequest Request Object
type DeletePrivateZoneRequest struct {

	// 待删除zone的ID。
	ZoneId string `json:"zone_id"`
}

func (o DeletePrivateZoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePrivateZoneRequest struct{}"
	}

	return strings.Join([]string{"DeletePrivateZoneRequest", string(data)}, " ")
}
