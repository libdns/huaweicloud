package model

import (
	"github.com/libdns/huaweicloud/sdk/core/utils"

	"strings"
)

// ListApiVersionsResponse Response Object
type ListApiVersionsResponse struct {
	Versions       *ValuesItem `json:"versions,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListApiVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListApiVersionsResponse", string(data)}, " ")
}
