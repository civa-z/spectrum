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

func (mql *Mysql) GetUsingFrequency(name string) ([]Freq_Using) {
    var freq_using_list []Freq_Using

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
		freq_using_list=append(freq_using_list, freq_using)
    }
    return freq_using_list
}

func (mql *Mysql) GetFrequency() ([]Frequency) {
log.Println("GetFrequency")
    var freq_list []Frequency

    command := `SELECT * FROM Frequency`
	log.Println(command)
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

func (mql *Mysql) GetLocationInfo() ([]Location_Info) {
    var location_infos []Location_Info

    command := `SELECT * FROM LocationInfo`
    rows, err := mql.Conn.Query(command)
    if err != nil {
        log.Println("command:", command)
        log.Println(err)
        return location_infos
    }
    defer rows.Close()

    for rows.Next() {
        var location_info Location_Info
        err := rows.Scan(
            &location_info.Id,
            &location_info.Province,
            &location_info.City,
            &location_info.District,
            &location_info.Code)
        if err != nil {
            log.Println(err)
        }
	location_infos=append(location_infos, location_info)
    }
    return location_infos
}

func (mql *Mysql) GetOnlineDevice(location_code string) ([]Online_Device) {
    var online_device_list []Online_Device
	var command string
	if location_code == "*"{
		command = fmt.Sprintf(`SELECT * FROM OnlineDevice`)
	} else {
		command = fmt.Sprintf(`SELECT * FROM OnlineDevice where DistrictCode == %s`, location_code)
	}
	log.Println(command)
    rows, err := mql.Conn.Query(command)
    if err != nil {
        log.Println("command:", command)
        log.Println(err)
        return online_device_list
    }
    defer rows.Close()

    for rows.Next() {
        var online_device Online_Device
        err := rows.Scan(
			&online_device.SerialNumber,
            &online_device.Latitude,
            &online_device.Longtitude,
            &online_device.FreqUsing.ID,
            &online_device.FreqUsing.DistrictCode,
            &online_device.FreqUsing.Channel,
            &online_device.FreqUsing.Power)
        if err != nil {
            log.Println(err)
        }
	online_device_list = append(online_device_list, online_device)
    }
    return online_device_list
}

func (mql *Mysql) InsertOnlineDevice(online_device Online_Device) {
	command := fmt.Sprintf(`INSERT into OnlineDevice(serialnumber districtcode latitude longtitude channel power) values(%s %s %f %f %d %f)`,
		online_device.SerialNumber,
		online_device.Latitude,
		online_device.Longtitude,
		online_device.FreqUsing.Channel,
		online_device.FreqUsing.Power)
		
	result, err := mql.Conn.Exec(command)
	if err != nil {
		fmt.Println("insert data failed:", err.Error())
		return
	}
	
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("fetch last insert id failed:", err.Error())
		return
	}
	fmt.Println("insert new record", id)
	return
}
