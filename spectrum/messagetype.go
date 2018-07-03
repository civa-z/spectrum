package main

import (
    "time"
)


type Location_ struct {
	Point struct {
		Center struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"center"`
		SemiMajorAxis int `json:"semiMajorAxis"`
		SemiMinorAxis int `json:"semiMinorAxis"`
		Orientation   int `json:"orientation"`
	} `json:"point"`
	Confidence int `json:"confidence"`
}

type Device_Desc struct {
	SerialNumber               string   `json:"serialNumber"`
	ManufacturerID             string   `json:"manufacturerId"`
	ModelID                    string   `json:"modelId"`
	RulesetIds                 []string `json:"rulesetIds"`
	EtsiEnDeviceType           string   `json:"etsiEnDeviceType"`
	EtsiEnDeviceCategory       string   `json:"etsiEnDeviceCategory"`
	EtsiEnDeviceEmissionsClass string   `json:"etsiEnDeviceEmissionsClass"`
	EtsiEnTechnologyID         string   `json:"etsiEnTechnologyId"`
}

type Init_Req struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Method  string `json:"method"`
	Params  struct {
		Type       string `json:"type"`
		Version    string `json:"version"`
		DeviceDesc Device_Desc `json:"deviceDesc"`
		Location Location_ `json:"location"`
	} `json:"params"`
}

type Ruleset_Info  struct {
	Authority         string `json:"authority"`
	RulesetID         string `json:"rulesetId"`
	MaxLocationChange int    `json:"maxLocationChange"`
	MaxPollingSecs    int    `json:"maxPollingSecs"`
	McwsdSupport      bool   `json:"mcwsdSupport"`
}

type Init_Resp struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Type          string      `json:"type"`
		Version       string      `json:"version"`
		ServerMessage interface{} `json:"serverMessage"`
		RulesetInfos  []Ruleset_Info `json:"rulesetInfos"`
		DatabaseChange interface{} `json:"databaseChange"`
	} `json:"result"`
}


type Avail_Spectrum_Req struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Method  string `json:"method"`
	Params  struct {
		Type       string `json:"type"`
		Version    string `json:"version"`
		DeviceDesc Device_Desc `json:"deviceDesc"`
		Location Location_ `json:"location"`
		Antenna struct {
			Height            int    `json:"height"`
			HeightType        string `json:"heightType"`
			HeightUncertainty int    `json:"heightUncertainty"`
		} `json:"antenna"`
        MasterDeviceDesc Device_Desc `json:"masterDeviceDesc"`//Slave req
        MasterDeviceLocation Location_ `json:"masterDeviceLocation"`//Slave req
        RequestType string `json:"requestType"`//GOP desc: request channel list for Slave(CPE) in Master
	} `json:"params"`
}

type Profile struct {
    Hz  float32 `json:"hz"`
    Dbm float32 `json:"dbm"`
}

type Spectrum struct {
    Profiles [][]Profile `json:"profiles"`
    ResolutionBwHz float32 `json:"resolutionBwHz"`
}

type Spectrum_Schedule struct {
    EventTime struct {
        StartTime time.Time `json:"startTime"`
        StopTime  time.Time `json:"stopTime"`
    } `json:"eventTime"`
    Spectra []Spectrum `json:"spectra"`
}

type Frequency_Range struct {
    StartHz int `json:"startHz"`
    StopHz  int `json:"stopHz"`
}

type Spectrum_Spec struct {
	RulesetInfo Ruleset_Info `json:"rulesetInfo"`
	SpectrumSchedules []Spectrum_Schedule `json:"spectrumSchedules"`
	TimeRange struct {
		StartTime time.Time `json:"startTime"`
		StopTime  time.Time `json:"stopTime"`
	} `json:"timeRange"`
	FrequencyRanges []Frequency_Range `json:"frequencyRanges"`
	NeedsSpectrumReport            bool   `json:"needsSpectrumReport"`
	MaxTotalBwHz                   int    `json:"maxTotalBwHz"`
	MaxContiguousBwHz              int    `json:"maxContiguousBwHz"`
	EtsiEnSimultaneousChannelOpera string `json:"etsiEnSimultaneousChannelOpera"`
}

type Avail_Spectrum_Resp struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Type          string      `json:"type"`
		Version       string      `json:"version"`
		DeviceDesc Device_Desc `json:"deviceDesc"`
		ServerMessage interface{} `json:"serverMessage"`
		Timestamp     time.Time   `json:"timestamp"`
		SpectrumSpecs []Spectrum_Spec  `json:"spectrumSpecs"`
	} `json:"result"`
}

type Spectrum_Use_Notify struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Method  string `json:"method"`
	Params  struct {
		Type       string `json:"type"`
		Version    string `json:"version"`
		DeviceDesc Device_Desc `json:"deviceDesc"`
		Location Location_ `json:"location"`
		Spectra []Spectrum `json:"spectra"`
        MasterDeviceDesc Device_Desc `json:"masterDeviceDesc"`//Slave req
        MasterDeviceLocation Location_ `json:"masterDeviceLocation"`//Slave req
	} `json:"params"`
}

type Spectrum_Use_Resp struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Type          string      `json:"type"`
		Version       string      `json:"version"`
		ServerMessage interface{} `json:"serverMessage"`
		Result        bool        `json:"result"`
		Message       string      `json:"message"`
	} `json:"result"`
}

