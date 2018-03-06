package main

import (
    "encoding/json"
    "net/http"
    "io/ioutil"
    "time"
)

const (
    INIT_REQ            = "INIT_REQ"
    INIT_RESP		= "INIT_RESP"

    AVAIL_SPECTRUM_REQ  = "AVAIL_SPECTRUM_REQ"
    AVAIL_SPECTRUM_RESP = "AVAIL_SPECTRUM_RESP"

    SPECTRUM_USE_NOTIFY = "SPECTRUM_USE_NOTIFY"
    SPECTRUM_USE_RESP	= "SPECTRUM_USE_RESP"
)


func ProcessReq(r *http.Request) ([]byte, int) {
    var resp[]byte
    var ret int

    body, _ := ioutil.ReadAll(r.Body)
    body_str := string(body)
    req_body_byte := []byte(body_str)

    t := PraseReqType(req_body_byte)
    switch t {
    case INIT_REQ:
        resp, ret = OnInitReq(req_body_byte)
	break
    default:
        resp, ret = nil, 404
    }

    return resp, ret
}

func PraseReqType(req_body_byte []byte) (string) {
    var req map[string]interface{}
    json.Unmarshal(req_body_byte, &req)

    if Params, ok := req["params"]; ok {
	ParamsMap := Params.(map[string]interface{})
	if Type, ok := ParamsMap["type"]; ok{
	    return Type.(string)
	}
    }
    return ""
}

func OnInitReq(req_body_byte []byte) ([]byte, int) {
    var req Init_Req
    var resp Init_Resp

    json.Unmarshal(req_body_byte, &req)
    resp.Jsonrpc = req.Jsonrpc
    resp.ID = req.ID
    resp.Result.Type = "INIT_RESP"
    resp.Result.Version = req.Params.Version

    //TO DO
    for _, resultid := range req.Params.DeviceDesc.RulesetIds {
        var ruleset_infos Ruleset_Infos
        ruleset_infos.Authority = "uk"
        ruleset_infos.RulesetID = resultid
        ruleset_infos.MaxLocationChange = 0
        ruleset_infos.MaxPollingSecs = 0
        ruleset_infos.McwsdSupport = false
	//ap := db.GetAllPtx(resultid)
    }


    resp_body_byte, e:= json.Marshal(resp)
    if e != nil{
        return nil, 404
    }
    return resp_body_byte, 0
}

func OnAvailSpectrumReq(req_body_byte []byte) ([]byte, int) {
    var req Avail_Spectrum_Req
    var resp Avail_Spectrum_Resp
    json.Unmarshal(req_body_byte, &req)

    //TO DO
    resp.Jsonrpc = req.Jsonrpc
    resp.ID = req.ID
    resp.Result.Type = "AVAIL_SPECTRUM_RESP"
    resp.Result.Version = req.Params.Version
    resp.Result.Timestamp = time.Now()

    resp_body_byte, e:= json.Marshal(resp)
    if e != nil{
        return nil, 404
    }
    return resp_body_byte, 0

}


