/*

	slookup - Lookup weather stations

	Prints area weather stations (from NOAA via wunderground.com)
	to the console.

	Written and maintained by Stephen Ramsay

	Last Modified: Wed Aug 03 14:02:45 CDT 2011

	Copyright © 2011 by Stephen Ramsay

	slookup is free software; you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation; either version 3, or (at your option) any
	later version.
	
	slookup is distributed in the hope that it will be useful, but
	WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY
	or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License
	for more details.
	
	You should have received a copy of the GNU General Public License
	along with slookup; see the file COPYING.  If not see
	<http://www.gnu.org/licenses/>

*/
package main

import (
	"./utils"
	"fmt"
	"http"
	"json"
	"io/ioutil"
)

type Conditions struct {
	Location Location
}

type Location struct {
	Nearby_weather_stations Nearby_weather_stations
}

type Nearby_weather_stations struct {
	Airport Airport
}

type Airport struct {
	Station []Station
}

type Station struct {
	City	string
	Icao	string
}

func main() {

	var stationId = utils.Options()
	var URL	string

	URL = utils.BuildURL("geolookup", stationId)

	res, err := http.Get(URL)

	var b []byte
	var obs Conditions

	if err == nil {
		b, err = ioutil.ReadAll(res.Body)
		res.Body.Close()
		jsonErr := json.Unmarshal(b, &obs)
		utils.CheckError(jsonErr)
		printWeather(&obs)
	}

}

func printWeather(obs *Conditions) {

	var stationNum = len(obs.Location.Nearby_weather_stations.Airport.Station)
	var stations	 = obs.Location.Nearby_weather_stations.Airport

	if stationNum == 0 {
		fmt.Println("No area stations")
	} else {
		for i := 0; i < stationNum; i++ {
			fmt.Println(stations.Station[i].City + ": " + stations.Station[i].Icao)
		}
	}
}
