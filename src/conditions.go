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
  "regexp"
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
  Full  string
}

func main() {

  key,_ := utils.GetConf()

  var stationId = utils.Options(key)
  var URL string

  URL = utils.BuildURL("conditions", stationId, key)

  res, err := http.Get(URL)
  var b []byte
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

  var current = obs.Current_observation

  fmt.Println("Current conditions at " + current.Observation_location.Full + " (" + current.Station_id + ")")
  fmt.Println(current.Observation_time)
  fmt.Println("   Temperature: " + current.Temperature_string)
  fmt.Println("   Sky Conditions: " + current.Weather)
  fmt.Println("   Wind: " + current.Wind_string)
  fmt.Println("   Relative humidity: " + current.Relative_humidity)
  switch current.Pressure_trend {
  case "+":
      fmt.Println("   Pressure: " + current.Pressure_in + " in (" + current.Pressure_mb + " mb) and rising")
  case "-":
    fmt.Println("   Pressure: " + current.Pressure_in + " in (" + current.Pressure_mb + " mb) and falling")
  case "0":
    fmt.Println("   Pressure: " + current.Pressure_in + " in (" + current.Pressure_mb + " mb) and holding steady")
  }
  fmt.Println("   Dewpoint: " + current.Dewpoint_string)
  if current.Heat_index_string != "NA" {
    fmt.Println("   Heat Index: " + current.Heat_index_string)
  }
  if current.Windchill_string != "NA" {
    fmt.Println("   Windchill: " + current.Windchill_string)
  }
  fmt.Println("   Visibility: " + current.Visibility_mi + " miles")
  m,_ := regexp.MatchString("0.0", current.Precip_today_string)
  if !m {
    fmt.Println("   Precipitation today: " + current.Precip_today_string)
  }
}
