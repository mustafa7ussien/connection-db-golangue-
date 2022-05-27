package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	fmt.Println("Go MySQL Tutorial")

	// if there is an error opening the connection, handle it

	pswd := "roooooot"
	db, err := sql.Open("mysql", "root:"+pswd+"@tcp(localhost:3306)/ddb")
	if err != nil {
		panic(err.Error())
	}
	// defer the close till after the main function has finished
	// executing
	defer db.Close()
	fmt.Println("opened")
	fmt.Println(pswd)

	//////////////delete
	// delete, err := db.Query("DELETE FROM t1 WHERE id = 1")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer delete.Close()
	// fmt.Println("deleted done")

	/////////////////insert
	insert, err := db.Query("INSERT INTO t1 VALUES (7,'Mostafa Hussien','mostafa.hussien.fci@gmail.com','software engineer',23,'problem solver','Open-Source Enthusiast.')")
	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
	fmt.Println("insert done")

	/////////////////Add clunmne in table with "Alter"
	// alter, err := db.Query("ALTER TABLE t1 ADD Email VARCHAR(200) DEFAULT 'Mostafa@gmail.com'")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer alter.Close()
	// fmt.Printf("alter table done")

}
