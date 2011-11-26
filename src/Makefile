
PROGS = conditions forecast alerts slookup

GOARCH := $(GOARCH)

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

slookup.o:
	$(C) slookup.go

utils.o:
	$(C) utils.go 

clean:
	rm -f $(PROGS)
	rm -f *.8
	rm -f *.6

install:
	cp conditions /usr/local/bin/
	cp forecast /usr/local/bin/
	cp alerts /usr/local/bin/
	cp slookup /usr/local/bin/

