import service from '@/utils/request'

// @Tags BTestSyn
// @Summary 创建测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BTestSyn true "创建测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /bTestSyn/createBTestSyn [post]
export const createBTestSyn = (data) => {
  return service({
    url: '/bTestSyn/createBTestSyn',
    method: 'post',
    data
  })
}

// @Tags BTestSyn
// @Summary 删除测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BTestSyn true "删除测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bTestSyn/deleteBTestSyn [delete]
export const deleteBTestSyn = (params) => {
  return service({
    url: '/bTestSyn/deleteBTestSyn',
    method: 'delete',
    params
  })
}

// @Tags BTestSyn
// @Summary 批量删除测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bTestSyn/deleteBTestSyn [delete]
export const deleteBTestSynByIds = (params) => {
  return service({
    url: '/bTestSyn/deleteBTestSynByIds',
    method: 'delete',
    params
  })
}

// @Tags BTestSyn
// @Summary 更新测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BTestSyn true "更新测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bTestSyn/updateBTestSyn [put]
export const updateBTestSyn = (data) => {
  return service({
    url: '/bTestSyn/updateBTestSyn',
    method: 'put',
    data
  })
}

// @Tags BTestSyn
// @Summary 用id查询测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BTestSyn true "用id查询测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bTestSyn/findBTestSyn [get]
export const findBTestSyn = (params) => {
  return service({
    url: '/bTestSyn/findBTestSyn',
    method: 'get',
    params
  })
}

// @Tags BTestSyn
// @Summary 分页获取测试列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取测试列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bTestSyn/getBTestSynList [get]
export const getBTestSynList = (params) => {
  return service({
    url: '/bTestSyn/getBTestSynList',
    method: 'get',
    params
  })
}
