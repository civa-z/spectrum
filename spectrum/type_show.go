package main

type WebSocketReqInfo struct {
    Type string
}

type Freq_Using_List_Res struct {
    Type string
    Freq_Using_List_ []Freq_Using
}

type Frequency_Res struct {
    Type string
    Frequency_ []Frequency
}

type LocationInfo_Res struct {
    Type string
    LocationInfo_ Location_Info
}

type OnlineDevice_Res struct {
    Type string
    OnlineDevicelist []Online_Device
}
