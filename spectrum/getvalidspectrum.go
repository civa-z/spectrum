package main

import (
    //"strings"

)

func getvalidspectrum(location Location_) ([]Profile) {

    var location_code string = "110105000000"
    var profiles []Profile

    freq_list := db.GetFrequency()
    usingfreq_global := getUsingFreq_global(location_code)
    usignfreq_local := getUsingFreq_local(location_code)

    //TODO using those there list to calculate the new free frequency list
    // that can be used for new request
    _ = freq_list
    _ = usingfreq_global
    _ = usignfreq_local

    return profiles
}

//Frequency used by CMMB, DTMB and TV
func getUsingFreq_global(location_code string) ([]Freq_Using) {
    var using_freq []Freq_Using

    using_freq_CMMB := db.GetUsingFrequency("CMMB")
    using_freq_DTMB := db.GetUsingFrequency("DTMB")
    using_freq_TV := db.GetUsingFrequency("TV")

    for _, value := range using_freq_CMMB.FreqUsingList{
        if value.DistrictCode == location_code{
            using_freq = append(using_freq, value)
	}
    }


    for _, value := range using_freq_DTMB.FreqUsingList{
        if value.DistrictCode == location_code{
            using_freq = append(using_freq, value)
	}
    }

    for _, value := range using_freq_TV.FreqUsingList{
        if value.DistrictCode == location_code{
            using_freq = append(using_freq, value)
	}
    }
    return using_freq
}

func getUsingFreq_local(location_code string) ([]Freq_Using_Local){
    var freq_using_local_list []Freq_Using_Local

    using_freq_local := db.GetUsingFreqLocal(location_code)
    for _, value := range using_freq_local {
        if value.FreqUsing.DistrictCode == location_code {
	    freq_using_local_list = append(freq_using_local_list, value)
	}
    }
    return freq_using_local_list

}

