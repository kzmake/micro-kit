package mysql

import (
	"time"

	// use gorm and mysql
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// BaseSchema はベースとなるスキーマです。
type BaseSchema struct {
	CreatedAt time.Time  `gorm:"type:datetime(6)"`
	UpdatedAt time.Time  `gorm:"type:datetime(6)"`
	DeletedAt *time.Time `gorm:"type:datetime(6);index"`
}
