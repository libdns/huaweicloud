package model

import (
	"github.com/libdns/huaweicloud/sdk/core/utils"

	"strings"
)

// RestorePtrRecordResponse Response Object
type RestorePtrRecordResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RestorePtrRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestorePtrRecordResponse struct{}"
	}

	return strings.Join([]string{"RestorePtrRecordResponse", string(data)}, " ")
}
