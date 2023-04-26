// 自动生成模板Fragment
package video

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Fragment 结构体
type Fragment struct {
      global.GVA_MODEL
      Author  string `json:"author" form:"author" gorm:"column:author;comment:碎片视频作者;size:64;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:碎片视频名称;size:64;"`
      Status  *bool `json:"status" form:"status" gorm:"column:status;comment:碎片视频审核状态（0 未审核 1 初审中 3 初审完成 4 复审审核中 5 复审完成）;"`
      Subtitles  string `json:"subtitles" form:"subtitles" gorm:"column:subtitles;comment:碎片视频字幕;"`
      UploadedAt  *time.Time `json:"uploadedAt" form:"uploadedAt" gorm:"column:uploaded_at;comment:碎片视频上传时间;"`
      Url  string `json:"url" form:"url" gorm:"column:url;comment:碎片视频存储地址;size:2000;"`
}


// TableName Fragment 表名
func (Fragment) TableName() string {
  return "fragment"
}

