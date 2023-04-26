package video

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PlayListVideoRouter struct {
}

// InitPlayListVideoRouter 初始化 PlayListVideo 路由信息
func (s *PlayListVideoRouter) InitPlayListVideoRouter(Router *gin.RouterGroup) {
	playListVideoRouter := Router.Group("playListVideo").Use(middleware.OperationRecord())
	playListVideoRouterWithoutRecord := Router.Group("playListVideo")
	var playListVideoApi = v1.ApiGroupApp.VideoApiGroup.PlayListVideoApi
	{
		playListVideoRouter.POST("createPlayListVideo", playListVideoApi.CreatePlayListVideo)   // 新建PlayListVideo
		playListVideoRouter.DELETE("deletePlayListVideo", playListVideoApi.DeletePlayListVideo) // 删除PlayListVideo
		playListVideoRouter.DELETE("deletePlayListVideoByIds", playListVideoApi.DeletePlayListVideoByIds) // 批量删除PlayListVideo
		playListVideoRouter.PUT("updatePlayListVideo", playListVideoApi.UpdatePlayListVideo)    // 更新PlayListVideo
	}
	{
		playListVideoRouterWithoutRecord.GET("findPlayListVideo", playListVideoApi.FindPlayListVideo)        // 根据ID获取PlayListVideo
		playListVideoRouterWithoutRecord.GET("getPlayListVideoList", playListVideoApi.GetPlayListVideoList)  // 获取PlayListVideo列表
	}
}
