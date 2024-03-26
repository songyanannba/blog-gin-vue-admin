package blog

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type B_testRouter struct {
}

// InitB_testRouter 初始化 测试 路由信息
func (s *B_testRouter) InitB_testRouter(Router *gin.RouterGroup) {
	b_testRouter := Router.Group("b_test").Use(middleware.OperationRecord())
	b_testRouterWithoutRecord := Router.Group("b_test")
	var b_testApi = v1.ApiGroupApp.BlogApiGroup.B_testApi
	{
		b_testRouter.POST("createB_test", b_testApi.CreateB_test)             // 新建测试
		b_testRouter.DELETE("deleteB_test", b_testApi.DeleteB_test)           // 删除测试
		b_testRouter.DELETE("deleteB_testByIds", b_testApi.DeleteB_testByIds) // 批量删除测试
		b_testRouter.PUT("updateB_test", b_testApi.UpdateB_test)              // 更新测试
	}
	{
		b_testRouterWithoutRecord.GET("findB_test", b_testApi.FindB_test)       // 根据ID获取测试
		b_testRouterWithoutRecord.GET("getB_testList", b_testApi.GetB_testList) // 获取测试列表
	}
}
