/*

	conditions - Report current weather conditions

	Prints current conditions (from NOAA via wunderground.com) to the
	console.

	Written and maintained by Stephen Ramsay

	Last Modified: Fri Nov 25 16:37:06 CST 2011

	Copyright © 2011 by Stephen Ramsay

	conditions is free software; you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation; either version 3, or (at your option) any
	later version.

	conditions is distributed in the hope that it will be useful, but
	WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY
	or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License
	for more details.

	You should have received a copy of the GNU General Public License
	along with conditions; see the file COPYING.  If not see
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
	Current_observation Current
}

type Current struct {
	Observation_time     string
	Observation_location Location
	Station_id           string
	Weather              string
	Temperature_string   string
	Relative_humidity    string
	Wind_string          string
	Pressure_mb          string
	Pressure_in          string
	Pressure_trend       string
	Dewpoint_string      string
	Heat_index_string    string
	Windchill_string     string
	Visibility_mi        string
	Precip_today_string  string
}

type Location struct {
	Full	string
}

func main() {

	var stationId = utils.Options()
	var URL string

	URL = utils.BuildURL("conditions", stationId)

	res, err := http.Get(URL)
	var b []byte;
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

	var observations = obs.Current_observation

	fmt.Println("Current conditions at " + observations.Observation_location.Full + " (" + observations.Station_id + ")")
	fmt.Println(observations.Observation_time)
	fmt.Println("   Temperature: " + observations.Temperature_string)
	fmt.Println("   Sky Conditions: " + observations.Weather)
	fmt.Println("   Wind: " + observations.Wind_string)
	fmt.Println("   Relative humidity: " + observations.Relative_humidity)
	switch observations.Pressure_trend {
	case "+":
		fmt.Println("   Pressure: " + observations.Pressure_in + " in (" + observations.Pressure_mb + " mb) and rising")
	case "-":
		fmt.Println("   Pressure: " + observations.Pressure_in + " in (" + observations.Pressure_mb + " mb) and falling")
	case "0":
		fmt.Println("   Pressure: " + observations.Pressure_in + " in (" + observations.Pressure_mb + " mb) and holding steady")
	}
	fmt.Println("   Dewpoint: " + observations.Dewpoint_string)
	if observations.Heat_index_string != "NA" {
		fmt.Println("   Heat Index: " + observations.Heat_index_string)
	}
	if observations.Windchill_string != "NA" {
		fmt.Println("   Windchill: " + observations.Windchill_string)
	}
	fmt.Println("   Visibility: " + observations.Visibility_mi + " miles")
	fmt.Println("   Precipitation today: " + observations.Precip_today_string)
}
