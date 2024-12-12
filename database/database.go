package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var (
	DB *sql.DB
)

func TestDB() {
	fmt.Println("open db")

	//打开数据库，如果不存在，则创建
	db, err := sql.Open("sqlite3", "./test.db")
	checkErr(err)

	//创建表
	fmt.Println("create table")
	sql_table := `
		    CREATE TABLE IF NOT EXISTS reg_logs(
		        lid INTEGER PRIMARY KEY AUTOINCREMENT,
		        name VARCHAR(64) NULL,
				label VARCHAR(64) NULL,
		        owner VARCHAR(64) NULL,
				premium VARCHAR(64) NULL,
				expires VARCHAR(64) NULL
		    );
		    `
	db.Exec(sql_table)

	// 插入数据
	fmt.Println("insert data")
	stmt, err := db.Prepare("INSERT INTO reg_logs(name, label, owner, premium, expires) values(?,?,?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("123", "123", "123", "123", "123")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("insert id: ", id)

	// //更新数据
	// fmt.Println("update data")
	// stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	// checkErr(err)

	// res, err = stmt.Exec("astaxieupdate", id)
	// checkErr(err)

	// affect, err := res.RowsAffected()
	// checkErr(err)

	// fmt.Println("affect: ", affect)

	//查询数据
	fmt.Println("show rows:")
	rows, err := db.Query("SELECT * FROM reg_logs")
	checkErr(err)

	for rows.Next() {
		var uid int
		var name string
		var label string
		var owner string
		var premium string
		var expires string
		err = rows.Scan(&uid, &name, &label, &owner, &premium, &expires)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(name)
		fmt.Println(label)
		fmt.Println(owner)
		fmt.Println(premium)
		fmt.Println(expires)
		fmt.Println()
	}

	// //删除数据
	// stmt, err = db.Prepare("delete from userinfo where uid=?")
	// checkErr(err)

	// res, err = stmt.Exec(id)
	// checkErr(err)

	// affect, err = res.RowsAffected()
	// checkErr(err)

	// fmt.Println(affect)

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// create db
func Open() {
	fmt.Println("open db")

	//打开数据库，如果不存在，则创建
	db, err := sql.Open("sqlite3", "./ens.db")
	checkErr(err)

	//创建表
	fmt.Println("create table")
	sql_table := `
    CREATE TABLE IF NOT EXISTS reg_logs(
        lid INTEGER PRIMARY KEY AUTOINCREMENT,
        name VARCHAR(64) NULL,
		label VARCHAR(64) NULL,
        owner VARCHAR(64) NULL,
		basecost VARCHAR(64) NULL,
		expires VARCHAR(64) NULL
    );
    `

	db.Exec(sql_table)

	DB = db
}

func Insert(name string, label string, owner string, basecost string, expires string) {
	// 插入数据
	stmt, err := DB.Prepare("INSERT INTO reg_logs(name, label, owner, basecost, expires) values(?,?,?,?,?)")
	checkErr(err)

	res, err := stmt.Exec(name, label, owner, basecost, expires)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	_ = id

	//fmt.Println("insert id: ", id)
}

// read an user from db by id
func ReadUserByID(id uint64) (string, error) {
	// check
	if DB == nil {
		return "", fmt.Errorf("db not open")
	}

	// 查询语句
	query := "SELECT owner FROM reg_logs WHERE lid = ?"

	// 准备查询
	stmt, err := DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// 执行查询
	rows, err := stmt.Query(fmt.Sprintf("%d", id))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var owner string
	// 遍历结果集
	for rows.Next() {
		if err := rows.Scan(&owner); err != nil {
			log.Fatal(err)
		}
	}

	// 检查遍历是否出错
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return owner, nil
}
