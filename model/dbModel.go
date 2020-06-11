package model

import (
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type Root struct {
	AccountID int64 `pg:",pk" json:"account_id"`
}

type Tenant struct {
	ID            int64     `pg:",pk" json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Email         string    `json:"email"`
	RootAccountID int64     `pg:",fk" json:"root_account_id"`
	IsActivated   bool      `json:"is_activated"`
	ActivateLater time.Time `json:"activate_later"`
	Root          *Root     `json:"root"`
	CreatedAt     time.Time `pg:"default:now()" json:"created_at"`
}

type Pod struct {
	ID                 int64             `pg:",pk" json:"id"`
	Name               string            `json:"name"`
	SubscriptionTypeID int64             `pg:",fk" json:"subscription_type_id"`
	SubscriptionType   *SubscriptionType `json:"subscription_type"`
	Status             string            `json:"status"`
	IpAddress          string            `json:"ip_address"`
	CreatedAt          time.Time         `pg:"default:now()" json:"created_at"`
}

type SubscriptionType struct {
	ID         int64     `pg:",pk" json:"id"`
	RefTypeID  int64     `pg:",fk" json:"ref_type_id"`
	RefType    *RefType  `json:"ref_type"`
	ExpiryTime time.Time `json:"expiry_time"`
	TenantID   int64     `pg:",fk" json:"tenant_id"`
	Tenant     *Tenant   `json:"tenant"`
}

type RefType struct {
	ID                  int64  `pg:",pk" json:"id"`
	TypeName            string `json:"type_name"`
	TrainingContentName string `json:"training_content_name"`
	TrainingContent     []byte `pg:"type:bytea" json:"training_content"`
}

// createSchema creates database schema for models.
func CreateSchema(db *pg.DB) error {
	models := []interface{}{
		(*Root)(nil),
		(*Tenant)(nil),
		(*Pod)(nil),
		(*SubscriptionType)(nil),
		(*RefType)(nil),
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
