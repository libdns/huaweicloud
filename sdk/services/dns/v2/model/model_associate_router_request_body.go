package model

import (
	"github.com/libdns/huaweicloud/sdk/core/utils"

	"strings"
)

type AssociateRouterRequestBody struct {
	Router *Router `json:"router"`
}

func (o AssociateRouterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateRouterRequestBody struct{}"
	}

	return strings.Join([]string{"AssociateRouterRequestBody", string(data)}, " ")
}
