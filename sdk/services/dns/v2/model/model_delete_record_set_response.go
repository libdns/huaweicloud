package model

import (
	"strings"

	"github.com/libdns/huaweicloud/sdk/core/utils"
)

// DeleteRecordSetResponse Response Object
type DeleteRecordSetResponse struct {

	// Record Set的ID。
	Id *string `json:"id,omitempty"`

	// Record Set的名称。
	Name *string `json:"name,omitempty"`

	// Record Set的描述信息。
	Description *string `json:"description,omitempty"`

	// 托管该记录的zone_id。
	ZoneId *string `json:"zone_id,omitempty"`

	// 托管该记录的zone_name。
	ZoneName *string `json:"zone_name,omitempty"`

	// 记录类型。
	Type *string `json:"type,omitempty"`

	// 解析记录在本地DNS服务器的缓存时间，缓存时间越长更新生效越慢，以秒为单位。
	Ttl *int32 `json:"ttl,omitempty"`

	// 域名解析后的值。
	Records *[]string `json:"records,omitempty"`

	// 创建时间。
	CreateAt *string `json:"create_at,omitempty"`

	// 更新时间。
	UpdateAt *string `json:"update_at,omitempty"`

	// 资源状态。
	Status *string `json:"status,omitempty"`

	// 标识是否由系统默认生成，系统默认生成的Record Set不能删除。
	Default *bool `json:"default,omitempty"`

	// 该Record Set所属的项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	Links          *PageLink `json:"links,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o DeleteRecordSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRecordSetResponse struct{}"
	}

	return strings.Join([]string{"DeleteRecordSetResponse", string(data)}, " ")
}
