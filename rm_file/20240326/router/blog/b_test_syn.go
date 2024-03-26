package blog

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BTestSynRouter struct {
}

// InitBTestSynRouter 初始化 测试 路由信息
func (s *BTestSynRouter) InitBTestSynRouter(Router *gin.RouterGroup) {
	bTestSynRouter := Router.Group("bTestSyn").Use(middleware.OperationRecord())
	bTestSynRouterWithoutRecord := Router.Group("bTestSyn")
	var bTestSynApi = v1.ApiGroupApp.BlogApiGroup.BTestSynApi
	{
		bTestSynRouter.POST("createBTestSyn", bTestSynApi.CreateBTestSyn)             // 新建测试
		bTestSynRouter.DELETE("deleteBTestSyn", bTestSynApi.DeleteBTestSyn)           // 删除测试
		bTestSynRouter.DELETE("deleteBTestSynByIds", bTestSynApi.DeleteBTestSynByIds) // 批量删除测试
		bTestSynRouter.PUT("updateBTestSyn", bTestSynApi.UpdateBTestSyn)              // 更新测试
	}
	{
		bTestSynRouterWithoutRecord.GET("findBTestSyn", bTestSynApi.FindBTestSyn)       // 根据ID获取测试
		bTestSynRouterWithoutRecord.GET("getBTestSynList", bTestSynApi.GetBTestSynList) // 获取测试列表
	}
}
