package model

import (
	"github.com/libdns/huaweicloud/sdk/core/utils"

	"strings"
)

// ListTagsResponse Response Object
type ListTagsResponse struct {

	// 标签列表。
	Tags           *[]TagValues `json:"tags,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagsResponse struct{}"
	}

	return strings.Join([]string{"ListTagsResponse", string(data)}, " ")
}
