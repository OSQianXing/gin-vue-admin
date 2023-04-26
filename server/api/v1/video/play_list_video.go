package video

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/video"
	videoReq "github.com/flipped-aurora/gin-vue-admin/server/model/video/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PlayListVideoApi struct {
}

var playListVideoService = service.ServiceGroupApp.VideoServiceGroup.PlayListVideoService

// CreatePlayListVideo 创建PlayListVideo
// @Tags PlayListVideo
// @Summary 创建PlayListVideo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body video.PlayListVideo true "创建PlayListVideo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /playListVideo/createPlayListVideo [post]
func (playListVideoApi *PlayListVideoApi) CreatePlayListVideo(c *gin.Context) {
	var playListVideo video.PlayListVideo
	err := c.ShouldBindJSON(&playListVideo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	playListVideo.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"UserId": {utils.NotEmpty()},
	}
	if err := utils.Verify(playListVideo, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := playListVideoService.CreatePlayListVideo(&playListVideo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePlayListVideo 删除PlayListVideo
// @Tags PlayListVideo
// @Summary 删除PlayListVideo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body video.PlayListVideo true "删除PlayListVideo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /playListVideo/deletePlayListVideo [delete]
func (playListVideoApi *PlayListVideoApi) DeletePlayListVideo(c *gin.Context) {
	var playListVideo video.PlayListVideo
	err := c.ShouldBindJSON(&playListVideo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	playListVideo.DeletedBy = utils.GetUserID(c)
	if err := playListVideoService.DeletePlayListVideo(playListVideo); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePlayListVideoByIds 批量删除PlayListVideo
// @Tags PlayListVideo
// @Summary 批量删除PlayListVideo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PlayListVideo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /playListVideo/deletePlayListVideoByIds [delete]
func (playListVideoApi *PlayListVideoApi) DeletePlayListVideoByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := playListVideoService.DeletePlayListVideoByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePlayListVideo 更新PlayListVideo
// @Tags PlayListVideo
// @Summary 更新PlayListVideo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body video.PlayListVideo true "更新PlayListVideo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /playListVideo/updatePlayListVideo [put]
func (playListVideoApi *PlayListVideoApi) UpdatePlayListVideo(c *gin.Context) {
	var playListVideo video.PlayListVideo
	err := c.ShouldBindJSON(&playListVideo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	playListVideo.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"UserId": {utils.NotEmpty()},
	}
	if err := utils.Verify(playListVideo, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := playListVideoService.UpdatePlayListVideo(playListVideo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPlayListVideo 用id查询PlayListVideo
// @Tags PlayListVideo
// @Summary 用id查询PlayListVideo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query video.PlayListVideo true "用id查询PlayListVideo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /playListVideo/findPlayListVideo [get]
func (playListVideoApi *PlayListVideoApi) FindPlayListVideo(c *gin.Context) {
	var playListVideo video.PlayListVideo
	err := c.ShouldBindQuery(&playListVideo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if replayListVideo, err := playListVideoService.GetPlayListVideo(playListVideo.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"replayListVideo": replayListVideo}, c)
	}
}

// GetPlayListVideoList 分页获取PlayListVideo列表
// @Tags PlayListVideo
// @Summary 分页获取PlayListVideo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query videoReq.PlayListVideoSearch true "分页获取PlayListVideo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /playListVideo/getPlayListVideoList [get]
func (playListVideoApi *PlayListVideoApi) GetPlayListVideoList(c *gin.Context) {
	var pageInfo videoReq.PlayListVideoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := playListVideoService.GetPlayListVideoInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
