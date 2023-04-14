package po

import "time"

type BasePO struct {
	// 主键ID
	Id uint `gorm:"primarykey;column:id"`
	// 创建时间
	CreateDate time.Time `gorm:"column:create_date"`
	// 删除标记
	Deleted bool `gorm:"column:deleted"`
}
