package blog

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BTestRouter struct {
}

// InitBTestRouter 初始化 测试 路由信息
func (s *BTestRouter) InitBTestRouter(Router *gin.RouterGroup) {
	bTestRouter := Router.Group("bTest").Use(middleware.OperationRecord())
	bTestRouterWithoutRecord := Router.Group("bTest")
	var bTestApi = v1.ApiGroupApp.BlogApiGroup.BTestApi
	{
		bTestRouter.POST("createBTest", bTestApi.CreateBTest)             // 新建测试
		bTestRouter.DELETE("deleteBTest", bTestApi.DeleteBTest)           // 删除测试
		bTestRouter.DELETE("deleteBTestByIds", bTestApi.DeleteBTestByIds) // 批量删除测试
		bTestRouter.PUT("updateBTest", bTestApi.UpdateBTest)              // 更新测试
	}
	{
		bTestRouterWithoutRecord.GET("findBTest", bTestApi.FindBTest)       // 根据ID获取测试
		bTestRouterWithoutRecord.GET("getBTestList", bTestApi.GetBTestList) // 获取测试列表
	}
}
