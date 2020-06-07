package model

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"time"
)

type Root struct {
	AccountID int64 `pg:",pk" json:"account_id"`
}

func (r *Root) String() string {
	return fmt.Sprintf("Root<%d>", r.AccountID)
}

type Tenant struct {
	ID            int64
	Name          string
	Email         string
	RootAccountID int64 `pg:",fk"`
	Root          *Root
	CreatedAt     time.Time `pg:"default:now()"`
}

func (t *Tenant) String() string {
	return fmt.Sprintf("Tenant<%d %s %s>", t.ID, t.Name, t.Email)
}

// createSchema creates database schema for models.
func CreateSchema(db *pg.DB) error {
	models := []interface{}{
		(*Root)(nil),
		(*Tenant)(nil),
	}

	for _, model := range models {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			//Temp: true, // temp table
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
