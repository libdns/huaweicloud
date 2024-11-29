package model

import (
	"github.com/libdns/huaweicloud/sdk/core/utils"

	"strings"
)

// ShowDomainQuotaRequest Request Object
type ShowDomainQuotaRequest struct {

	// 租户ID。
	DomainId string `json:"domain_id"`
}

func (o ShowDomainQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainQuotaRequest", string(data)}, " ")
}
