package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // user:password@tcp(localhost:5555)/dbname?charset=utf8
    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
    fmt.Println(db, err)
    defer db.Close()

    stmt, err := db.Prepare("insert into student(name, age) values(?,?)")
    defer stmt.Close()
    fmt.Println(stmt, err)

    rs, err := stmt.Exec("xwz", 1000)
    fmt.Println(rs, err)

    id, err := rs.LastInsertId()
    fmt.Println(id, err)

    rows, err := db.Query("SELECT * FROM student")
    fmt.Println(rows, err)
    cols, _ := rows.Columns()
    for i := range cols {
        fmt.Println("col:", cols[i])
    }
    for rows.Next() {
        var uid int
        var name string
        var age int
        err = rows.Scan(&uid, &name, &age)
        fmt.Println(err)
        fmt.Println(uid, name, age)
    }
}
