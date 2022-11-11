package main

import (
	"fmt"
	"frame_one/mysql"
)

func main() {

	db := mysql.NewMySQL(&mysql.Config{
		Dsn:    "root:123456@tcp(127.0.0.1:3306)/picture?charset=utf8&parseTime=True&loc=Asia%2FShanghai",
		Driver: "mysql",
		// ShowSql: true,
	})
	fmt.Println(db)
	// rows, err := db.Query("select * from app")
	// // rows, err := db.Query("select * from app")
	// fmt.Println(err)
	// fmt.Println(rows)

}
