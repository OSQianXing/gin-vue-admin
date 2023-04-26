package video

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FragmentAuditRouter struct {
}

// InitFragmentAuditRouter 初始化 FragmentAudit 路由信息
func (s *FragmentAuditRouter) InitFragmentAuditRouter(Router *gin.RouterGroup) {
	fragmentAuditRouter := Router.Group("fragmentAudit").Use(middleware.OperationRecord())
	fragmentAuditRouterWithoutRecord := Router.Group("fragmentAudit")
	var fragmentAuditApi = v1.ApiGroupApp.VideoApiGroup.FragmentAuditApi
	{
		fragmentAuditRouter.POST("createFragmentAudit", fragmentAuditApi.CreateFragmentAudit) // 新建FragmentAudit
		fragmentAuditRouter.PUT("updateFragmentAudit", fragmentAuditApi.UpdateFragmentAudit)  // 更新FragmentAudit
	}
	{
		fragmentAuditRouterWithoutRecord.GET("findFragmentAudit", fragmentAuditApi.FindFragmentAudit)       // 根据ID获取FragmentAudit
		fragmentAuditRouterWithoutRecord.GET("getFragmentAuditList", fragmentAuditApi.GetFragmentAuditList) // 获取FragmentAudit列表
	}
}
