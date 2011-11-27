/*

  almanac - Report high- and low-temperature averages.

  Written and maintained by Stephen Ramsay

  Last Modified: Sun Nov 27 16:33:18 CST 2011

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
  Almanac Almanac
}

type Almanac struct {
  Temp_high Temp_high
  Temp_low  Temp_low
}

type Temp_high struct {
  Normal  Normal
  Record  Record
  Recordyear  string
}

type Temp_low struct {
  Normal  Normal
  Record  Record
  Recordyear  string
}

type Normal struct {
  F string
  C string
}

type Record struct {
  F string
  C string
}


func main() {

  var stationId = utils.Options()
  var URL string

  URL = utils.BuildURL("almanac", stationId)

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

  normalHighF := obs.Almanac.Temp_high.Normal.F
  normalHighC := obs.Almanac.Temp_high.Normal.C
  normalLowF  := obs.Almanac.Temp_low.Normal.F
  normalLowC  := obs.Almanac.Temp_low.Normal.C

  recordHighF := obs.Almanac.Temp_high.Record.F
  recordHighC := obs.Almanac.Temp_high.Record.C
  recordHYear := obs.Almanac.Temp_high.Recordyear
  recordLowF  := obs.Almanac.Temp_low.Record.F
  recordLowC  := obs.Almanac.Temp_low.Record.C
  recordLYear := obs.Almanac.Temp_low.Recordyear

  fmt.Println("Normal high: " + normalHighF + " F (" + normalHighC + " C)")
  fmt.Println("Record high: " + recordHighF + " F (" + recordHighC + " C) [" + recordHYear + "]")
 fmt.Println("Normal low: " + normalLowF + " F (" + normalLowC + " C)")
  fmt.Println("Record low: " + recordLowF + " F (" + recordLowC + " C) [" + recordLYear + "]")

}
