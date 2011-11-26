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
	"fmt"
	"http"
	"os"
	"strings"
	"json"
	"io/ioutil"
	"github.com/jteeuwen/go-pkg-optarg"
)

const URLstem = "http://api.wunderground.com/api/bc5deaeccb858c43/conditions/q/"

const VERS = "1.3.0"

type Observation struct {
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

	optarg.Add("s", "station", "Weather station.  May be indicated using city, state, CITY,STATE, country, (US or Canadian) zipcode, 3- or 4-letter airport code, or LAT,LONG", "KLNK")
	optarg.Add("h", "help", "Print this message", false)
	optarg.Add("V", "version", "Print version number", false)

	var station = "KLNK"
	var help, version bool
	var URL string

	for opt := range optarg.Parse() {
		switch opt.ShortName {
		case "s":
			station = opt.String()
		case "h":
			help = opt.Bool()
		case "V":
			version = opt.Bool()
		}
	}

	if help {
		optarg.Usage()
		os.Exit(0)
	}

	if version {
		fmt.Println("conditions " + VERS)
		fmt.Println("Copyright (C) 2011 by Stephen Ramsay")
		fmt.Println("Data courtesy of Weather Underground, Inc.")
		fmt.Println("is subject to Weather Underground Data Feed")
		fmt.Println("Terms of Service.  The program itself is free")
		fmt.Println("software, and you are welcome to redistribute")
		fmt.Println("it under certain conditions.  See LICENSE for")
		fmt.Println("details.")
		os.Exit(0)
	}

	// Temporarily trim whitespace locations with spaces
	// (e.g. "New York, NY" -> "NewYork,NY")
	var station_components = strings.Fields(station)
	var station_id = ""
	for i := 0; i < len(station_components); i++ {
		station_id = station_id + station_components[i]
	}

	URL = URLstem + station_id + ".json"

	res, err := http.Get(URL)
	var b []byte;
	var obs Observation

	if err == nil {
		b, err = ioutil.ReadAll(res.Body);
		res.Body.Close();
		jsonErr := json.Unmarshal(b, &obs)
		res.Body.Close()
		checkError(jsonErr)
		printWeather(&obs)
	}
}

func printWeather(obs *Observation) {
	fmt.Println("Current conditions at " + obs.Current_observation.Observation_location.Full + " (" + obs.Current_observation.Station_id + ")")
	fmt.Println(obs.Current_observation.Observation_time)
	fmt.Println("   Temperature: " + obs.Current_observation.Temperature_string)
	fmt.Println("   Sky Conditions: " + obs.Current_observation.Weather)
	fmt.Println("   Wind: " + obs.Current_observation.Wind_string)
	fmt.Println("   Relative humidity: " + obs.Current_observation.Relative_humidity)
	switch obs.Current_observation.Pressure_trend {
	case "+":
		fmt.Println("   Pressure: " + obs.Current_observation.Pressure_in + " in (" + obs.Current_observation.Pressure_mb + " mb) and rising")
	case "-":
		fmt.Println("   Pressure: " + obs.Current_observation.Pressure_in + " in (" + obs.Current_observation.Pressure_mb + " mb) and falling")
	case "0":
		fmt.Println("   Pressure: " + obs.Current_observation.Pressure_in + " in (" + obs.Current_observation.Pressure_mb + " mb) and holding steady")
	}
	fmt.Println("   Dewpoint: " + obs.Current_observation.Dewpoint_string)
	if obs.Current_observation.Heat_index_string != "NA" {
		fmt.Println("   Heat Index: " + obs.Current_observation.Heat_index_string)
	}
	if obs.Current_observation.Windchill_string != "NA" {
		fmt.Println("   Windchill: " + obs.Current_observation.Windchill_string)
	}
	fmt.Println("   Visibility: " + obs.Current_observation.Visibility_mi + " miles")
	fmt.Println("   Precipitation today: " + obs.Current_observation.Precip_today_string)
}

func checkError(err os.Error) {
	if err != nil {
		fmt.Println(os.Stderr, "Fatal error ", err.String())
		os.Exit(1)
	}
}
