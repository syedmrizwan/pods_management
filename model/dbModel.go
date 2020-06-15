package model

import (
	"time"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type Root struct {
	AccountID int64 `pg:",pk" json:"account_id"`
}

type Tenant struct {
	ID                int64     `pg:",pk" json:"id"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	Email             string    `json:"email"`
	RootAccountID     int64     `pg:",fk" json:"root_account_id"`
	IsActivated       bool      `json:"is_activated"`
	ActivateLater     time.Time `json:"activate_later"`
	Root              *Root     `json:"root"`
	CreatedAt         time.Time `pg:"default:now()" json:"created_at"`
	SubscriptionTypes []*SubscriptionType
}

type Pod struct {
	ID                 int64             `pg:",pk" json:"id"`
	Name               string            `json:"name"`
	DatastoreID        int64             `pg:",fk" json:"datastore_id"`
	Datastore          *Datastore        `json:"datastore"`
	ClusterID          int64             `pg:",fk" json:"cluster_id"`
	Cluster            *Cluster          `json:"cluster"`
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
	Pods       []*Pod
}

type RefType struct {
	ID               int64  `pg:",pk" json:"id"`
	TypeName         string `json:"type_name"`
	VappTemplateName string `json:"vapp_template_name"`
}

type TrainingContent struct {
	ID        int64    `pg:",pk" json:"id"`
	Name      string   `json:"name"`
	Content   []byte   `pg:"type:bytea" json:"content"`
	RefTypeID int64    `pg:",fk" json:"ref_type_id"`
	RefType   *RefType `json:"ref_type"`
}

type Vcenter struct {
	ID        int64  `pg:",pk" json:"id"`
	IpAddress string `json:"ip_address"`
	UserName  string `json:"user_name"`
	Password  string `json:"password"`
}

type Cluster struct {
	ID        int64    `pg:",pk" json:"id"`
	Name      string   `json:"name"`
	VcenterID int64    `pg:",fk" json:"vcenter_id"`
	Vcenter   *Vcenter `json:"vcenter"`
}

type Datastore struct {
	ID        int64    `pg:",pk" json:"id"`
	Name      string   `json:"name"`
	VcenterID int64    `pg:",fk" json:"vcenter_id"`
	Vcenter   *Vcenter `json:"vcenter"`
}

// createSchema creates database schema for models.
func CreateSchema(db *pg.DB) error {
	models := []interface{}{
		(*Root)(nil),
		(*Tenant)(nil),
		(*Pod)(nil),
		(*SubscriptionType)(nil),
		(*RefType)(nil),
		(*Vcenter)(nil),
		(*Cluster)(nil),
		(*Datastore)(nil),
		(*TrainingContent)(nil),
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

// Struct to String conversion
func (r *Root) String() string {
	return fmt.Sprintf("Root<%d>", r.AccountID)
}

func (t *Tenant) String() string {
	return fmt.Sprintf("Tenant<%d %s %s %s>", t.ID, t.Name, t.Description, t.Email)
}

func (p *Pod) String() string {
	return fmt.Sprintf("Pod<%d %s %s %s>", p.ID, p.Name, p.Status, p.IpAddress)
}

func (s *SubscriptionType) String() string {
	return fmt.Sprintf("Pod<%d %s>", s.ID, s.ExpiryTime)
}
