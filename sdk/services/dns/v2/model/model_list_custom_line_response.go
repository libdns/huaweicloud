package model

import (
	"github.com/libdns/huaweicloud/sdk/core/utils"

	"strings"
)

// ListCustomLineResponse Response Object
type ListCustomLineResponse struct {

	// 线路列表。
	Lines *[]Line `json:"lines,omitempty"`

	Metadata       *Metadata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListCustomLineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomLineResponse struct{}"
	}

	return strings.Join([]string{"ListCustomLineResponse", string(data)}, " ")
}
