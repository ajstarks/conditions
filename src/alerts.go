/*

  alerts - Report three-day forecast

  Prints active alerts (from NOAA via wunderground.com) to the console.

  Written and maintained by Stephen Ramsay

  Last Modified: Fri Nov 25 16:39:46 CST 2011

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
  "./utils"
  "fmt"
  "http"
  "json"
  "io/ioutil"
)

type Conditions struct {
  Alerts []Alerts
}

type Alerts struct {
  Date        string
  Expires     string
  Description string
  Message     string
}

func main() {

  var stationId = utils.Options()
  var URL string

  URL = utils.BuildURL("alerts", stationId)

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

  var alertsNum = len(obs.Alerts)
  var alerts    = obs.Alerts

  if alertsNum == 0 {
    fmt.Println("No active alerts")
  } else {
    for i := 0; i < alertsNum; i++ {
      fmt.Println("### " +  alerts[i].Description + " ###")
      fmt.Println()
      fmt.Println("Issued at " + alerts[i].Date)
      fmt.Println("Expires at " + alerts[i].Expires)
      fmt.Println(alerts[i].Message)
    }
  }
}
