// 自动生成模板PlayList
package video

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// PlayList 结构体
type PlayList struct {
	global.GVA_MODEL
	Name          string `json:"name" form:"name" gorm:"column:name;comment:播单名称;size:64;"`
	OnlineVersion string `json:"onlineVersion" form:"onlineVersion" gorm:"column:online_version;comment:上线版本;size:16;"`
}

// TableName PlayList 表名
func (PlayList) TableName() string {
	return "play_list"
}
