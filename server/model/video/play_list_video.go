// 自动生成模板PlayListVideo
package video

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// PlayListVideo 结构体
type PlayListVideo struct {
      global.GVA_MODEL
      AuditReason  string `json:"auditReason" form:"auditReason" gorm:"column:audit_reason;comment:审核不通过原因;size:2000;"`
      AuditedAt  *time.Time `json:"auditedAt" form:"auditedAt" gorm:"column:audited_at;comment:;"`
      PlayListId  *int `json:"playListId" form:"playListId" gorm:"column:play_list_id;comment:播单ID;size:20;"`
      ProdAuditReason  string `json:"prodAuditReason" form:"prodAuditReason" gorm:"column:prod_audit_reason;comment:成品审核不通过原因;size:2000;"`
      ProdAuditedAt  *time.Time `json:"prodAuditedAt" form:"prodAuditedAt" gorm:"column:prod_audited_at;comment:;"`
      ProdId  *int `json:"prodId" form:"prodId" gorm:"column:prod_id;comment:成品ID;size:10;"`
      ProdStatus  *bool `json:"prodStatus" form:"prodStatus" gorm:"column:prod_status;comment:成品状态 0-待审核 1-审核通过 2-审核不通过;"`
      ProdUserId  *int `json:"prodUserId" form:"prodUserId" gorm:"column:prod_user_id;comment:;size:20;"`
      Status  *bool `json:"status" form:"status" gorm:"column:status;comment:短视频状态 0-待审核 1-审核通过 2-审核不通过;"`
      Url  string `json:"url" form:"url" gorm:"column:url;comment:合成视频地址;size:2000;"`
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:审核人ID;size:20;"`
      Version  string `json:"version" form:"version" gorm:"column:version;comment:版本;size:16;"`
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName PlayListVideo 表名
func (PlayListVideo) TableName() string {
  return "play_list_video"
}

