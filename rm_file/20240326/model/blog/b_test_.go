// 自动生成模板B_test
package blog

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 测试 结构体  B_test
type B_test struct {
	global.GVA_MODEL
	Id   *int   `json:"id" form:"id" gorm:"primarykey;column:主键;comment:;" binding:"required"` //主键
	Name string `json:"name" form:"name" gorm:"column:名称;comment:;"`                           //名称
}

// TableName 测试 B_test自定义表名 b_test
func (B_test) TableName() string {
	return "b_test"
}
