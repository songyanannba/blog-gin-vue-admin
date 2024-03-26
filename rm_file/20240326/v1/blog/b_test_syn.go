package blog

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/blog"
	blogReq "github.com/flipped-aurora/gin-vue-admin/server/model/blog/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BTestSynApi struct {
}

var bTestSynService = service.ServiceGroupApp.BlogServiceGroup.BTestSynService

// CreateBTestSyn 创建测试
// @Tags BTestSyn
// @Summary 创建测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body blog.BTestSyn true "创建测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /bTestSyn/createBTestSyn [post]
func (bTestSynApi *BTestSynApi) CreateBTestSyn(c *gin.Context) {
	var bTestSyn blog.BTestSyn
	err := c.ShouldBindJSON(&bTestSyn)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	bTestSyn.CreatedBy = utils.GetUserID(c)

	if err := bTestSynService.CreateBTestSyn(&bTestSyn); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBTestSyn 删除测试
// @Tags BTestSyn
// @Summary 删除测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body blog.BTestSyn true "删除测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bTestSyn/deleteBTestSyn [delete]
func (bTestSynApi *BTestSynApi) DeleteBTestSyn(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := bTestSynService.DeleteBTestSyn(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBTestSynByIds 批量删除测试
// @Tags BTestSyn
// @Summary 批量删除测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /bTestSyn/deleteBTestSynByIds [delete]
func (bTestSynApi *BTestSynApi) DeleteBTestSynByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := bTestSynService.DeleteBTestSynByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBTestSyn 更新测试
// @Tags BTestSyn
// @Summary 更新测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body blog.BTestSyn true "更新测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bTestSyn/updateBTestSyn [put]
func (bTestSynApi *BTestSynApi) UpdateBTestSyn(c *gin.Context) {
	var bTestSyn blog.BTestSyn
	err := c.ShouldBindJSON(&bTestSyn)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	bTestSyn.UpdatedBy = utils.GetUserID(c)

	if err := bTestSynService.UpdateBTestSyn(bTestSyn); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBTestSyn 用id查询测试
// @Tags BTestSyn
// @Summary 用id查询测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query blog.BTestSyn true "用id查询测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bTestSyn/findBTestSyn [get]
func (bTestSynApi *BTestSynApi) FindBTestSyn(c *gin.Context) {
	ID := c.Query("ID")
	if rebTestSyn, err := bTestSynService.GetBTestSyn(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebTestSyn": rebTestSyn}, c)
	}
}

// GetBTestSynList 分页获取测试列表
// @Tags BTestSyn
// @Summary 分页获取测试列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query blogReq.BTestSynSearch true "分页获取测试列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bTestSyn/getBTestSynList [get]
func (bTestSynApi *BTestSynApi) GetBTestSynList(c *gin.Context) {
	var pageInfo blogReq.BTestSynSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := bTestSynService.GetBTestSynInfoList(pageInfo); err != nil {
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
