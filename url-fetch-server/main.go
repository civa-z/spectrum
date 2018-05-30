//gohttps/1-http/spectrum.go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "strings"
    "github.com/robfig/config"
)

const(
    Page404 = "<html>\r\n<head><title>404 Not Found</title></head>\r\n<body bgcolor=\"white\">\r\n<center><h1>404 Not Found</h1></center>\r\n"
    configurationFile = "url-fetch-server.config"
)

type DB struct {
    MCWSDSupport   bool   `json:"MCWSD_support"`
    DbProviderName string `json:"db_provider_name"`
    URL            string `json:"url"`
    WsDbID         string `json:"ws_db_id"`
}


type Response struct {
	WsDatabases struct {
		Db []DB `json:"db"`
		LastUpdate  string `json:"last_update"`
		RefreshRate string `json:"refresh_rate"`
	} `json:"ws_databases"`
}

type Configuration struct{
    IP string
    port int
    Url_DB string
}

var conf Configuration

func ProcessReq(r *http.Request) ([]byte, int) {
    url := r.URL.String()
    if !strings.Contains(url, "weblist.json") {
        return nil, 404
    }

    var resp Response
    var db DB
    db.URL = conf.Url_DB
    db.DbProviderName = "Nominet UK"
    db.WsDbID = "2"
    db.MCWSDSupport = true

    resp.WsDatabases.LastUpdate = "2016-06-30T12:00:00"
    resp.WsDatabases.RefreshRate = "1440"
    resp.WsDatabases.Db = append(resp.WsDatabases.Db, db)

    resp_body_byte, e := json.Marshal(resp)
    if e != nil{
        return nil, 404
    }

    return resp_body_byte, 0
}

func readConfig(file string, conf *Configuration) (bool){
    configdata, err := config.ReadDefault(file)
    if err != nil {
        log.Println(err)
        return false
	}

    var IP string
    var port int
    var url_db string

    IP, err = configdata.String("server", "IP")
    if err != nil {
        log.Println(err)
        return false
    }
    conf.IP = IP

    port, err = configdata.Int("server", "port")
    if err != nil {
        log.Println(err)
        return false
    }
    conf.port = port

    url_db, err = configdata.String("server", "url_db")
    if err != nil {
        log.Println(err)
        return false
    }
    conf.Url_DB = url_db
    return true
}

func handler(w http.ResponseWriter, r *http.Request) {
    log.Println(r.Method, r.URL.String(), r.RemoteAddr)
    var resp []byte
    var err_num int
    resp, err_num = ProcessReq(r)
    if err_num == 0 {
        w.Header().Set("content-length", strconv.Itoa(len(resp)))
	    w.Write(resp)
    } else {
        w.WriteHeader(404)
        fmt.Fprintf(w, Page404)
    }

    return
}

func main() {
    readConfig(configurationFile, &conf)
    http.HandleFunc("/", handler)
    log.Println(conf.IP + ":" + strconv.Itoa(conf.port))
    http.ListenAndServe(conf.IP + ":" + strconv.Itoa(conf.port), nil)
    //http.ListenAndServeTLS(":8081", "server.crt", "server.key", nil)
}

