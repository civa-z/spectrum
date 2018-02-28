//gohttps/1-http/spectrum.go
package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
)

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

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe("127.0.0.1:8080", nil)
    //http.ListenAndServeTLS(":8081", "server.crt", "server.key", nil)
}
