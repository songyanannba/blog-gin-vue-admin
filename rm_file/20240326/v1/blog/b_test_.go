package blog

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/blog"
	blogReq "github.com/flipped-aurora/gin-vue-admin/server/model/blog/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type B_testApi struct {
}

var b_testService = service.ServiceGroupApp.BlogServiceGroup.B_testService

// CreateB_test 创建测试
// @Tags B_test
// @Summary 创建测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body blog.B_test true "创建测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /b_test/createB_test [post]
func (b_testApi *B_testApi) CreateB_test(c *gin.Context) {
	var b_test blog.B_test
	err := c.ShouldBindJSON(&b_test)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := b_testService.CreateB_test(&b_test); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteB_test 删除测试
// @Tags B_test
// @Summary 删除测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body blog.B_test true "删除测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /b_test/deleteB_test [delete]
func (b_testApi *B_testApi) DeleteB_test(c *gin.Context) {
	ID := c.Query("ID")
	if err := b_testService.DeleteB_test(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteB_testByIds 批量删除测试
// @Tags B_test
// @Summary 批量删除测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /b_test/deleteB_testByIds [delete]
func (b_testApi *B_testApi) DeleteB_testByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := b_testService.DeleteB_testByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateB_test 更新测试
// @Tags B_test
// @Summary 更新测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body blog.B_test true "更新测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /b_test/updateB_test [put]
func (b_testApi *B_testApi) UpdateB_test(c *gin.Context) {
	var b_test blog.B_test
	err := c.ShouldBindJSON(&b_test)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := b_testService.UpdateB_test(b_test); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindB_test 用id查询测试
// @Tags B_test
// @Summary 用id查询测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query blog.B_test true "用id查询测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /b_test/findB_test [get]
func (b_testApi *B_testApi) FindB_test(c *gin.Context) {
	ID := c.Query("ID")
	if reb_test, err := b_testService.GetB_test(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reb_test": reb_test}, c)
	}
}

// GetB_testList 分页获取测试列表
// @Tags B_test
// @Summary 分页获取测试列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query blogReq.B_testSearch true "分页获取测试列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /b_test/getB_testList [get]
func (b_testApi *B_testApi) GetB_testList(c *gin.Context) {
	var pageInfo blogReq.B_testSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := b_testService.GetB_testInfoList(pageInfo); err != nil {
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
