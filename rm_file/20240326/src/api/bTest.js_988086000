import service from '@/utils/request'

// @Tags BTest
// @Summary 创建测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BTest true "创建测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /bTest/createBTest [post]
export const createBTest = (data) => {
  return service({
    url: '/bTest/createBTest',
    method: 'post',
    data
  })
}

// @Tags BTest
// @Summary 删除测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BTest true "删除测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bTest/deleteBTest [delete]
export const deleteBTest = (params) => {
  return service({
    url: '/bTest/deleteBTest',
    method: 'delete',
    params
  })
}

// @Tags BTest
// @Summary 批量删除测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bTest/deleteBTest [delete]
export const deleteBTestByIds = (params) => {
  return service({
    url: '/bTest/deleteBTestByIds',
    method: 'delete',
    params
  })
}

// @Tags BTest
// @Summary 更新测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BTest true "更新测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bTest/updateBTest [put]
export const updateBTest = (data) => {
  return service({
    url: '/bTest/updateBTest',
    method: 'put',
    data
  })
}

// @Tags BTest
// @Summary 用id查询测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BTest true "用id查询测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bTest/findBTest [get]
export const findBTest = (params) => {
  return service({
    url: '/bTest/findBTest',
    method: 'get',
    params
  })
}

// @Tags BTest
// @Summary 分页获取测试列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取测试列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bTest/getBTestList [get]
export const getBTestList = (params) => {
  return service({
    url: '/bTest/getBTestList',
    method: 'get',
    params
  })
}
