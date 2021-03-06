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
	@echo 'Building mnt...'
	@go build -o mnt -ldflags '-s -w -X main.Version='${VERSION}

clean:
	@echo 'Cleaning...'
	@rm -f mnt
	@go clean

install: build
	@echo installing executable file to /usr/bin/mnt
	@sudo cp mnt /usr/bin/mnt

uninstall: clean
	@echo removing executable file from /usr/bin/mnt
	@sudo rm -f /usr/bin/mnt
