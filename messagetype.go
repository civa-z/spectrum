package main

import (
    "time"
)

type Init_Req struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Method  string `json:"method"`
	Params  struct {
		Type       string `json:"type"`
		Version    string `json:"version"`
		DeviceDesc struct {
			SerialNumber               string   `json:"serialNumber"`
			ManufacturerID             string   `json:"manufacturerId"`
			ModelID                    string   `json:"modelId"`
			RulesetIds                 []string `json:"rulesetIds"`
			EtsiEnDeviceType           string   `json:"etsiEnDeviceType"`
			EtsiEnDeviceCategory       string   `json:"etsiEnDeviceCategory"`
			EtsiEnDeviceEmissionsClass string   `json:"etsiEnDeviceEmissionsClass"`
			EtsiEnTechnologyID         string   `json:"etsiEnTechnologyId"`
		} `json:"deviceDesc"`
		Location struct {
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
		} `json:"location"`
	} `json:"params"`
}

type Init_Resp struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Type          string      `json:"type"`
		Version       string      `json:"version"`
		ServerMessage interface{} `json:"serverMessage"`
		RulesetInfos  []struct {
			Authority         string `json:"authority"`
			RulesetID         string `json:"rulesetId"`
			MaxLocationChange int    `json:"maxLocationChange"`
			MaxPollingSecs    int    `json:"maxPollingSecs"`
			McwsdSupport      bool   `json:"mcwsdSupport"`
		} `json:"rulesetInfos"`
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
		DeviceDesc struct {
			SerialNumber               string   `json:"serialNumber"`
			ManufacturerID             string   `json:"manufacturerId"`
			ModelID                    string   `json:"modelId"`
			RulesetIds                 []string `json:"rulesetIds"`
			EtsiEnDeviceType           string   `json:"etsiEnDeviceType"`
			EtsiEnDeviceCategory       string   `json:"etsiEnDeviceCategory"`
			EtsiEnDeviceEmissionsClass string   `json:"etsiEnDeviceEmissionsClass"`
			EtsiEnTechnologyID         string   `json:"etsiEnTechnologyId"`
		} `json:"deviceDesc"`
		Location struct {
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
		} `json:"location"`
		Antenna struct {
			Height            int    `json:"height"`
			HeightType        string `json:"heightType"`
			HeightUncertainty int    `json:"heightUncertainty"`
		} `json:"antenna"`
	} `json:"params"`
}

type Avail_Spectrum_Resp struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Type          string      `json:"type"`
		Version       string      `json:"version"`
		ServerMessage interface{} `json:"serverMessage"`
		Timestamp     time.Time   `json:"timestamp"`
		DeviceDesc    struct {
			SerialNumber               string   `json:"serialNumber"`
			ManufacturerID             string   `json:"manufacturerId"`
			ModelID                    string   `json:"modelId"`
			RulesetIds                 []string `json:"rulesetIds"`
			EtsiEnDeviceType           string   `json:"etsiEnDeviceType"`
			EtsiEnDeviceEmissionsClass string   `json:"etsiEnDeviceEmissionsClass"`
			EtsiEnTechnologyID         string   `json:"etsiEnTechnologyId"`
			EtsiEnDeviceCategory       string   `json:"etsiEnDeviceCategory"`
		} `json:"deviceDesc"`
		SpectrumSpecs []struct {
			RulesetInfo struct {
				Authority         string `json:"authority"`
				RulesetID         string `json:"rulesetId"`
				MaxLocationChange int    `json:"maxLocationChange"`
				MaxPollingSecs    int    `json:"maxPollingSecs"`
				McwsdSupport      bool   `json:"mcwsdSupport"`
			} `json:"rulesetInfo"`
			SpectrumSchedules []struct {
				EventTime struct {
					StartTime time.Time `json:"startTime"`
					StopTime  time.Time `json:"stopTime"`
				} `json:"eventTime"`
				Spectra []struct {
					Profiles [][]struct {
						Hz  int `json:"hz"`
						Dbm int `json:"dbm"`
					} `json:"profiles"`
					ResolutionBwHz int `json:"resolutionBwHz"`
				} `json:"spectra"`
			} `json:"spectrumSchedules"`
			TimeRange struct {
				StartTime time.Time `json:"startTime"`
				StopTime  time.Time `json:"stopTime"`
			} `json:"timeRange"`
			FrequencyRanges []struct {
				StartHz int `json:"startHz"`
				StopHz  int `json:"stopHz"`
			} `json:"frequencyRanges"`
			NeedsSpectrumReport            bool   `json:"needsSpectrumReport"`
			MaxTotalBwHz                   int    `json:"maxTotalBwHz"`
			MaxContiguousBwHz              int    `json:"maxContiguousBwHz"`
			EtsiEnSimultaneousChannelOpera string `json:"etsiEnSimultaneousChannelOpera"`
		} `json:"spectrumSpecs"`
	} `json:"result"`
}

type Spectrum_Use_Notify struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Method  string `json:"method"`
	Params  struct {
		Type       string `json:"type"`
		Version    string `json:"version"`
		DeviceDesc struct {
			SerialNumber               string   `json:"serialNumber"`
			ManufacturerID             string   `json:"manufacturerId"`
			ModelID                    string   `json:"modelId"`
			RulesetIds                 []string `json:"rulesetIds"`
			EtsiEnDeviceType           string   `json:"etsiEnDeviceType"`
			EtsiEnDeviceCategory       string   `json:"etsiEnDeviceCategory"`
			EtsiEnDeviceEmissionsClass string   `json:"etsiEnDeviceEmissionsClass"`
			EtsiEnTechnologyID         string   `json:"etsiEnTechnologyId"`
		} `json:"deviceDesc"`
		Location struct {
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
		} `json:"location"`
		Spectra []struct {
			ResolutionBwHz int `json:"resolutionBwHz"`
			Profiles       [][]struct {
				Hz  int `json:"hz"`
				Dbm int `json:"dbm"`
			} `json:"profiles"`
		} `json:"spectra"`
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

