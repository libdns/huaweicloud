package model

import (
	"github.com/libdns/huaweicloud/sdk/core/utils"

	"strings"
)

// UpdatePublicZoneStatusRequest Request Object
type UpdatePublicZoneStatusRequest struct {

	// 待设置状态Zone的ID
	ZoneId string `json:"zone_id"`

	Body *UpdatePublicZoneStatusRequestBody `json:"body,omitempty"`
}

func (o UpdatePublicZoneStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePublicZoneStatusRequest struct{}"
	}

	return strings.Join([]string{"UpdatePublicZoneStatusRequest", string(data)}, " ")
}
