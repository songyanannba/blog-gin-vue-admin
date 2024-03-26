package blog

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/blog"
	blogReq "github.com/flipped-aurora/gin-vue-admin/server/model/blog/request"
)

type B_testService struct {
}

// CreateB_test 创建测试记录
// Author [piexlmax](https://github.com/piexlmax)
func (b_testService *B_testService) CreateB_test(b_test *blog.B_test) (err error) {
	err = global.GVA_DB.Create(b_test).Error
	return err
}

// DeleteB_test 删除测试记录
// Author [piexlmax](https://github.com/piexlmax)
func (b_testService *B_testService) DeleteB_test(ID string) (err error) {
	err = global.GVA_DB.Delete(&blog.B_test{}, "id = ?", ID).Error
	return err
}

// DeleteB_testByIds 批量删除测试记录
// Author [piexlmax](https://github.com/piexlmax)
func (b_testService *B_testService) DeleteB_testByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]blog.B_test{}, "id in ?", IDs).Error
	return err
}

// UpdateB_test 更新测试记录
// Author [piexlmax](https://github.com/piexlmax)
func (b_testService *B_testService) UpdateB_test(b_test blog.B_test) (err error) {
	err = global.GVA_DB.Save(&b_test).Error
	return err
}

// GetB_test 根据ID获取测试记录
// Author [piexlmax](https://github.com/piexlmax)
func (b_testService *B_testService) GetB_test(ID string) (b_test blog.B_test, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&b_test).Error
	return
}

// GetB_testInfoList 分页获取测试记录
// Author [piexlmax](https://github.com/piexlmax)
func (b_testService *B_testService) GetB_testInfoList(info blogReq.B_testSearch) (list []blog.B_test, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&blog.B_test{})
	var b_tests []blog.B_test
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&b_tests).Error
	return b_tests, total, err
}
