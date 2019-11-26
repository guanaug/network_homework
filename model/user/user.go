package user

import (
	"database/sql"
	"github.com/go-pg/pg/orm"
	"network/global/constant"
	"network/global/pgdb"
	"time"
)

type User struct {
	tableName  struct{}  `sql:"tb_user"`
	ID         int64     `pg:"id, pk"`
	Account    string    `pg:"account, notnull"`
	Name       string    `pg:"name, notnull"`
	Type       int8      `pg:"type, notnull"` // 1:市级单位 2:市级各辖区单位 3:受监管企业单位 4:签约技术支持/安全服务单位
	Phone      string    `pg:"phone" binding:"phone"`
	Email      string    `pg:"email" binding:"email"`
	WeChat     string    `pg:"wechat"`
	Department int64     `pg:"department"`
	Valid      bool      `pg:"valid, notnull"`
	CreatedAt  time.Time `pg:"created_at, notnull"`
	ModifiedAt time.Time `pg:"modified_at, notnull"`
	DeletedAt  time.Time `pg:"deleted_at, soft_delete"`
	Password   string    `pg:"password"`
}

func New() *User {
	return &User{}
}

func (u *User) Model() *orm.Query {
	return pgdb.DB().Model(u)
}

func (u *User) Insert() error {
	_, err := u.Model().Returning("*").Insert()

	return err
}

func (u *User) Delete() error {
	_, err := u.Model().WherePK().Delete()

	return err
}

func (u *User) Update() error {
	_, err := u.Model().WherePK().Update()

	return err
}

func List(offset int, limit int) ([]User, int, error) {
	users := make([]User, 0)
	count, err := pgdb.DB().Model(&users).Offset(offset).Limit(limit).Order("id asc").SelectAndCount()

	return users, count, err
}

func (u *User) Login() (bool, error) {
	err := u.Model().Where("account = ? and password = ?", u.Account, u.Password).Select()
	if sql.ErrNoRows == err {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func (u *User) IsAdmin() (bool, error) {
	err := u.Model().Where("account = ? and type = ?", u.Account, constant.TypeUserAdministrator).Select()
	if sql.ErrNoRows == err {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}
