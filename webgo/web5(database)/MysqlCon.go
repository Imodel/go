package main

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/test?charset=utf8")
	checkErr(err)

	/*stmt, err := db.Prepare("INSERT INTO userinfo (username,departname,created) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("liuxk", "study", "2017-05-22")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)*/

	stmt, err := db.Prepare("UPDATE userinfo SET username=? where uid = ?")
	checkErr(err)

	res, err := stmt.Exec("xkliue", 1)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
