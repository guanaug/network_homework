package transaciton

import (
	"github.com/go-pg/pg/orm"
	"network/global/pgdb"
	"time"
)

type Transaction struct {
	tableName         struct{}  `sql:"network_homework.tb_transaction,discard_unknown_columns"`
	ID                int64     `sql:"id,pk"`
	Publisher         int64     `pg:"publisher"`
	CreatedAt         time.Time `pg:"created_at"`
	Type              int8      `pg:"type"`
	Status            int8      `pg:"status"`
	Detail            string    `pg:"detail"`
	TranType          int8      `pg:"tran_type"`
	HandlerDepartment int64     `pg:"handler_department"`
	Handler           int64     `pg:"handler"`
	ModifiedAt        time.Time `pg:"modified_at"`
}

func New() *Transaction {
	return &Transaction{}
}

func (t *Transaction) Model() *orm.Query {
	return pgdb.DB().Model(t)
}

func (t *Transaction) Insert() error {
	_, err := t.Model().WherePK().Returning("*").Insert()

	return err
}

func (t *Transaction) Update() error {
	_, err := t.Model().WherePK().UpdateNotNull()

	return err
}