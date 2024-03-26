import service from '@/utils/request'

// @Tags B_test
// @Summary 创建测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.B_test true "创建测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /b_test/createB_test [post]
export const createB_test = (data) => {
  return service({
    url: '/b_test/createB_test',
    method: 'post',
    data
  })
}

// @Tags B_test
// @Summary 删除测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.B_test true "删除测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /b_test/deleteB_test [delete]
export const deleteB_test = (params) => {
  return service({
    url: '/b_test/deleteB_test',
    method: 'delete',
    params
  })
}

// @Tags B_test
// @Summary 批量删除测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /b_test/deleteB_test [delete]
export const deleteB_testByIds = (params) => {
  return service({
    url: '/b_test/deleteB_testByIds',
    method: 'delete',
    params
  })
}

// @Tags B_test
// @Summary 更新测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.B_test true "更新测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /b_test/updateB_test [put]
export const updateB_test = (data) => {
  return service({
    url: '/b_test/updateB_test',
    method: 'put',
    data
  })
}

// @Tags B_test
// @Summary 用id查询测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.B_test true "用id查询测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /b_test/findB_test [get]
export const findB_test = (params) => {
  return service({
    url: '/b_test/findB_test',
    method: 'get',
    params
  })
}

// @Tags B_test
// @Summary 分页获取测试列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取测试列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /b_test/getB_testList [get]
export const getB_testList = (params) => {
  return service({
    url: '/b_test/getB_testList',
    method: 'get',
    params
  })
}
