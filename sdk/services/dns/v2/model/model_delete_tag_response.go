package model

import (
	"github.com/libdns/huaweicloud/sdk/core/utils"

	"strings"
)

// DeleteTagResponse Response Object
type DeleteTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteTagResponse", string(data)}, " ")
}
