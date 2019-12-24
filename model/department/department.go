package department

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"network/global/pgdb"
	"time"
)

type Department struct {
	tableName    struct{}  `sql:"network_homework.tb_department,discard_unknown_columns"`
	ID           int64     `pg:"id,pk"`
	Name         string    `pg:"name,notnull"`
	Address      string    `pg:"address,notnull"`
	Type         int8      `pg:"type,notnull"`
	Owner        string    `pg:"owner,notnull"`
	OwnerContact string    `pg:"owner_contact,notnull"`
	Admin        string    `pg:"admin,notnull"`
	AdminContact string    `pg:"admin_contact,notnull"`
	CreatedAt    time.Time `pg:"created_at,notnull"`
	ModifiedAt   time.Time `pg:"modified_at,notnull"`
	DeletedAt    time.Time `pg:"deleted_at,soft_delete"`
}

func New() *Department {
	return &Department{}
}

func (d *Department) Model() *orm.Query {
	return pgdb.DB().Model(d)
}

func (d *Department) Insert() error {
	_, err := d.Model().Returning("*").Insert()

	return err
}

func (d *Department) Delete() error {
	_, err := pgdb.DB().Model(d).WherePK().Delete()

	return err
}

func (d *Department) Update() error {
	_, err := pgdb.DB().Model(d).WherePK().UpdateNotNull()

	return err
}

func List(offset int, limit int, t ...int8) ([]Department, int, error) {
	departs := make([]Department, 0)

	query := pgdb.DB().Model(&departs)
	if len(t) > 0 && t[0] > 0 {
		query.Where("department.type IN (?)", pg.In(t))
	}

	count, err := query.Offset(offset).Limit(limit).Order("id asc").SelectAndCount()

	return departs, count, err
}

func (d *Department) Info() (Department, error) {
	depart := Department{}

	err := pgdb.DB().Model(d).WherePK().Select(&depart)

	return depart, err
}

// TODO 如果ID数量过大可能会有问题
func MapID2Department(id ...int64) (map[int64]Department, error) {
	departs := make([]Department, 0)
	mapID2Department := make(map[int64]Department)

	query := pgdb.DB().Model(&departs)
	if len(id) > 0 {
		query.Where("id in ?", pg.Array(id))
	}
	err := query.Select()
	if err != nil {
		return nil, err
	}

	for _, depart := range departs {
		mapID2Department[depart.ID] = depart
	}

	return mapID2Department, nil
}

func (d *Department) IsRoleOr(role ...int8) (bool, error) {
	if len(role) == 0 {
		return false, nil
	}

	count, err := d.Model().Where("type IN (?)", pg.In(role)).WherePK().Count()
	if err != nil {
		return false, err
	}
	if 0 == count {
		return false, nil
	}

	return true, nil
}
