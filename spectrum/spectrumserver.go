package main

import (
    "encoding/json"
    "net/http"
    "io/ioutil"
    "time"
    "log"
    "strings"
	"fmt"
    "golang.org/x/net/websocket"
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
    //Url: /data
    var resp[]byte
    var ret int

    if !strings.Contains(r.URL.String(), "/data") {
        return nil, 404
    }

    body, _ := ioutil.ReadAll(r.Body)
    body_str := string(body)
    req_body_byte := []byte(body_str)

    t := PraseReqType(req_body_byte)
    log.Println(t)
    switch t {
    case INIT_REQ:
        resp, ret = OnInitReq(req_body_byte)
	break
    case AVAIL_SPECTRUM_REQ:
        resp, ret = OnAvailSpectrumReq(req_body_byte)
        break
    case SPECTRUM_USE_NOTIFY:
	resp, ret = OnSpectrumUseNotify(req_body_byte)
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

    for _, resultid := range req.Params.DeviceDesc.RulesetIds {
        var ruleset_info Ruleset_Info
        ruleset_info.Authority = "uk"
        ruleset_info.RulesetID = resultid
        ruleset_info.MaxLocationChange = 50
        ruleset_info.MaxPollingSecs = 900
        ruleset_info.McwsdSupport = true
	resp.Result.RulesetInfos = append(resp.Result.RulesetInfos, ruleset_info)
    }

    resp_body_byte, e:= json.Marshal(resp)
    if e != nil{
        return nil, 404
    }
    return resp_body_byte, 0
}

var frequency_list []Frequency
func GetChannelID(start_freq float32, end_freq float32) (int, string) {
	if(len(frequency_list) == 0){
		frequency_list = db.GetFrequency()
	}
	
	for _, frequency := range frequency_list{
		if start_freq <= frequency.Center && end_freq >= frequency.Center {
			return frequency.ChannelID, frequency.Channel
		}
	}
	return -1, ""
}

func OnAvailSpectrumReq(req_body_byte []byte) ([]byte, int) {
    var req Avail_Spectrum_Req
    var resp Avail_Spectrum_Resp
	var online_device Online_Device
    json.Unmarshal(req_body_byte, &req)

    //TO DO
    resp.Jsonrpc = req.Jsonrpc
    resp.ID = req.ID
    resp.Result.Type = "AVAIL_SPECTRUM_RESP"
    resp.Result.Version = req.Params.Version
    resp.Result.Timestamp = time.Now()
    resp.Result.DeviceDesc = req.Params.DeviceDesc
	
	online_device.SerialNumber = req.Params.DeviceDesc.SerialNumber
	online_device.FreqUsing.DistrictCode = "UnKnown"
	online_device.Latitude = req.Params.Location.Point.Center.Latitude
	online_device.Longtitude = req.Params.Location.Point.Center.Longitude

    for _, resultid := range req.Params.DeviceDesc.RulesetIds{
        time_now := time.Now()
        var profile_start Profile
        profile_start.Hz = 470000000
        profile_start.Dbm = 16

		var profile_end Profile
        profile_end.Hz = 574000000
        profile_end.Dbm = 16
		
		online_device.FreqUsing.Channel, _ = GetChannelID(profile_start.Hz, profile_end.Hz)
		online_device.FreqUsing.Power = profile_start.Dbm

        var profile_list []Profile
        profile_list = append(profile_list, profile_start)
        profile_list = append(profile_list, profile_end)

        var spectrum Spectrum
        spectrum.Profiles = append(spectrum.Profiles, profile_list)
		spectrum.ResolutionBwHz = 8000000

        var spectrumSchedule Spectrum_Schedule
        spectrumSchedule.Spectra = append(spectrumSchedule.Spectra, spectrum)
        spectrumSchedule.EventTime.StartTime = time_now
        spectrumSchedule.EventTime.StopTime = time_now.AddDate(1,2,3)

        var spectrumSpec Spectrum_Spec
        spectrumSpec.SpectrumSchedules = append(spectrumSpec.SpectrumSchedules, spectrumSchedule)

		spectrumSpec.TimeRange.StartTime = time_now
		spectrumSpec.TimeRange.StopTime = time_now.AddDate(1,2,3)

        var ruleset_info Ruleset_Info
        ruleset_info.Authority = "uk"
        ruleset_info.RulesetID = resultid
        ruleset_info.MaxLocationChange = 50
        ruleset_info.MaxPollingSecs = 900
        ruleset_info.McwsdSupport = true
        spectrumSpec.RulesetInfo = ruleset_info

		var frequencyRange Frequency_Range
		frequencyRange.StartHz = 470000000
		frequencyRange.StopHz = 790000000
		spectrumSpec.FrequencyRanges = append(spectrumSpec.FrequencyRanges, frequencyRange)

		spectrumSpec.NeedsSpectrumReport = true
		spectrumSpec.MaxTotalBwHz = 24000000
		spectrumSpec.MaxContiguousBwHz = 24000000
		spectrumSpec.EtsiEnSimultaneousChannelOpera = "0"

        resp.Result.SpectrumSpecs = append(resp.Result.SpectrumSpecs, spectrumSpec)
    }

	if db.HasOnlineDevice(online_device.SerialNumber){
		db.UpdateOnlineDevice(online_device)
	} else {
		db.InsertOnlineDevice(online_device)
	}
	if g_ws_conn != nil {
		var online_device_res OnlineDevice_Res
		online_device_res.Type = "OnlineDevice"
		online_device_res.OnlineDevicelist = db.GetOnlineDevice("*")
		b, errMarshl:=json.Marshal(online_device_res)
		if errMarshl != nil {
			fmt.Println("取得数据异常online_device_res...")
		}
		errMarshl=websocket.Message.Send(g_ws_conn, string(b))
        if errMarshl != nil {
            //移除出错的链接
            fmt.Println("发送出错...")
        }
	} else {
		fmt.Println("g_ws_conn == nil")
	}
	
    resp_body_byte, e:= json.Marshal(resp)
    if e != nil{
        return nil, 404
    }
    return resp_body_byte, 0
}

func OnSpectrumUseNotify(req_body_byte []byte) ([]byte, int) {
    var req Spectrum_Use_Notify
    var resp Spectrum_Use_Resp
    json.Unmarshal(req_body_byte, &req)

    //TO DO
    resp.Jsonrpc = req.Jsonrpc
    resp.ID = req.ID
    resp.Result.Type = "SPECTRUM_USE_RESP"
    resp.Result.Version = req.Params.Version
    resp.Result.Result = true
    resp.Result.Message = "OK"

    resp_body_byte, e:= json.Marshal(resp)
    if e != nil{
	return nil, 404
    }
    return resp_body_byte, 0
}


