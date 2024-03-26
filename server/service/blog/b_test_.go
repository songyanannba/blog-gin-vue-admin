package blog

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/blog"
	blogReq "github.com/flipped-aurora/gin-vue-admin/server/model/blog/request"
	"gorm.io/gorm"
)

type BTestService struct {
}

// CreateBTest 创建测试记录
// Author [piexlmax](https://github.com/piexlmax)
func (bTestService *BTestService) CreateBTest(bTest *blog.BTest) (err error) {
	err = global.GVA_DB.Create(bTest).Error
	return err
}

// DeleteBTest 删除测试记录
// Author [piexlmax](https://github.com/piexlmax)
func (bTestService *BTestService) DeleteBTest(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&blog.BTest{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&blog.BTest{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteBTestByIds 批量删除测试记录
// Author [piexlmax](https://github.com/piexlmax)
func (bTestService *BTestService) DeleteBTestByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&blog.BTest{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&blog.BTest{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateBTest 更新测试记录
// Author [piexlmax](https://github.com/piexlmax)
func (bTestService *BTestService) UpdateBTest(bTest blog.BTest) (err error) {
	err = global.GVA_DB.Save(&bTest).Error
	return err
}

// GetBTest 根据ID获取测试记录
// Author [piexlmax](https://github.com/piexlmax)
func (bTestService *BTestService) GetBTest(ID string) (bTest blog.BTest, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&bTest).Error
	return
}

// GetBTestInfoList 分页获取测试记录
// Author [piexlmax](https://github.com/piexlmax)
func (bTestService *BTestService) GetBTestInfoList(info blogReq.BTestSearch) (list []blog.BTest, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&blog.BTest{})
	var bTests []blog.BTest
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

	err = db.Find(&bTests).Error
	return bTests, total, err
}
