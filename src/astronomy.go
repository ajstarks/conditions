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
	"http"
	"json"
	"io/ioutil"
	"strconv"
)


type Conditions struct {
	Moon_phase Moon_phase
	Sunrise		 Sunrise
	Sunset		 Sunset
}

type Moon_phase struct {
	PercentIlluminated	string
	AgeOfMoon						string
	Sunrise							Sunrise
	Sunset							Sunset
}

type Sunrise struct {
	Hour		string
	Minute	string
}

type Sunset struct {
	Hour		string
	Minute	string
}


func main() {

	var stationId = utils.Options()
	var URL string

	URL = utils.BuildURL("astronomy", stationId)

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

	var age,_ = strconv.Atof32(obs.Moon_phase.AgeOfMoon)
	var percent = obs.Moon_phase.PercentIlluminated

	// Calculate traditional description of lunar phase

	var moonDesc string

	switch {
	case age < 1.84566:  moonDesc = "New moon"
	case age < 5.53699:  moonDesc = "Waxing crescent"
	case age < 9.22831:  moonDesc = "First quarter"
	case age < 12.91963: moonDesc = "Waxing gibbous"
	case age < 16.61096: moonDesc = "Full moon"
	case age < 20.30228: moonDesc = "Waning gibbous"
	case age < 23.99361: moonDesc = "Last quarter"
	case age < 27.68493: moonDesc = "Waning crescent"
	}

	fmt.Println("Moon Phase: " + moonDesc + " (" + percent + "% illuminated)")
	fmt.Println("Sunrise: " + obs.Moon_phase.Sunrise.Hour + ":" + obs.Moon_phase.Sunrise.Minute)
	fmt.Println("Sunset: " + obs.Moon_phase.Sunset.Hour + ":" + obs.Moon_phase.Sunset.Minute)

}
