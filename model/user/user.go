package user

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"network/global/constant"
	"network/global/pgdb"
	"time"
)

// TODO 这里 pg tag貌似不生效，得用sql tag才行，但是，用sql tag time不会用默认值
type User struct {
	tableName  struct{}  `sql:"network_homework.tb_user,discard_unknown_columns"`
	ID         int64     `pg:"id,pk"`
	Account    string    `pg:"account,notnull"`
	Name       string    `sql:"name,notnull"`
	Type       int8      `pg:"type,notnull"` // 1:市级单位 2:市级各辖区单位 3:受监管企业单位 4:签约技术支持/安全服务单位
	Phone      string    `pg:"phone" binding:"phone"`
	Email      string    `pg:"email" binding:"email"`
	WeChat     string    `sql:"wechat"`
	Department int64     `pg:"department"`
	Valid      bool      `pg:"valid,notnull"`
	CreatedAt  time.Time `pg:"created_at,notnull"`
	ModifiedAt time.Time `pg:"modified_at,notnull"`
	DeletedAt  time.Time `pg:"deleted_at,soft_delete"`
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
	_, err := u.Model().WherePK().Returning("*").UpdateNotNull()

	return err
}

// TODO ugly code
func (u *User) Restore() error {
	_, err := u.Model().ExecOne("UPDATE network_homework.tb_user SET deleted_at = NULL")
	return err
}

func List(offset int, limit int, types ...int8) ([]User, int, error) {
	users := make([]User, 0)
	count := 0
	var err error

	if len(types) > 0 {
		str := `select * from network_homework.tb_user 
				where department  IN (select id from network_homework.tb_department
    			where tb_department.type in ?);`
		_, err = pgdb.DB().Query(&users, str, pg.InMulti(types))
	} else {
		query := pgdb.DB().Model(&users)
		count, err = query.Where("type != ?", constant.TypeUserAdministrator).
			Offset(offset).Limit(limit).Order("id asc").SelectAndCount()
	}

	return users, count, err
}

func (u *User) Login() (bool, error) {
	err := u.Model().Where("account = ? and password = ?", u.Account, u.Password).Select()
	if pg.ErrNoRows == err {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func (u *User) IsAdmin() (bool, error) {
	err := u.Model().Where("account = ? and type = ?", u.Account, constant.TypeUserAdministrator).Select()
	if pg.ErrNoRows == err {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}

// TODO 如果ID数量过大可能会有问题
func MapID2User(id ...int64) (map[int64]User, error) {
	users := make([]User, 0)
	mapID2User := make(map[int64]User)

	query := pgdb.DB().Model(&users)
	if len(id) > 0 {
		query.Where("id in ?", pg.Array(id))
	}
	err := query.Select()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		mapID2User[user.ID] = user
	}

	return mapID2User, nil
}

func OneByAccount(account string) (User, error) {
	u := User{}
	err := pgdb.DB().Model(&u).Where("account = ?", account).Select()
	return u, err
}

func OneWithDeletedByAccount(account string) (User, error) {
	u := User{}
	_, err := pgdb.DB().QueryOne(&u, "SELECT * FROM network_homework.tb_user where account = ?", account)
	return u, err
}

func (u *User) SimDepUser() ([]User, error) {
	us := make([]User, 0)
	err := u.Model().Where("department = ?", u.Department).Select(&us)
	return us, err
}

func GetRoleByAccount(account string) (int8, error) {
	var role int8
	_, err := pgdb.DB().QueryOne(&role, "SELECT type FROM network_homework.tb_department where id = " +
		"(SELECT department FROM network_homework.tb_user where account = ?)", account)

	return role, err
}