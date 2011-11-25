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
	"xml"
	"github.com/jteeuwen/go-pkg-optarg"
)

const URLstem = "http://api.wunderground.com/api/" + api_key + "/conditions/q/"

const VERS = "1.3.0"

type Weather struct {
	Observation_time    string `xml:"current_observation>observation_time"`
	Full                string `xml:"current_observation>display_location>full"`
	Station_id          string `xml:"current_observation>station_id"`
	Weather             string `xml:"current_observation>weather"`
	Temperature_string  string `xml:"current_observation>temperature_string"`
	Relative_humidity   string `xml:"current_observation>relative_humidity"`
	Wind_string         string `xml:"current_observation>wind_string"`
	Pressure_mb         string `xml:"current_observation>pressure_mb"`
	Pressure_in         string `xml:"current_observation>pressure_in"`
	Pressure_trend      string `xml:"current_observation>pressure_trend"`
	Dewpoint_string     string `xml:"current_observation>dewpoint_string"`
	Heat_index_string   string `xml:"current_observation>heat_index_string"`
	Windchill_string    string `xml:"current_observation>windchill_string"`
	Visibility_mi       string `xml:"current_observation>visibility_mi"`
	Precip_today_string string `xml:"current_observation>precip_today_string"`
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

	URL = URLstem + station_id + ".xml"

	res, err := http.Get(URL)

	if err == nil {
		var weather Weather
		xmlErr := xml.Unmarshal(res.Body, &weather)
		res.Body.Close()
		checkError(xmlErr)
		printWeather(&weather)
	}
}

func printWeather(weather *Weather) {
	fmt.Println("Current conditions at " + weather.Full + " (" + weather.Station_id + ")")
	fmt.Println(weather.Observation_time)
	fmt.Println("   Temperature: " + weather.Temperature_string)
	fmt.Println("   Sky Conditions: " + weather.Weather)
	fmt.Println("   Wind: " + weather.Wind_string)
	fmt.Println("   Relative humidity: " + weather.Relative_humidity)
	switch weather.Pressure_trend {
	case "+":
		fmt.Println("   Pressure: " + weather.Pressure_in + " in (" + weather.Pressure_mb + " mb) and rising")
	case "-":
		fmt.Println("   Pressure: " + weather.Pressure_in + " in (" + weather.Pressure_mb + " mb) and falling")
	case "0":
		fmt.Println("   Pressure: " + weather.Pressure_in + " in (" + weather.Pressure_mb + " mb) and holding steady")
	}
	fmt.Println("   Dewpoint: " + weather.Dewpoint_string)
	if weather.Heat_index_string != "NA" {
		fmt.Println("   Heat Index: " + weather.Heat_index_string)
	}
	if weather.Windchill_string != "NA" {
		fmt.Println("   Windchill: " + weather.Windchill_string)
	}
	fmt.Println("   Visibility: " + weather.Visibility_mi + " miles")
	fmt.Println("   Precipitation today: " + weather.Precip_today_string)
}

func checkError(err os.Error) {
	if err != nil {
		fmt.Println(os.Stderr, "Fatal error ", err.String())
		os.Exit(1)
	}
}
