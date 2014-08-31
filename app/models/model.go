package models

import (
	"database/sql"
	"revel-boilerplate/app/entities"

	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
)

type Model struct {
	Dbm *gorp.DbMap
}

func (self *Model) Init() {
	// initialize the DbMap
	var err error
	self.Dbm, err = initDb()
	if err != nil {
		panic(err)
	}
}

func initDb() (*gorp.DbMap, error) {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := sql.Open("sqlite3", "/tmp/db.bin")
	if err != nil {
		return nil, err
	}

	// construct a gorp DbMap
	dbm := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	dbm.AddTableWithName(entities.User{}, "users").SetKeys(true, "id")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbm.CreateTablesIfNotExists()
	if err != nil {
		return nil, err
	}

	return dbm, nil
}
