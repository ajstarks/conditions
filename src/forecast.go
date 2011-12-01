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
	"json"
)

type Conditions struct {
	Forecast Forecast
}

type Forecast struct {
	Txt_forecast Txt_forecast
}

type Txt_forecast struct {
	Date        string
	Forecastday []Forecastday
}

type Forecastday struct {
	Title   string
	Fcttext string
}

func main() {
	var obs Conditions
	key, _ := utils.GetConf()
	stationId := utils.Options()
	url := utils.BuildURL("forecast", stationId, key)
	b, err := utils.Fetch(url)
	utils.CheckError(err)
	jsonErr := json.Unmarshal(b, &obs)
	utils.CheckError(jsonErr)
	printWeather(&obs, stationId)
}

func printWeather(obs *Conditions, stationId string) {
	t := obs.Forecast.Txt_forecast
	fmt.Printf("Forecast for %s\nIssued at %s\n", stationId, t.Date)
	for _, f := range t.Forecastday {
		fmt.Printf("%s: %s\n", f.Title, f.Fcttext)
	}
}
