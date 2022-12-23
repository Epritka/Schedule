package pg

import (
	"fmt"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type dbLogger struct{}

func (d dbLogger) BeforeQuery(q *pg.QueryEvent) {
}

func (d dbLogger) AfterQuery(q *pg.QueryEvent) {
	sql, _ := q.FormattedQuery()
	fmt.Println(sql)
}

func New(dbUrl string, IsDebug bool) *pg.DB {
	// Нужно для того что бы ORM искал таблицы с именем в единственном числе
	orm.SetTableNameInflector(func(s string) string {
		return s
	})
	opt, err := pg.ParseURL(dbUrl)
	if err != nil {
		panic(err.Error())
	}
	db := pg.Connect(opt)
	if IsDebug {
		db.AddQueryHook(dbLogger{})
	}
	return db
}
