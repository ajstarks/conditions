/*

	forecast

	Prints three-day forecast (via wunderground.com) to the
	console.

	Written and maintained by Stephen Ramsay

	Last Modified: Fri Jul 15 23:56:53 CDT 2011

	Copyright (c) 2011 by Stephen Ramsay

	forecast is free software; you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation; either version 3, or (at your option) any
	later version.

	forecast is distributed in the hope that it will be useful, but
	WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY
	or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License
	for more details.

	You should have received a copy of the GNU General Public License
	along with forecast; see the file COPYING.  If not see
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

const URLstem = "http://api.wunderground.com/auto/wui/geo/ForecastXML/index.xml?query="

const VERS = "1.2.1"

type Day struct {
	Title   string
	Fcttext string
}

type Forecast struct {
	Date        string `xml:"txt_forecast>date"`
	Forecastday []Day  `xml:"txt_forecast>forecastday"`
}

func main() {

	optarg.Add("s", "station", "Weather station.  May be indicated using city, state, CITY,STATE, country, (US or Canadian) zipcode, 3- or 4-letter airport code, or LAT,LONG", "Lincoln, NE")
	optarg.Add("h", "help", "Print this message", false)
	optarg.Add("V", "version", "Print version number", false)

	var station = "Lincoln, NE"
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
		var forecast Forecast
		xmlErr := xml.Unmarshal(res.Body, &forecast)
		checkError(xmlErr)
		printWeather(&forecast, station)
		res.Body.Close()
	}
}

func printWeather(forecast *Forecast, station string) {
	fmt.Println("Forecast for " + station)
	fmt.Println("Issued at " + forecast.Date)
	for i := 0; i < len(forecast.Forecastday); i++ {
		fmt.Println(forecast.Forecastday[i].Title + ": " + forecast.Forecastday[i].Fcttext)
	}
}

func checkError(err os.Error) {
	if err != nil {
		fmt.Println(os.Stderr, "Fatal error ", err.String())
		os.Exit(1)
	}
}
