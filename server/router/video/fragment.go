package video

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FragmentRouter struct {
}

// InitFragmentRouter 初始化 Fragment 路由信息
func (s *FragmentRouter) InitFragmentRouter(Router *gin.RouterGroup) {
	fragmentRouter := Router.Group("fragment").Use(middleware.OperationRecord())
	fragmentRouterWithoutRecord := Router.Group("fragment")
	var fragmentApi = v1.ApiGroupApp.VideoApiGroup.FragmentApi
	{
		fragmentRouter.POST("createFragment", fragmentApi.CreateFragment)             // 新建Fragment
		fragmentRouter.DELETE("deleteFragment", fragmentApi.DeleteFragment)           // 删除Fragment
		fragmentRouter.DELETE("deleteFragmentByIds", fragmentApi.DeleteFragmentByIds) // 批量删除Fragment
	}
	{
		fragmentRouterWithoutRecord.GET("findFragment", fragmentApi.FindFragment)       // 根据ID获取Fragment
		fragmentRouterWithoutRecord.GET("getFragmentList", fragmentApi.GetFragmentList) // 获取Fragment列表
	}
}
