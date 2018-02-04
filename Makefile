# Version
VERSION = `date +%y.%m`

# If unable to grab the version, default to N/A
ifndef VERSION
    VERSION = "n/a"
endif

#
# Makefile options
#


# State the "phony" targets
.PHONY: all clean build install uninstall


all: build

build: clean
	@echo 'Building qmount...'
	@go build -ldflags '-s -w -X main.Version='${VERSION}

clean:
	@echo 'Cleaning...'
	@go clean

install: build
	@echo installing executable file to /usr/bin/qmount
	@sudo cp qmount /usr/bin/qmount

uninstall: clean
	@echo removing executable file from /usr/bin/qmount
	@sudo rm /usr/bin/qmount