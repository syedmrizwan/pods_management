package main

import (
	"fmt"
	"github.com/syedmrizwan/pods_management/database"
	"github.com/syedmrizwan/pods_management/model"
)

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

}
