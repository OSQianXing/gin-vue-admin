package video

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/video"
	videoReq "github.com/flipped-aurora/gin-vue-admin/server/model/video/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FragmentAuditApi struct {
}

var fragmentAuditService = service.ServiceGroupApp.VideoServiceGroup.FragmentAuditService

// CreateFragmentAudit 创建FragmentAudit
// @Tags FragmentAudit
// @Summary 创建FragmentAudit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body video.FragmentAudit true "创建FragmentAudit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fragmentAudit/createFragmentAudit [post]
func (fragmentAuditApi *FragmentAuditApi) CreateFragmentAudit(c *gin.Context) {
	var fragmentAudit video.FragmentAudit
	err := c.ShouldBindJSON(&fragmentAudit)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fragmentAudit.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"FragmentId": {utils.NotEmpty()},
		"Type":       {utils.NotEmpty()},
		"UserId":     {utils.NotEmpty()},
	}
	if err := utils.Verify(fragmentAudit, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fragmentAuditService.CreateFragmentAudit(&fragmentAudit); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// UpdateFragmentAudit 更新FragmentAudit
// @Tags FragmentAudit
// @Summary 更新FragmentAudit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body video.FragmentAudit true "更新FragmentAudit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fragmentAudit/updateFragmentAudit [put]
func (fragmentAuditApi *FragmentAuditApi) UpdateFragmentAudit(c *gin.Context) {
	var fragmentAudit video.FragmentAudit
	err := c.ShouldBindJSON(&fragmentAudit)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fragmentAudit.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"FragmentId": {utils.NotEmpty()},
		"Type":       {utils.NotEmpty()},
		"UserId":     {utils.NotEmpty()},
	}
	if err := utils.Verify(fragmentAudit, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fragmentAuditService.UpdateFragmentAudit(fragmentAudit); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFragmentAudit 用id查询FragmentAudit
// @Tags FragmentAudit
// @Summary 用id查询FragmentAudit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query video.FragmentAudit true "用id查询FragmentAudit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fragmentAudit/findFragmentAudit [get]
func (fragmentAuditApi *FragmentAuditApi) FindFragmentAudit(c *gin.Context) {
	var fragmentAudit video.FragmentAudit
	err := c.ShouldBindQuery(&fragmentAudit)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if refragmentAudit, err := fragmentAuditService.GetFragmentAudit(fragmentAudit.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refragmentAudit": refragmentAudit}, c)
	}
}

// GetFragmentAuditList 分页获取FragmentAudit列表
// @Tags FragmentAudit
// @Summary 分页获取FragmentAudit列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query videoReq.FragmentAuditSearch true "分页获取FragmentAudit列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fragmentAudit/getFragmentAuditList [get]
func (fragmentAuditApi *FragmentAuditApi) GetFragmentAuditList(c *gin.Context) {
	var pageInfo videoReq.FragmentAuditSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := fragmentAuditService.GetFragmentAuditInfoList(pageInfo); err != nil {
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
