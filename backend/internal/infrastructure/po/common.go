package po

import (
	"context"

	"gorm.io/gorm"

	"github.com/wanyuqin/etcd-ui/backend/db"
)

var (
	DescCreatedAt = "created_at desc"
	DescId        = "id desc"
	WhereTrue     = "1=1"
)

type BaseDao struct {
	DB *gorm.DB
}

func NewBaseDao(ctx context.Context, tn string) BaseDao {
	return BaseDao{
		DB: db.Mysql.Table(tn).WithContext(ctx),
	}
}
