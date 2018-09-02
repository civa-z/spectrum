package main

import (
    //"strings"

)

func getvalidspectrum(location Location_) ([]Profile) {

    var location_code string = "110105000000"
    var profiles []Profile

    freq_list := db.GetFrequency()
    usingfreq_global := getUsingFreq_global(location_code)
    onlinedevice_local := db.GetOnlineDevice(location_code)

    //TODO using those there list to calculate the new free frequency list
    // that can be used for new request
    _ = freq_list
    _ = usingfreq_global
    _ = onlinedevice_local

    return profiles
}

//Frequency used by CMMB, DTMB and TV
func getUsingFreq_global(location_code string) ([]Freq_Using) {
    var using_freq_list []Freq_Using

    using_freq_list_CMMB := db.GetUsingFrequency("CMMB")
    using_freq_list_DTMB := db.GetUsingFrequency("DTMB")
    using_freq_list_TV := db.GetUsingFrequency("TV")

    for _, value := range using_freq_list_CMMB{
        if value.DistrictCode == location_code{
            using_freq_list = append(using_freq_list, value)
		}
    }


    for _, value := range using_freq_list_DTMB{
        if value.DistrictCode == location_code{
            using_freq_list = append(using_freq_list, value)
		}
    }

    for _, value := range using_freq_list_TV{
        if value.DistrictCode == location_code{
            using_freq_list = append(using_freq_list, value)
		}
    }
    return using_freq_list
}
