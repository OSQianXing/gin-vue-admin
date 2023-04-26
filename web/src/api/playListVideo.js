import service from '@/utils/request'

// @Tags PlayListVideo
// @Summary 创建PlayListVideo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PlayListVideo true "创建PlayListVideo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /playListVideo/createPlayListVideo [post]
export const createPlayListVideo = (data) => {
  return service({
    url: '/playListVideo/createPlayListVideo',
    method: 'post',
    data
  })
}

// @Tags PlayListVideo
// @Summary 删除PlayListVideo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PlayListVideo true "删除PlayListVideo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /playListVideo/deletePlayListVideo [delete]
export const deletePlayListVideo = (data) => {
  return service({
    url: '/playListVideo/deletePlayListVideo',
    method: 'delete',
    data
  })
}

// @Tags PlayListVideo
// @Summary 删除PlayListVideo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PlayListVideo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /playListVideo/deletePlayListVideo [delete]
export const deletePlayListVideoByIds = (data) => {
  return service({
    url: '/playListVideo/deletePlayListVideoByIds',
    method: 'delete',
    data
  })
}

// @Tags PlayListVideo
// @Summary 更新PlayListVideo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PlayListVideo true "更新PlayListVideo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /playListVideo/updatePlayListVideo [put]
export const updatePlayListVideo = (data) => {
  return service({
    url: '/playListVideo/updatePlayListVideo',
    method: 'put',
    data
  })
}

// @Tags PlayListVideo
// @Summary 用id查询PlayListVideo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PlayListVideo true "用id查询PlayListVideo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /playListVideo/findPlayListVideo [get]
export const findPlayListVideo = (params) => {
  return service({
    url: '/playListVideo/findPlayListVideo',
    method: 'get',
    params
  })
}

// @Tags PlayListVideo
// @Summary 分页获取PlayListVideo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取PlayListVideo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /playListVideo/getPlayListVideoList [get]
export const getPlayListVideoList = (params) => {
  return service({
    url: '/playListVideo/getPlayListVideoList',
    method: 'get',
    params
  })
}
