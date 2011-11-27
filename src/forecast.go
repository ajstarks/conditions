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
	"./utils"
	"fmt"
	"http"
	"json"
	"io/ioutil"
)

type Conditions struct {
	Forecast Forecast
}

type Forecast struct {
	Txt_forecast Txt_forecast
}

type Txt_forecast struct {
	Date	string
	Forecastday []Forecastday
}

type Forecastday struct {
	Title		string
	Fcttext	string
}

func main() {

	var stationId = utils.Options()
	var URL string

	URL = utils.BuildURL("forecast", stationId)

	res, err := http.Get(URL)

	var b []byte
	var obs Conditions

	if err == nil {
		b, err = ioutil.ReadAll(res.Body)
		res.Body.Close()
		jsonErr := json.Unmarshal(b, &obs)
		utils.CheckError(jsonErr)
		printWeather(&obs, stationId)
	}
}

func printWeather(obs *Conditions, stationId string) {

	var forecastNum = len(obs.Forecast.Txt_forecast.Forecastday)
	var forecasts		= obs.Forecast.Txt_forecast.Forecastday
	var date        = obs.Forecast.Txt_forecast.Date

	fmt.Println("Forecast for " + stationId)
	fmt.Println("Issued at " + date)
	for i := 0; i < forecastNum; i++ {
		fmt.Println(forecasts[i].Title + ": " + forecasts[i].Fcttext)
	}
}
