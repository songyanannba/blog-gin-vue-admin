// 自动生成模板BTestSyn
package blog

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 测试 结构体  BTestSyn
type BTestSyn struct {
	global.GVA_MODEL
	Id        *int   `json:"id" form:"id" gorm:"primarykey;column:id;comment:;" binding:"required"` //x
	Name      string `json:"name" form:"name" gorm:"column:name;comment:;"`                         //name
	Str       string `json:"str" form:"str" gorm:"column:str;comment:;"`                            //str
	CreatedBy uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 测试 BTestSyn自定义表名 b_test_syn
func (BTestSyn) TableName() string {
	return "b_test_syn"
}
