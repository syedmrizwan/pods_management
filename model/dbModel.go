package model

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"time"
)

type Company struct {
	ID                int64     `pg:",pk" json:"id"`
	Name              string    `json:"name"`
	Email             string    `json:"email"`
	CreatedAt         time.Time `pg:"default:now()" json:"created_at"`
	SubscriptionTypes []*SubscriptionType
}
type Root struct {
	AccountID int64    `pg:",pk" json:"account_id"`
	CompanyID int64    `pg:",fk" json:"company_id"`
	Company   *Company `json:"company"`
}

type Tenant struct {
	ID             int64     `pg:",pk" json:"id"`
	CompanyID      int64     `pg:",fk" json:"company_id"`
	Company        *Company  `json:"company"`
	RootAccountID  int64     `pg:",fk" json:"root_account_id"`
	ActivateLater  bool      `pg:",use_zero" json:"activate_later"`
	ActivationTime time.Time `json:"activation_time"`
	Root           *Root     `json:"root"`
}

type SubTenant struct {
	ID        int64    `pg:",pk" json:"id"`
	CompanyID int64    `pg:",fk" json:"company_id"`
	Company   *Company `json:"company"`
	TenantID  int64    `pg:",fk" json:"tenant_id"`
	Tenant    *Tenant  `json:"tenant"`
}

type Pod struct {
	ID                          int64             `pg:",pk" json:"id"`
	Name                        string            `json:"name"`
	DisplayName                 string            `json:"display_name"`
	DatastoreID                 int64             `pg:",fk" json:"datastore_id"`
	Datastore                   *Datastore        `json:"datastore"`
	ClusterID                   int64             `pg:",fk" json:"cluster_id"`
	Cluster                     *Cluster          `json:"cluster"`
	TenantSubscriptionTypeID    int64             `pg:",fk:subscription_type_id" json:"tenant_subscription_type_id"`
	TenantSubscriptionType      *SubscriptionType `json:"tenant_subscription_type"`
	SubTenantSubscriptionTypeID int64             `pg:",fk:subscription_type_id" json:"sub_tenant_subscription_type_id"`
	SubTenantSubscriptionType   *SubscriptionType `json:"sub_tenant_subscription_type"`
	Status                      string            `json:"status"`
	IpAddress                   string            `json:"ip_address"`
	StudentID                   int64             `pg:",fk" json:"student_id"`
	Student                     *Student          `json:"student"`
	TaskID                      string            `json:"task_id"`
	IsExpired                   bool              `pg:"default:false" json:"is_expired"`
	CreatedAt                   time.Time         `pg:"default:now()" json:"created_at"`
}

type SubscriptionType struct {
	ID          int64     `pg:",pk" json:"id"`
	Description string    `json:"description"`
	RefTypeID   int64     `pg:",fk" json:"ref_type_id"`
	RefType     *RefType  `json:"ref_type"`
	ExpiryTime  time.Time `json:"expiry_time"`
	CompanyID   int64     `pg:",fk" json:"company_id"`
	Company     *Company  `json:"company"`
	IsActivated bool      `json:"is_activated"`
	Pods        []*Pod
}

type VcenterTemplate struct {
	ID               int64    `pg:",pk" json:"id"`
	VappTemplateName string   `json:"vapp_template_name"`
	RefTypeID        int64    `pg:",fk" json:"ref_type_id"`
	RefType          *RefType `json:"ref_type"`
	VcenterID        int64    `pg:",fk" json:"vcenter_id"`
	Vcenter          *Vcenter
}

type RefType struct {
	ID               int64  `pg:",pk" json:"id"`
	TypeName         string `json:"type_name"`
	Description      string `json:"description"`
	TrainingContents []*TrainingContent
}

type TrainingContent struct {
	ID        int64  `pg:",pk" json:"id"`
	Name      string `json:"name"`
	Content   []byte `pg:"type:bytea" json:"content"`
	RefTypeID int64  `pg:",fk" json:"ref_type_id"`
}

type Vcenter struct {
	ID               int64              `pg:",pk" json:"id"`
	IpAddress        string             `json:"ip_address"`
	UserName         string             `json:"user_name"`
	Password         string             `json:"password"`
	VcenterTemplates []*VcenterTemplate `json:"vcenter_templates"`
	Clusters         []*Cluster
	Datastores       []*Datastore
}

type Cluster struct {
	ID        int64  `pg:",pk" json:"id"`
	Name      string `json:"name"`
	VcenterID int64  `pg:",fk" json:"vcenter_id"`
}

type Datastore struct {
	ID        int64  `pg:",pk" json:"id"`
	Name      string `json:"name"`
	VcenterID int64  `pg:",fk" json:"vcenter_id"`
}

type Student struct {
	ID        int64     `pg:",pk" json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CompanyID int64     `pg:",fk" json:"company_id"`
	Company   *Company  `json:"company"`
	CreatedAt time.Time `pg:"default:now()" json:"created_at"`
}

type Request struct {
	ID               int64           `pg:",pk" json:"id"`
	Title            string          `json:"title"`
	Description      string          `json:"description"`
	Status           string          `json:"status"`
	ToCompanyID      int64           `pg:",fk:company_id" json:"to_company_id"`
	ToCompany        *Company        `json:"company"`
	FromCompanyID    int64           `pg:",fk:subscription_type_id" json:"from_company_id"`
	FromCompany      *Company        `json:"company"`
	RefTypeID        int64           `pg:",fk" json:"ref_type_id"`
	RefType          *RefType        `json:"ref_type"`
	RequestRefTypeID int64           `pg:",fk" json:"request_ref_type_id"`
	RequestRefType   *RequestRefType `json:"request_ref_type"`
	CreatedAt        time.Time       `pg:"default:now()" json:"created_at"`
}

type RequestRefType struct {
	ID          int64  `pg:",pk" json:"id"`
	RequestType string `json:"request_type"`
}

type RequestDetail struct {
	ID           int64    `pg:",pk" json:"id"`
	DurationDays int64    `pg:json:"duration_days"`
	NumberOfPods int64    `pg:json:"number_of_pods"`
	RequestID    int64    `pg:",fk" json:"request_id"`
	Request      *Request `json:"request"`
}

// createSchema creates database schema for models.
func CreateSchema(db *pg.DB) error {
	models := []interface{}{
		(*Root)(nil),
		(*Company)(nil),
		(*Tenant)(nil),
		(*SubTenant)(nil),
		(*Pod)(nil),
		(*SubscriptionType)(nil),
		(*RefType)(nil),
		(*Vcenter)(nil),
		(*Cluster)(nil),
		(*Datastore)(nil),
		(*TrainingContent)(nil),
		(*Student)(nil),
		(*Request)(nil),
		(*RequestRefType)(nil),
		(*VcenterTemplate)(nil),
		(*RequestDetail)(nil),
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
