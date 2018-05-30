//gohttps/1-http/spectrum.go
package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
	"golang.org/x/net/websocket"
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
var users map[*websocket.Conn]string

func h_index(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index.html")
}


type WebSocketReqInfo{
    Type string
}

type CMMB_Res struct {
    Type string
    Cmmb CMMB
}

type DTMB_Res struct {
    Type string
    Dtmb DTMB
}

type TV_Res struct {
    Type string
    Tv TV
}

type Frequency_Res struct {
    Type string
    Frequency_ Frequency
}

type LocationInfo_Res struct {
    Type string
    LocationInfo_ LocationInfo
}

fucn h_webSocket(ws *websocket.Conn){
    for {
        fmt.Println("开始解析数据...")
        var data string
        err := websocket.Message.Receive(ws, &data)
        fmt.Println("data：", data)
        if err != nil {
            fmt.Println("接收出错...")
            continue
        }

        data = strings.Replace(data, "\n", "", 0)
        var webSocketReqInfo WebSocketReqInfo
        err = json.Unmarshal([]byte(data), &webSocketReqInfo)
        if err != nil {
            fmt.Println("解析数据异常...")
            break
        }
        fmt.Println("请求数据类型：", webSocketReqInfo.Type)

        switch webSocketReqInfo.Type {
            case "CMMB":{
                b, errMarshl := json.Marshal(datas)
                if errMarshl != nil {
                    fmt.Println("全局消息内容异常...")
                    break
                }
                errMarshl = websocket.Message.Send(key, string(b))
                if errMarshl != nil {
                    //移除出错的链接
                    fmt.Println("发送出错...")
                    break
                }
            }
            case "DTMB":
               
        }
    }

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
    db.MysqlOpen("testbed_v3_1", "127.0.0.1" ,3306)
    return
}



func main() {
    spectrum_init()
    http.HandleFunc("/", h_index)
    http.HandleFunc("/data", handler)
	http.HandleFunc("/websocket", websocket.Handler(h_webSocket))
    http.ListenAndServe("127.0.0.1:443", nil)
    //http.ListenAndServeTLS(":8081", "server.crt", "server.key", nil)
}
