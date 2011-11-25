
Conditions
==========

__[Weather Alert: Conditions is currently broken, since Wunderground has changed their API.  I'll issue an all clear as soon as I get the code in order.]__

Conditions is a suite of small command-line applications that retrieve weather data from [Weather Underground](http://www.wunderground.com).  The programs are designed to be extremely fast and generally conformant with the [UNIX philosophy](http://en.wikipedia.org/wiki/Unix_philosophy).

Description
-----------

Conditions consists of forecast command-line applications.

* _conditions_ reports the current weather conditions.

* _forecast_ gives the current forecast.

* _alerts_ reports any active weather alerts.

* _slookup_ allows you to determine the codes for the various weather stations in a particular area.
	
All three commands understand the following switches:

* -s location (which can be a city, a state, a (U.S. or Canadian zip code), a 3- or 4-letter airport code, or lat,long).  Default location is KLNK (Lincoln, NE).  See note in "Building Conditions" below.
* -h help
* -V version

Building Conditions
-------------------

Conditions is written in [Go](http://golang.org), and thus requires a working Go compiler.  Assuming you have one of those:

	cd src
	make

Note that the default location for conditions is "KLNK" (Lincoln, NE).  You can change this by altering the contents of this line in each file:

	var station = "KLNK"

License(s)
---------

Conditions is written and maintained by Stephen Ramsay.

This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program.  If not, see [http://www.gnu.org/licenses/](http://www.gnu.org/licenses/).

Data courtesy of Weather Underground, Inc. (WUI) is subject to the [Weather Underground API Terms and Conditions of Use](http://www.wunderground.com/weather/api/d/terms.html).  The author of this software is not affiliated with WUI, and the software is neither sponsored nor endorsed by WUI.

Thanks
------

Conditions was heavily inspired -- and indeed, might be considered a clean-room implementation of -- Jeremy Stanley's [weather](http://fungi.yuggoth.org/weather/).  This is a lovely Python script that does more-or-less the same thing as conditions.  I reimplemented the system because Stanley's had stopped working and I wanted a program that was faster.
