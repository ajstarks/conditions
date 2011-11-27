
PROGS = opt-arg conditions forecast alerts slookup
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

opt-arg:
	GOPATH=$(PWD) goinstall github.com/jteeuwen/go-pkg-optarg

conditions: conditions.o
	$(L) -L $(GOPATH)/pkg/linux_amd64 -o conditions conditions.$(SFX)

conditions.o: utils.o
	$(C) conditions.go

forecast: forecast.o
	$(L) -L $(GOPATH)/pkg/linux_amd64 -o forecast forecast.$(SFX)

forecast.o: utils.o
	$(C) forecast.go

alerts: alerts.o
	$(L) -L $(GOPATH)/pkg/linux_amd64 -o alerts alerts.$(SFX)

alerts.o: utils.o
	$(C) alerts.go

slookup: slookup.o
	$(L) -L $(GOPATH)/pkg/linux_amd64 -o slookup slookup.$(SFX)

slookup.o: utils.o
	$(C) slookup.go

utils.o:
	$(C) -I $(GOPATH)/pkg/linux_amd64 utils.go 

clean:
	rm -f $(PROGS)
	rm -f *.8
	rm -f *.6
	rm -rf pkg
	cd src/github.com/jteeuwen/go-pkg-optarg/; make clean 

install:
	cp conditions /usr/local/bin/
	cp forecast /usr/local/bin/
	cp alerts /usr/local/bin/
	cp slookup /usr/local/bin/

