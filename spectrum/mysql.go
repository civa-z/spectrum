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
	log.Println(url)
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

func (mql *Mysql) GetUsingFrequency(name string) (Freq_Using_List) {
    var freq_using_list Freq_Using_List
    freq_using_list.Name=name

    command := fmt.Sprintf(`select id, districtcode, channel, power from %s`, name)
	log.Println(command)
    rows, err := mql.Conn.Query(command)
    if err != nil {
        log.Println("command:", command)
        log.Println(err)
        return freq_using_list
    }
    defer rows.Close()

    for rows.Next() {
        var freq_using Freq_Using
        err := rows.Scan(
            &freq_using.ID,
            &freq_using.DistrictCode,
            &freq_using.Channel,
            &freq_using.Power)
        if err != nil {
            log.Println(err)
        }
	freq_using_list.FreqUsingList=append(freq_using_list.FreqUsingList, freq_using)
    }
    return freq_using_list
}

func (mql *Mysql) GetFrequency() ([]Frequency) {
    var freq_list []Frequency

    command := `SELECT * FROM Frequency`
    rows, err := mql.Conn.Query(command)
    if err != nil {
        log.Println("command:", command)
        log.Println(err)
        return freq_list
    }
    defer rows.Close()

    for rows.Next() {
        var freq Frequency
        err := rows.Scan(
            &freq.ChannelID,
            &freq.Channel,
            &freq.Video,
            &freq.Audio,
            &freq.Center,
            &freq.Low,
            &freq.High)
        if err != nil {
            log.Println(err)
        }
	freq_list=append(freq_list, freq)
    }
    return freq_list
}

func (mql *Mysql) GetLocationInfo() ([]LocationInfo) {
    var locationinfos []LocationInfo

    command := `SELECT * FROM LocationInfo`
    rows, err := mql.Conn.Query(command)
    if err != nil {
        log.Println("command:", command)
        log.Println(err)
        return locationinfos
    }
    defer rows.Close()

    for rows.Next() {
        var locationinfo LocationInfo
        err := rows.Scan(
            &locationinfo.Id,
            &locationinfo.Province,
            &locationinfo.City,
            &locationinfo.District,
            &locationinfo.Code)
        if err != nil {
            log.Println(err)
        }
	locationinfos=append(locationinfos, locationinfo)
    }
    return locationinfos
}

func (mql *Mysql) GetUsingFreqLocal(location_code string) ([]Freq_Using_Local) {
    var freq_using_local_list []Freq_Using_Local

    command := fmt.Sprintf(`SELECT * FROM FreqUsing where districtcode = '%s'`, location_code)
    rows, err := mql.Conn.Query(command)
    if err != nil {
        log.Println("command:", command)
        log.Println(err)
        return freq_using_local_list
    }
    defer rows.Close()

    for rows.Next() {
        var freq_using_local Freq_Using_Local
        err := rows.Scan(
            &freq_using_local.FreqUsing.ID,
            &freq_using_local.FreqUsing.DistrictCode,
            &freq_using_local.Latitude,
            &freq_using_local.Longtitude,
            &freq_using_local.FreqUsing.Channel,
            &freq_using_local.FreqUsing.Power)
        if err != nil {
            log.Println(err)
        }
	freq_using_local_list = append(freq_using_local_list, freq_using_local)
    }
    return freq_using_local_list
}

