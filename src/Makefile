
PROGS = opt-arg conditions forecast alerts slookup astronomy almanac
PREFIX = /usr/local/bin
GOARCH := $(GOARCH)
GOPATH := $(PWD)

ifeq ($(GOARCH), amd64)
	C   = 6g
	L   = 6l
	SFX = 6
endif

ifeq ($(GOARCH), 386)
	C   = 8g
	L   = 8l
	SFX = 8
endif

all: $(PROGS)

conditions: conditions.o
	$(L) -o conditions conditions.$(SFX)

conditions.o: utils.o
	$(C) conditions.go

forecast: forecast.o
	$(L) -o forecast forecast.$(SFX)

forecast.o: utils.o
	$(C) forecast.go

alerts: alerts.o
	$(L) -o alerts alerts.$(SFX)

alerts.o: utils.o
	$(C) alerts.go

slookup: slookup.o
	$(L) -o slookup slookup.$(SFX)

slookup.o: utils.o
	$(C) slookup.go

astronomy: astronomy.o
	$(L) -o astronomy astronomy.$(SFX)

astronomy.o: utils.o
	$(C) astronomy.go

almanac: almanac.o
	$(L) -o almanac almanac.$(SFX)

almanac.o: utils.o
	$(C) almanac.go

utils.o:
	$(C) utils.go 

clean:
	rm -f $(PROGS)
	rm -f *.8
	rm -f *.6

install:
	cp conditions $(PREFIX)
	cp forecast $(PREFIX)
	cp alerts $(PREFIX)
	cp slookup $(PREFIX)
	cp astronomy $(PREFIX)
	cp almanac $(PREFIX)

uninstall:
	rm $(PREFIX)/conditions
	rm $(PREFIX)/forecast
	rm $(PREFIX)/alerts
	rm $(PREFIX)/slookup
	rm $(PREFIX)/astronomy
	rm $(PREFIX)/almanac
