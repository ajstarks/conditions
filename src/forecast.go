/*

	forecast

	Prints three-day forecast (via wunderground.com) to the
	console.

	Written and maintained by Stephen Ramsay

	Last Modified: Fri Nov 25 22:03:31 CST 2011

	Copyright © 2011 by Stephen Ramsay

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

const URLstem = "http://api.wunderground.com/api/bc5deaeccb858c43/forecast/q/"

const VERS = "1.3.0"

type Forecastday struct {
	Title   string `xml:"forecastday>title"`
	Fcttext string `xml:"forecastday>fcttext"`
}

type Txt_forecast struct {
	Date         string `xml:"forecast>txt_forecast>date"`
	Forecastdays []Forecastday
}

type Forecast struct {
	Txt_forecast Txt_forecast
}

type Response struct {
	Forecast Forecast
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

	fmt.Println(URL)

	res, err := http.Get(URL)

	if err == nil {
		var response Response
		xmlErr := xml.Unmarshal(res.Body, &response)
		checkError(xmlErr)
		printWeather(&response, station)
		res.Body.Close()
	}
}

func printWeather(response *Response, station string) {
	fmt.Println("Forecast for " + station)
	fmt.Println("Issued at " + response.Forecast.Txt_forecast.Date)
	fmt.Println(response)
	fmt.Println(len(response.Forecast.Txt_forecast.Forecastdays))
	for i := 0; i < len(response.Forecast.Txt_forecast.Forecastdays); i++ {
		fmt.Println(response.Forecast.Txt_forecast.Forecastdays[i].Title + ": " + response.Forecast.Txt_forecast.Forecastdays[i].Fcttext)
	}
}

func checkError(err os.Error) {
	if err != nil {
		fmt.Println(os.Stderr, "Fatal error ", err.String())
		os.Exit(1)
	}
}
