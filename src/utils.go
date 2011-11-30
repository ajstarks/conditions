/*

  utils - General functions common to all three utils.

  Written and maintained by Stephen Ramsay

  Last Modified: Sun Nov 27 14:04:35 CST 2011

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

package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"json"
	"flag"
)

// GetVersion returns the version of the package
func GetVersion() string {
	const VERS = "2.0.0"
	return VERS
}

// GetConf returns the API key and weather station from
// the configuration file at $HOME/.condrc
func GetConf() (string, string) {

	type config struct {
		Key     string
		Station string
	}

	var b []byte
	var conf config

	var confFile = os.Getenv("HOME") + "/.condrc"
	b, err := ioutil.ReadFile(confFile)

	if err == nil {
		jsonErr := json.Unmarshal(b, &conf)
		CheckError(jsonErr)
	} else {
		fmt.Println("You must create a .condrc file in $HOME.")
		os.Exit(1)
	}
	return conf.Key, conf.Station
}

// Options handles commandline options and returns a 
// possibly updated weather station string
func Options() string {

	var help, version bool
	var station, sconf string

	_, sconf = GetConf()
	if sconf == "" {
		sconf = "KLNK"
	}

	flag.BoolVar(&help, "h", false, "print this message")
	flag.BoolVar(&version, "v", false, "Print version number")
	flag.StringVar(&station, "s", sconf, "Weather station: \"city, state-abbreviation\", (US or Canadian) zipcode, 3- or 4-letter airport code, or LAT,LONG")
	flag.Parse()

	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if version {
		fmt.Println("conditions " + GetVersion())
		fmt.Println("Copyright (C) 2011 by Stephen Ramsay")
		fmt.Println("Data courtesy of Weather Underground, Inc.")
		fmt.Println("is subject to Weather Underground Data Feed")
		fmt.Println("Terms of Service.  The program itself is free")
		fmt.Println("software, and you are welcome to redistribute")
		fmt.Println("it under certain conditions.  See LICENSE for")
		fmt.Println("details.")
		os.Exit(0)
	}

	// Trap for city-state combinations (e.g. "San Francisco, CA") and
	// make them URL-friendly (e.g. "CA/SanFranciso")

	cityStatePattern := regexp.MustCompile("([A-Za-z ]+), ([A-Za-z ]+)")
	cityState := cityStatePattern.FindStringSubmatch(station)

	if cityState != nil {
		station = cityState[2] + "/" + cityState[1]
		station = strings.Replace(station, " ", "_", -1)
	}
	return station
}

// BuildURL returns the URL required by the Weather Underground API
// from the query type, station id, and API key
func BuildURL(infoType string, stationId string, key string) string {

	const URLstem = "http://api.wunderground.com/api/"
	const query = "/q/"
	const format = ".json"

	return URLstem + key + "/" + infoType + query + stationId + format
}

// CheckError exits on error with a message
func CheckError(err os.Error) {
	if err != nil {
		fmt.Println(os.Stderr, "Fatal error ", err.String())
		os.Exit(1)
	}
}
