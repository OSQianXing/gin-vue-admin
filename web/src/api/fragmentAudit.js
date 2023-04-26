import service from '@/utils/request'

// @Tags FragmentAudit
// @Summary 创建FragmentAudit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FragmentAudit true "创建FragmentAudit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fragmentAudit/createFragmentAudit [post]
export const createFragmentAudit = (data) => {
  return service({
    url: '/fragmentAudit/createFragmentAudit',
    method: 'post',
    data
  })
}

// @Tags FragmentAudit
// @Summary 删除FragmentAudit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FragmentAudit true "删除FragmentAudit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fragmentAudit/deleteFragmentAudit [delete]
export const deleteFragmentAudit = (data) => {
  return service({
    url: '/fragmentAudit/deleteFragmentAudit',
    method: 'delete',
    data
  })
}

// @Tags FragmentAudit
// @Summary 删除FragmentAudit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FragmentAudit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fragmentAudit/deleteFragmentAudit [delete]
export const deleteFragmentAuditByIds = (data) => {
  return service({
    url: '/fragmentAudit/deleteFragmentAuditByIds',
    method: 'delete',
    data
  })
}

// @Tags FragmentAudit
// @Summary 更新FragmentAudit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FragmentAudit true "更新FragmentAudit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fragmentAudit/updateFragmentAudit [put]
export const updateFragmentAudit = (data) => {
  return service({
    url: '/fragmentAudit/updateFragmentAudit',
    method: 'put',
    data
  })
}

// @Tags FragmentAudit
// @Summary 用id查询FragmentAudit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.FragmentAudit true "用id查询FragmentAudit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fragmentAudit/findFragmentAudit [get]
export const findFragmentAudit = (params) => {
  return service({
    url: '/fragmentAudit/findFragmentAudit',
    method: 'get',
    params
  })
}

// @Tags FragmentAudit
// @Summary 分页获取FragmentAudit列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取FragmentAudit列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fragmentAudit/getFragmentAuditList [get]
export const getFragmentAuditList = (params) => {
  return service({
    url: '/fragmentAudit/getFragmentAuditList',
    method: 'get',
    params
  })
}
