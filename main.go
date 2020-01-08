package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-oci8"
)

func main() {
	username := "username"
	password := "password"
	host := "host"
	port := "port"
	serviceName := "servicename"

	db, err := sql.Open("oci8", username+"/"+password+"@"+"//"+host+":"+port+"/"+serviceName)

	if err != nil {
		fmt.Println("... DB Setup Failed")

		fmt.Println(err)
		return
	}

	defer db.Close()
	fmt.Println("... Opening Database Connection")

	if err = db.Ping(); err != nil {
		fmt.Println("error")
		fmt.Println(err)
		return
	}

	fmt.Println("... Connected to Database")

	dbQuery := "select username from nettv_auto_bonus where username='harshu2066_fbnpa'"
	rows, err := db.Query(dbQuery)

	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}

	defer rows.Close()
	var userName string
	for rows.Next() {
		if err = rows.Scan(&userName); err != nil {
			log.Fatal(err)
		}

		fmt.Println(userName)
	}

}
