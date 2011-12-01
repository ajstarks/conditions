/*

  astronomy - Report current astronomical observations.

  Written and maintained by Stephen Ramsay

  Last Modified: Sun Nov 27 15:33:08 CST 2011

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
	"json"
	"strconv"
)

type Conditions struct {
	Moon_phase Moon_phase
	Sunrise    Sunrise
	Sunset     Sunset
}

type Moon_phase struct {
	PercentIlluminated string
	AgeOfMoon          string
	Sunrise            Sunrise
	Sunset             Sunset
}

type Sunrise struct {
	Hour   string
	Minute string
}

type Sunset struct {
	Hour   string
	Minute string
}

func main() {
	var obs Conditions
	key, _ := utils.GetConf()
	stationId := utils.Options()
	url := utils.BuildURL("astronomy", stationId, key)
	b, err := utils.Fetch(url)
	utils.CheckError(err)
	jsonErr := json.Unmarshal(b, &obs)
	utils.CheckError(jsonErr)
	printWeather(&obs, stationId)
}

func printWeather(obs *Conditions, stationId string) {

	var age, _ = strconv.Atoi(obs.Moon_phase.AgeOfMoon)
	var moonDesc string

	// Calculate traditional description of lunar phase

	switch {
	case age < 2:
		moonDesc = "New moon"
	case age < 6:
		moonDesc = "Waxing crescent"
	case age < 9:
		moonDesc = "First quarter"
	case age < 13:
		moonDesc = "Waxing gibbous"
	case age < 17:
		moonDesc = "Full moon"
	case age < 20:
		moonDesc = "Waning gibbous"
	case age < 24:
		moonDesc = "Last quarter"
	case age < 28:
		moonDesc = "Waning crescent"
	}
	sr := obs.Moon_phase.Sunrise
	ss := obs.Moon_phase.Sunset
	percent := obs.Moon_phase.PercentIlluminated
	fmt.Printf("Moon Phase: %s (%s%% illuminated)\n", moonDesc, percent)
	fmt.Printf("Sunrise   : %s:%s\n", sr.Hour, sr.Minute)
	fmt.Printf("Sunset    : %s:%s\n", ss.Hour, ss.Minute)
}
