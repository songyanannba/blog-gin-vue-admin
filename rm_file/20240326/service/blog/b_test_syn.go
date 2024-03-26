package blog

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/blog"
	blogReq "github.com/flipped-aurora/gin-vue-admin/server/model/blog/request"
	"gorm.io/gorm"
)

type BTestSynService struct {
}

// CreateBTestSyn 创建测试记录
// Author [piexlmax](https://github.com/piexlmax)
func (bTestSynService *BTestSynService) CreateBTestSyn(bTestSyn *blog.BTestSyn) (err error) {
	err = global.GVA_DB.Create(bTestSyn).Error
	return err
}

// DeleteBTestSyn 删除测试记录
// Author [piexlmax](https://github.com/piexlmax)
func (bTestSynService *BTestSynService) DeleteBTestSyn(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&blog.BTestSyn{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&blog.BTestSyn{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteBTestSynByIds 批量删除测试记录
// Author [piexlmax](https://github.com/piexlmax)
func (bTestSynService *BTestSynService) DeleteBTestSynByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&blog.BTestSyn{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&blog.BTestSyn{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateBTestSyn 更新测试记录
// Author [piexlmax](https://github.com/piexlmax)
func (bTestSynService *BTestSynService) UpdateBTestSyn(bTestSyn blog.BTestSyn) (err error) {
	err = global.GVA_DB.Save(&bTestSyn).Error
	return err
}

// GetBTestSyn 根据ID获取测试记录
// Author [piexlmax](https://github.com/piexlmax)
func (bTestSynService *BTestSynService) GetBTestSyn(ID string) (bTestSyn blog.BTestSyn, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&bTestSyn).Error
	return
}

// GetBTestSynInfoList 分页获取测试记录
// Author [piexlmax](https://github.com/piexlmax)
func (bTestSynService *BTestSynService) GetBTestSynInfoList(info blogReq.BTestSynSearch) (list []blog.BTestSyn, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&blog.BTestSyn{})
	var bTestSyns []blog.BTestSyn
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

	err = db.Find(&bTestSyns).Error
	return bTestSyns, total, err
}
