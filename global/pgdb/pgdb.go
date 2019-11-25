package pgdb

import (
	"github.com/go-pg/pg"
	"log"
	"sync"
)

var db *pg.DB

func init()  {
	once := sync.Once{}

	once.Do(func() {
		db = pg.Connect(&pg.Options{
			User:                  "postgres",
			Password:              "root",
			Database:              "dogod",
			OnConnect: func(conn *pg.Conn) error {
				return nil
			},
		})

		db.AddQueryHook(queryHook{})
	})
}

func DB() *pg.DB {
	return db
}

type queryHook struct {
}

func (q queryHook) BeforeQuery(event *pg.QueryEvent) {
	sql, err := event.FormattedQuery()
	if err != nil {
		log.Println("SQL error.")
	} else {
		log.Printf("SQL: %s\n", sql)
	}
}

func (q queryHook) AfterQuery(event *pg.QueryEvent) {
}

