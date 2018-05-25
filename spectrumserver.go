package main

import (
    "encoding/json"
    "net/http"
    "io/ioutil"
    "time"
    "log"
    "strings"
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

    if !strings.Contains(r.URL.String(), "/date") {
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
    resp.Result.DeviceDesc = req.Params.DeviceDesc
    resp.Result.DeviceDesc.SerialNumber = "750000105"
    resp.Result.DeviceDesc.EtsiEnDeviceEmissionsClass = "3"

    for _, resultid := range req.Params.DeviceDesc.RulesetIds{
        var profile Profile
        profile.Hz = 470000000
        profile.Dbm = 20

        var profile_list []Profile
        profile_list = append(profile_list, profile)

        var spectrum Spectrum
        spectrum.Profiles = append(spectrum.Profiles, profile_list)
	spectrum.ResolutionBwHz = 8000000

        var spectrumSchedule Spectrum_Schedule
        spectrumSchedule.Spectra = append(spectrumSchedule.Spectra, spectrum)
        spectrumSchedule.EventTime.StartTime = time.Now()
        spectrumSchedule.EventTime.StopTime = time.Now()

        var spectrumSpec Spectrum_Spec
        spectrumSpec.SpectrumSchedules = append(spectrumSpec.SpectrumSchedules, spectrumSchedule)

	spectrumSpec.TimeRange.StartTime = time.Now()
	spectrumSpec.TimeRange.StopTime = time.Now()

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


