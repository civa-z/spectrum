package main

import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/go-sql-driver/mysql"
)


type Mysql struct{
    Conn *sql.DB
}

func (mql *Mysql) MysqlOpen(db_name string, mysql_ip string, mysql_port int) int {
    var err error
    url := fmt.Sprintf("sony:sony@tcp(%s:%d)/%s", mysql_ip, mysql_port, db_name)
    mql.Conn, err = sql.Open("mysql", url)
        if err == nil {
        return 0
    } else {
        log.Println(err)
        return -1
    }
}

func (mql *Mysql) MySqlClose() {
	mql.Conn.Close()
}


func (mql *Mysql) GetAllPtx(IP_address string) (All_Ptx) {
    var ap All_Ptx

    //command := fmt.Sprintf(`SELECT * FROM all_ptx where ip_address='%s'`, IP_address)
    command := fmt.Sprintf(`SELECT * FROM all_ptx where ip_address='192.168.191.1'`)
    rows, err := mql.Conn.Query(command)
    if err != nil {
        log.Println("command:", command)
        log.Println(err)
        return ap
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(
            &ap.id,
            &ap.IP_address,
            &ap.location_x,
            &ap.location_y,
            &ap.channel,
            &ap.PTX,
            &ap.OFCOM,
            &ap.ECC,
            &ap.QoS,)
        if err != nil {
            log.Println(err)
        }
    }
    return ap
}

