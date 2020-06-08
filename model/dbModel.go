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

type Pod struct {
	ID                 int64
	Name               string
	Status             string
	IpAddress          string
	SubscriptionTypeID int64
	SubscriptionType   *SubType
	CreatedAt          time.Time `pg:"default:now()"`
}

func (p *Pod) String() string {
	return fmt.Sprintf("Pod<%d %s %s %s>", p.ID, p.Name, p.Status, p.IpAddress)
}

type SubType struct {
	ID         int64
	TypeName   string
	ExpiryTime time.Time
	TenantID   int64
	Tenant     *Tenant
}

func (s *SubType) String() string {
	return fmt.Sprintf("Pod<%d %s %s>", s.ID, s.TypeName, s.ExpiryTime)
}

// createSchema creates database schema for models.
func CreateSchema(db *pg.DB) error {
	models := []interface{}{
		(*Root)(nil),
		(*Tenant)(nil),
		(*Pod)(nil),
		(*SubType)(nil),
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
