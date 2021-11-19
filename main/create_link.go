package main

import "fmt"

func createLink(sliceIP []string) []string {
	var sliceString []string
	for _, v := range sliceIP {
		sliceString = append(sliceString, fmt.Sprintf("http://ip-api.com/json/%v?fields=status,message,continent,continentCode,country,countryCode,region,regionName,city,district,zip,lat,lon,timezone,offset,currency,isp,org,as,asname,reverse,mobile,proxy,hosting,query", v))
	}

	return sliceString
}
