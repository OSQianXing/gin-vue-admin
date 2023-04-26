import service from '@/utils/request'

// @Tags Fragment
// @Summary 创建Fragment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Fragment true "创建Fragment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fragment/createFragment [post]
export const createFragment = (data) => {
  return service({
    url: '/fragment/createFragment',
    method: 'post',
    data
  })
}

// @Tags Fragment
// @Summary 删除Fragment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Fragment true "删除Fragment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fragment/deleteFragment [delete]
export const deleteFragment = (data) => {
  return service({
    url: '/fragment/deleteFragment',
    method: 'delete',
    data
  })
}

// @Tags Fragment
// @Summary 删除Fragment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Fragment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fragment/deleteFragment [delete]
export const deleteFragmentByIds = (data) => {
  return service({
    url: '/fragment/deleteFragmentByIds',
    method: 'delete',
    data
  })
}

// @Tags Fragment
// @Summary 更新Fragment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Fragment true "更新Fragment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fragment/updateFragment [put]
export const updateFragment = (data) => {
  return service({
    url: '/fragment/updateFragment',
    method: 'put',
    data
  })
}

// @Tags Fragment
// @Summary 用id查询Fragment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Fragment true "用id查询Fragment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fragment/findFragment [get]
export const findFragment = (params) => {
  return service({
    url: '/fragment/findFragment',
    method: 'get',
    params
  })
}

// @Tags Fragment
// @Summary 分页获取Fragment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Fragment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fragment/getFragmentList [get]
export const getFragmentList = (params) => {
  return service({
    url: '/fragment/getFragmentList',
    method: 'get',
    params
  })
}
