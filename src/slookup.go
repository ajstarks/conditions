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
	"json"
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
	City string
	Icao string
}

func main() {
	var obs Conditions
	key, _ := utils.GetConf()
	stationId := utils.Options()
	url := utils.BuildURL("geolookup", stationId, key)
	b, err := utils.Fetch(url)
	utils.CheckError(err)
	jsonErr := json.Unmarshal(b, &obs)
	utils.CheckError(jsonErr)
	printWeather(&obs)
}

func printWeather(obs *Conditions) {
	station := obs.Location.Nearby_weather_stations.Airport.Station
	if len(station) == 0 {
		fmt.Println("No area stations")
	} else {
		for _, s := range station {
			fmt.Printf("%s: %s\n", s.City, s.Icao)
		}
	}
}
