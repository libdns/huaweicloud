package model

import (
	"github.com/libdns/huaweicloud/sdk/core/utils"

	"strings"
)

// UpdatePublicZoneRequest Request Object
type UpdatePublicZoneRequest struct {

	// 待修改zone的ID
	ZoneId string `json:"zone_id"`

	Body *UpdatePublicZoneInfo `json:"body,omitempty"`
}

func (o UpdatePublicZoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePublicZoneRequest struct{}"
	}

	return strings.Join([]string{"UpdatePublicZoneRequest", string(data)}, " ")
}
