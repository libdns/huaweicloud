package model

import (
	"github.com/libdns/huaweicloud/sdk/core/utils"

	"strings"
)

// ListApiVersionsRequest Request Object
type ListApiVersionsRequest struct {
}

func (o ListApiVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiVersionsRequest struct{}"
	}

	return strings.Join([]string{"ListApiVersionsRequest", string(data)}, " ")
}
