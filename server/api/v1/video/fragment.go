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

type FragmentApi struct {
}

var fragmentService = service.ServiceGroupApp.VideoServiceGroup.FragmentService

// CreateFragment 创建Fragment
// @Tags Fragment
// @Summary 创建Fragment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body video.Fragment true "创建Fragment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fragment/createFragment [post]
func (fragmentApi *FragmentApi) CreateFragment(c *gin.Context) {
	var fragment video.Fragment
	err := c.ShouldBindJSON(&fragment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Author":    {utils.NotEmpty()},
		"Name":      {utils.NotEmpty()},
		"Subtitles": {utils.NotEmpty()},
	}
	if err := utils.Verify(fragment, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fragmentService.CreateFragment(&fragment); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFragment 删除Fragment
// @Tags Fragment
// @Summary 删除Fragment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body video.Fragment true "删除Fragment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fragment/deleteFragment [delete]
func (fragmentApi *FragmentApi) DeleteFragment(c *gin.Context) {
	var fragment video.Fragment
	err := c.ShouldBindJSON(&fragment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fragmentService.DeleteFragment(fragment); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFragmentByIds 批量删除Fragment
// @Tags Fragment
// @Summary 批量删除Fragment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Fragment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /fragment/deleteFragmentByIds [delete]
func (fragmentApi *FragmentApi) DeleteFragmentByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fragmentService.DeleteFragmentByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// FindFragment 用id查询Fragment
// @Tags Fragment
// @Summary 用id查询Fragment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query video.Fragment true "用id查询Fragment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fragment/findFragment [get]
func (fragmentApi *FragmentApi) FindFragment(c *gin.Context) {
	var fragment video.Fragment
	err := c.ShouldBindQuery(&fragment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if refragment, err := fragmentService.GetFragment(fragment.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refragment": refragment}, c)
	}
}

// GetFragmentList 分页获取Fragment列表
// @Tags Fragment
// @Summary 分页获取Fragment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query videoReq.FragmentSearch true "分页获取Fragment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fragment/getFragmentList [get]
func (fragmentApi *FragmentApi) GetFragmentList(c *gin.Context) {
	var pageInfo videoReq.FragmentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := fragmentService.GetFragmentInfoList(pageInfo); err != nil {
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
