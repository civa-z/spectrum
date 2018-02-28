package main

import (
    "database/sql"
    "fmt"
    "log"
)


type Mysql struct{
    Conn *sql.DB
}

func (mql *Mysql) MysqlOpen(db_name string, mysql_port int) int {
    var err error
    url := fmt.Sprintf("root:meimima.1@tcp(127.0.0.1:%d)/%s", mysql_port, db_name)
    mql.Conn, err = sql.Open("mysql", url)
        if err == nil {
        return 0
    } else {
        log.Println(err)
        return -1
    }
}


