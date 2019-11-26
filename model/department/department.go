package department

import (
	"network/global/pgdb"
	"time"
)

type Department struct {
	tableName    struct{}  `sql:"network_homework.tb_department, discard_unknown_columns"`
	ID           int64     `pg:"id, pk"`
	Name         string    `pg:"name, notnull"`
	Address      string    `pg:"address, notnull"`
	Type         int8      `pg:"type, notnull"`
	Owner        string    `pg:"owner, notnull"`
	OwnerContact string    `pg:"owner_contact, notnull"`
	Admin        string    `pg:"admin, notnull"`
	AdminContact string    `pg:"admin_contact, notnull"`
	CreatedAt    time.Time `pg:"created_at, notnull"`
	ModifiedAt   time.Time `pg:"modified_at, notnull"`
	DeletedAt    time.Time `pg:"deleted_at,soft_delete"`
}

func New() *Department {
	return &Department{}
}

func (d *Department) Add() error {
	_, err := pgdb.DB().Model(d).Returning("*").Insert()

	return err
}

func (d *Department) Delete() error {
	_, err := pgdb.DB().Model(d).WherePK().Delete()

	return err
}

func (d *Department) Update() error {
	_, err := pgdb.DB().Model(d).WherePK().Update()

	return err
}

func (d *Department) List(offset int, limit int) ([]Department, int, error) {
	departs := make([]Department, 0)

	count, err := pgdb.DB().Model(d).Offset(offset).Limit(limit).Order("id asc").SelectAndCount(&departs)

	return departs, count, err
}

func (d *Department) Info() (Department, error) {
	depart := Department{}

	err := pgdb.DB().Model(d).WherePK().Select(&depart)

	return depart, err
}
