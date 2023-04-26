// 自动生成模板Fragment2words
package video

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Fragment2words 结构体
type Fragment2words struct {
	global.GVA_MODEL
	FragmentId *int   `json:"fragmentId" form:"fragmentId" gorm:"column:fragment_id;comment:碎片视频ID;size:19;"`
	Removed    *bool  `json:"removed" form:"removed" gorm:"column:removed;comment:是否移除（0-默认 1-已移除）;"`
	Words      string `json:"words" form:"words" gorm:"column:words;comment:单词;size:64;"`
}

// TableName Fragment2words 表名
func (Fragment2words) TableName() string {
	return "fragment2words"
}
