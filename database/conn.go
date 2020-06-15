package database

import (
	"github.com/go-pg/pg/v10"
	"github.com/syedmrizwan/pods_management/env"
	"github.com/syedmrizwan/pods_management/model"
)

var db *pg.DB

func init() {
	//todo add more config to the connection like idle timeout etc
	db = pg.Connect(&pg.Options{
		Addr:     env.Env.GetAddr(),
		User:     env.Env.DbUsername,
		Password: env.Env.DbPassword,
		Database: env.Env.DbName,
		PoolSize: env.Env.DbPoolSize,
	})
	if err := model.CreateSchema(db); err != nil {
		panic(err)
	}
	//Insert prerequisite data
	InsertPrerequisite()
}

func GetConnection() *pg.DB {
	return db
}

func populateRoot() error {
	root := &model.Root{AccountID: 1}
	if _, err := db.Model(root).Where("account_id = ?account_id").SelectOrInsert(); err != nil {
		return err
	}
	return nil
}

func selectOrInsertRefType(typeName string, template string) (*model.RefType, error) {
	refType := &model.RefType{
		TypeName: typeName,
		VappTemplateName: template,
	}
	if _, err := db.Model(refType).Where("type_name = ?type_name").SelectOrInsert(); err != nil {
		return nil, err
	}
	return refType, nil
}

//TODO: populate configuration dynamically
func populateConfiguration() error {
	vcenter := &model.Vcenter{
		IpAddress: "127.0.0.1",
		UserName:  "rizwan",
		Password:  "password",
	}
	if _, err := db.Model(vcenter).Where("ip_address = ?ip_address").SelectOrInsert(); err != nil {
		return err
	}

	datastore := &model.Datastore{
		Name:      "Datastore",
		VcenterID: vcenter.ID,
	}
	if _, err := db.Model(datastore).Where("name = ?name").SelectOrInsert(); err != nil {
		return err
	}

	cluster := &model.Cluster{
		Name:      "Cluster",
		VcenterID: vcenter.ID,
	}
	if _, err := db.Model(cluster).Where("name = ?name").SelectOrInsert(); err != nil {
		return err
	}

	return nil
}

func InsertPrerequisite() error {
	if err := populateRoot(); err != nil {
		return err
	}
	if _, err := selectOrInsertRefType("Type-A","Type-A"); err != nil {
		return err
	}
	if _, err := selectOrInsertRefType("Type-B", "Type-B"); err != nil {
		return err
	}
	if err := populateConfiguration(); err != nil {
		return err
	}
	return nil

}
