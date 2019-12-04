package loginlog

import (
	"github.com/go-pg/pg/orm"
	"network/global/pgdb"
	"time"
)

type UserLog struct {
	tableName struct{}  `sql:"network_homework.tb_login_log,discard_unknown_columns"`
	ID        int64     `sql:"id,pk"`
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

type LoginLog struct {
	ID          int64     `pg:"id"`
	UserAccount string    `sql:"account"`
	UserName    string    `sql:"name"`
	IP          string    `pg:"ip"`
	CreatedAt   time.Time `pg:"created_at"`
}

func List(offset int, limit int) ([]LoginLog, int, error) {
	loginLog := make([]LoginLog, 0)
	count, err := pgdb.DB().Model((*UserLog)(nil)).
		Column("user_log.id", "tb_user.account", "tb_user.name", "user_log.ip", "user_log.created_at").
		Join("LEFT JOIN network_homework.tb_user ON user_log.user_id = tb_user.id").
		Offset(offset).Limit(limit).Order("id asc").
		SelectAndCount(&loginLog)
	return loginLog, count, err
}
