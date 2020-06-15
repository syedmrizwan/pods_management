package main

import (
	"fmt"
	"github.com/syedmrizwan/pods_management/database"
	"github.com/syedmrizwan/pods_management/model"
)

func testJoins() {
	db := database.GetConnection()

	var tenants []model.Tenant
	err := db.Model(&tenants).
		Column("tenant.*").
		Relation("SubscriptionTypes").
		Relation("SubscriptionTypes.Pods").
		Relation("SubscriptionTypes.RefType").
		Select()
	if err != nil {
		// Error Handler
	} else {
		fmt.Println(tenants)
		fmt.Println(tenants[0].SubscriptionTypes[1].Pods[0].Name)
		fmt.Println(tenants[0].SubscriptionTypes[1].RefType.TypeName)
	}
}

func main() {
	db := database.GetConnection()

	tenant := &model.Tenant{
		Name:          "Rizwan123",
		Email:         "ee@gmail.com",
		RootAccountID: 1,
	}

	err := db.Insert(tenant)
	if err != nil {
		panic(err)
	}
	fmt.Println("Tenant added")

	//inserting pod
	pod := &model.Pod{
		Name:      "Pod 1",
		Status:    "Pending",
		IpAddress: "10.2.0.11",
	}

	err = db.Insert(pod)
	if err != nil {
		panic(err)
	}
	fmt.Println("Pod added")

	testJoins()

}
