package loginlog

import (
	"github.com/go-pg/pg/orm"
	"network/global/pgdb"
	"time"
)

type UserLog struct {
	tableName struct{}  `sql:"network_homework.tb_login_log"`
	ID        int64     `sql:"id, pk"`
	UserID    int64     `sql:"user_id, notnull"`
	IP        string    `sql:"ip"`
	CreatedAt time.Time `sql:"created_at"`
}

func New() *UserLog {
	return &UserLog{}
}

func (ul *UserLog) Model() *orm.Query {
	return pgdb.DB().Model(ul)
}

func (ul *UserLog) Insert() error {
	_, err := ul.Model().Returning("*").Insert()

	return err
}
