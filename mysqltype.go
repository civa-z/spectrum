package main

type All_Ptx struct {
    id	        int16
    IP_address  string
    location_x  float64
    location_y  float64
    channel     int16
    PTX         float64
    OFCOM       float64
    ECC         float64
    QoS         float64
}

type CMMB struct {
    ID           int
    DistrictCode string
    Channel      int
    Power        float32
}

type DTMB struct {
    ID           int
    DistrictCode string
    Channel      int
    Power        float32
}

type TV struct {
    ID           int
    DistrictCode string
    Channel      int
    Power        float32
}

type Frequency struct {
    ChannelID    int
    Channel      string
    Video        float32
    Audio        float32
    Center       float32
    Low          float32
    High         float32
}

type LocationInfo {
    Id           int
    Province     string
    City         string
    District     string
    Code         string
}
