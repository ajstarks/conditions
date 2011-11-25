/*

	conditions - Report current weather conditions

	Prints current conditions (from NOAA via wunderground.com) to the
	console.

	Written and maintained by Stephen Ramsay

	Last Modified: Fri Jul 15 23:59:07 CDT 2011

	Copyright (c) 2011 by Stephen Ramsay

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

const URLstem = "http://api.wunderground.com/auto/wui/geo/WXCurrentObXML/index.xml?query="

const VERS = "1.2.1"

type Weather struct {
	Observation_time   string
	City               string `xml:"display_location>city"`
	State              string `xml:"display_location>state"`
	Station_id         string
	Weather            string
	Temperature_string string
	Relative_humidity  string
	Wind_string        string
	Pressure_string    string
	Dewpoint_string    string
	Heat_index_string  string
	Windchill_string   string
	Visibility_mi      string
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

	URL = URLstem + station_id

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
	fmt.Println("Current conditions at " + weather.City + ", " + weather.State + " (" + weather.Station_id + ")")
	fmt.Println(weather.Observation_time)
	fmt.Println("   Temperature: " + weather.Temperature_string)
	fmt.Println("   Sky Conditions: " + weather.Weather)
	fmt.Println("   Wind: " + weather.Wind_string)
	fmt.Println("   Relative humidity: " + weather.Relative_humidity)
	fmt.Println("   Pressure: " + weather.Pressure_string)
	fmt.Println("   Dewpoint: " + weather.Dewpoint_string)
	if weather.Heat_index_string != "NA" {
		fmt.Println("   Heat Index: " + weather.Heat_index_string)
	}
	if weather.Windchill_string != "NA" {
		fmt.Println("   Windchill: " + weather.Windchill_string)
	}
	fmt.Println("   Visibility: " + weather.Visibility_mi + " miles")
}

func checkError(err os.Error) {
	if err != nil {
		fmt.Println(os.Stderr, "Fatal error ", err.String())
		os.Exit(1)
	}
}
