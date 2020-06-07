package database

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/syedmrizwan/pods_management/env"
	"github.com/syedmrizwan/pods_management/model"
)

var db *pg.DB

func init() {
	db = pg.Connect(&pg.Options{
		Addr:     env.Env.GetAddr(),
		User:     env.Env.DbUsername,
		Password: env.Env.DbPassword,
		Database: env.Env.DbName,
		PoolSize: env.Env.DbPoolSize,
	})
	err := model.CreateSchema(db)
	if err != nil {
		panic(err)
	}
	//check and create root user
	GetRoot()

}

// GetConnection returns pg connection
func GetConnection() *pg.DB {
	return db
}

func GetRoot() *model.Root {
	// Select user by primary key.
	root := &model.Root{AccountID: 1}
	err := db.Select(root)
	if err != nil {
		// adding root user
		root := &model.Root{
			AccountID: 1,
		}
		err := db.Insert(root)
		if err != nil {
			panic(err)
		}
		fmt.Println("Root user added")
	}
	return root
}
