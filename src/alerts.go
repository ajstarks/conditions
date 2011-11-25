/*

	alerts - Report three-day forecast

	Prints active alerts (from NOAA via wunderground.com) to the console.

	Written and maintained by Stephen Ramsay

	Last Modified: Fri Jul 15 23:55:17 CDT 2011

	Copyright © 2011 by Stephen Ramsay

	alerts is free software; you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation; either version 3, or (at your option) any
	later version.
	
	alerts is distributed in the hope that it will be useful, but
	WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY
	or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License
	for more details.
	
	You should have received a copy of the GNU General Public License
	along with alerts; see the file COPYING.  If not see
	<http://www.gnu.org/licenses/>

*/
package main

import (
	"fmt"
	"http"
	"os"
	"strconv"
	"strings"
	"xml"
	"github.com/jteeuwen/go-pkg-optarg"
)

const URLstem = "http://api.wunderground.com/auto/wui/geo/AlertsXML/index.xml?query="

const VERS = "1.2.1"

type Alerts struct {
	Alert Alert
}

type Alert struct {
	Count     string      `xml:"attr"`
	Alertitem []Alertitem `xml:"alertitem"`
}

type Alertitem struct {
	Description string
	Message     string
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
		var alerts Alerts
		xmlErr := xml.Unmarshal(res.Body, &alerts)
		checkError(xmlErr)
		printWeather(&alerts)
		res.Body.Close()
	}
}

func printWeather(alerts *Alerts) {
	alertCount, err := strconv.Atoi(alerts.Alert.Count)
	checkError(err)
	if alertCount == 0 {
		fmt.Println("No active alerts")
	} else {
		for i := 0; i < alertCount; i++ {
			fmt.Println(alerts.Alert.Alertitem[i].Description)
			fmt.Println(alerts.Alert.Alertitem[i].Message)
		}
	}
}

func checkError(err os.Error) {
	if err != nil {
		fmt.Println(os.Stderr, "Fatal error ", err.String())
		os.Exit(1)
	}
}
