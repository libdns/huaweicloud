package model

import (
	"github.com/libdns/huaweicloud/sdk/core/utils"

	"strings"
)

// ListRecordSetsWithLineResponse Response Object
type ListRecordSetsWithLineResponse struct {
	Links *PageLink `json:"links,omitempty"`

	// recordset列表信息。
	Recordsets *[]QueryRecordSetWithLineAndTagsResp `json:"recordsets,omitempty"`

	Metadata       *Metadata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListRecordSetsWithLineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecordSetsWithLineResponse struct{}"
	}

	return strings.Join([]string{"ListRecordSetsWithLineResponse", string(data)}, " ")
}
