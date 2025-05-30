package models

import (
	"database/sql"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbDriver := "mysql"
	dbHost :="127.0.0.1"
	dbPort := "3306"
	dbUser := "root"
	dbPass := "yash@jain123"
	dbName := "crud_db"

	orm.RegisterDriver(dbDriver,orm.DRMySQL)
	connStr := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName

	err := orm.RegisterDataBase("default", dbDriver, connStr)
	if err != nil {
		fmt.Println("Database registration error:", err)
	}

	db, err := sql.Open(dbDriver, connStr)
	if err != nil {
		fmt.Println("Failed to open database:", err)
	} else {
		err = db.Ping()
		if err != nil {
			fmt.Println("Database connection failed:", err)
		} else {
			fmt.Println("✅ Database connected successfully.")
		}
	}

	orm.RunSyncdb("default", false, true)
}