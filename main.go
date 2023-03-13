//
//  go-unit-test-sql
//
//  Copyright © 2020. All rights reserved.
//

package main

import (
	"fmt"

	"go-unit-test-sql/repository/mysql"
)

func main() {
	// build the DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "root", "23031996", "127.0.0.1", 3306, "user")
	repo, err := mysql.NewRepository("mysql", dsn, 1000, 1000)
	if err != nil {
		panic(err)
	}

	user, err := repo.FindByID("1")
	if err != nil {
		fmt.Printf("error FindByID %v", err)
	}
	defer repo.Close()

	fmt.Println(user)
	fmt.Println("DONE")
}
