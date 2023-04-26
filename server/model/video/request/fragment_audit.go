package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/video"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type FragmentAuditSearch struct{
    video.FragmentAudit
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
