// 自动生成模板FragmentAudit
package video

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// FragmentAudit 结构体
type FragmentAudit struct {
      global.GVA_MODEL
      ClippingRange  string `json:"clippingRange" form:"clippingRange" gorm:"column:clipping_range;comment:剪辑范围（掐头去尾）;size:256;"`
      FragmentId  *int `json:"fragmentId" form:"fragmentId" gorm:"column:fragment_id;comment:碎片视频ID;size:19;"`
      Reason  string `json:"reason" form:"reason" gorm:"column:reason;comment:审核失败原因;size:2000;"`
      Status  *bool `json:"status" form:"status" gorm:"column:status;comment:审核结果  0 未审核 1 不通过 2 neededit 3 pending  4 ok  5 great;"`
      Type  *bool `json:"type" form:"type" gorm:"column:type;comment:审核类型  0 初审一审  1 初审二审 2 复审;"`
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:审核人 ID;size:19;"`
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName FragmentAudit 表名
func (FragmentAudit) TableName() string {
  return "fragment_audit"
}

