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
func testRawQuery() {
	db := database.GetConnection()

	var podConfigurations []model.PodConfiguration
	query := `select p.id                  pod_id,
				   p.name                pod_name,
				   p.status              status,
				   d.id                  "configuration__datastore_id",
				   d.name                "configuration__datastore_name",
				   c.id                  "configuration__cluster_id",
				   c.name                "configuration__cluster_name",
				   v.id                  "configuration__vcenter_id",
				   v.ip_address          "configuration__ip_address",
				   v.user_name           "configuration__user_name",
				   v.password            "configuration__password",
				   rt.type_name          "configuration__type_name",
				   rt.vapp_template_name "configuration__template_name"
			from pods p
					 inner join datastores d on p.datastore_id = d.id
					 inner join clusters c on p.cluster_id = c.id
					 inner join vcenters v on c.vcenter_id = v.id
					 inner join subscription_types st on p.subscription_type_id = st.id
					 inner join ref_types rt on st.ref_type_id = rt.id
			where p.id in (6, 7, 8)`

	_, err := db.Query(&podConfigurations, query)

	if err != nil {
		// Error Handler
	}
	fmt.Println(podConfigurations)
	fmt.Println(podConfigurations[0].Configuration.IpAddress)

}

func main() {
	testRawQuery()
	//db := database.GetConnection()
	//
	//tenant := &model.Tenant{
	//	Name:          "Rizwan123",
	//	Email:         "ee@gmail.com",
	//	RootAccountID: 1,
	//}
	//
	//err := db.Insert(tenant)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("Tenant added")
	//
	////inserting pod
	//pod := &model.Pod{
	//	Name:      "Pod 1",
	//	Status:    "Pending",
	//	IpAddress: "10.2.0.11",
	//}
	//
	//err = db.Insert(pod)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("Pod added")
	//
	//testJoins()

}
