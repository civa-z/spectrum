//gohttps/1-http/spectrum.go
package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
    "golang.org/x/net/websocket"
    "strings"
    "encoding/json"
)


type UserMsg struct {
    UserName string
    Msg      string
    DataType string
}

type UserData struct {
    UserName string
}

type Datas struct {
    UserMsgs  []UserMsg
    UserDatas []UserData
}

//Global Information
var datas Datas
var mysql Mysql


var users map[*websocket.Conn]string

func h_index(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index.html")
}

func h_webSocket(ws *websocket.Conn){
    for {
        fmt.Println("开始解析数据...")
        var data string
        err := websocket.Message.Receive(ws, &data)
        fmt.Println("data：", data)
        if err != nil {
            fmt.Println(`接收出错... :%s`, err)
			ws.Close()
            break
        }

        data = strings.Replace(data, "\n", "", 0)
		fmt.Println(`接收:`, data)
        var webSocketReqInfo WebSocketReqInfo
        err = json.Unmarshal([]byte(data), &webSocketReqInfo)
        if err != nil {
            fmt.Println("解析数据异常...")
            break
        }
        fmt.Println("请求数据类型：", webSocketReqInfo.Type)
        b := getData(webSocketReqInfo.Type)

        errMarshl:=websocket.Message.Send(ws, string(b))
        if errMarshl != nil {
            //移除出错的链接
            fmt.Println("发送出错...")
            break
        }
    }

}

func getData(name string) ([]byte){
    var b []byte
    var errMarshl error
    switch name{
        case "CMMB":
			var freq_using_res Freq_Using_List_Res
			freq_using_res.Type = "CMMB"
		    freq_using_res.Freq_Using_List_ = db.GetUsingFrequency(name)
            b, errMarshl=json.Marshal(freq_using_res)
        case "DTMB":
			var freq_using_res Freq_Using_List_Res
			freq_using_res.Type = "DTMB"
		    freq_using_res.Freq_Using_List_ = db.GetUsingFrequency(name)
            b, errMarshl=json.Marshal(freq_using_res)
        case "TV":
			var freq_using_res Freq_Using_List_Res
			freq_using_res.Type = "TV"
            freq_using_res.Freq_Using_List_ = db.GetUsingFrequency(name)
            b, errMarshl=json.Marshal(freq_using_res)
        case "Frequency":
			var freq_res Frequency_Res
			freq_res.Type = "Frequency"
            freq_res.Frequency_ = db.GetFrequency()
            b, errMarshl=json.Marshal(freq_res)
        case "OnlineDevice":
			var online_device_res OnlineDevice_Res
			online_device_res.Type = "OnlineDevice"
			online_device_res.OnlineDevicelist = db.GetOnlineDevice("*")
            b, errMarshl=json.Marshal(online_device_res)
        case "LocationInfo"://Not implemented
            locationinfo:=db.GetLocationInfo()
            b, errMarshl=json.Marshal(locationinfo)
    }

    if errMarshl != nil {
        fmt.Println("取得数据异常...")
    }
    return b

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
        w.WriteHeader(400)
        fmt.Fprintf(w, "Hi, This is an example of http service in golang{!")
    }

    return
}

var db Mysql

func spectrum_init() {
    db.MysqlOpen("spectrum_v1", "127.0.0.1" ,3306)
    return
}



func main() {
    spectrum_init()
    http.HandleFunc("/", h_index)
    http.HandleFunc("/data", handler)
    http.Handle("/webSocket", websocket.Handler(h_webSocket))
    //err := http.ListenAndServe("192.168.1.110:443", nil)
    err := http.ListenAndServeTLS("192.168.1.112:443", "server.crt", "server.key", nil)
    fmt.Println(err)
}
