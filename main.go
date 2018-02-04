//
// Package
//
package main

//
// Imports
//
import (
	"flag"
	"fmt"
	"strings"
	"time"
)

//
// Globals
//
var (
	// string name of device to be mounted
	deviceArg = ""

	// Whether or not to print the current version of the program
	printVersion = false

	// default version value
	Version = "0.0"

	// default drive name
	defaultName = "_drive_"
)

// Initialize the argument input flags.
func init() {

	// Version mode flag
	flag.StringVar(&deviceArg, "device", "",
		"The device to be mounted; e.g. /dev/sdb3")

	// Version mode flag
	flag.BoolVar(&printVersion, "version", false,
		"Print the current version of this program and exit.")
}

//
// PROGRAM MAIN
//
func main() {

	// Parse the flags, if any.
	flag.Parse()

	// if requested, go ahead and print the version; afterwards exit the
	// program, since this is all done
	if printVersion {
		fmt.Println("qmount v" + Version)
		return
	}

	// quit the program if the device argument is blank
	deviceToMount := strings.TrimSpace(deviceArg)
	if deviceToMount == "" {
		flag.Usage()
		return
	}

	// obtain a timestamp, this is used for drive directory name
	datetime := time.Now().Format(time.UnixDate)

	// TODO: complete this pseudo code
	datetime = datetime

	// check if the device in question is a proper directory; if that
	// didn't work then exit

	// execute `lsblk`

	// if that worked, string.Split() it via " "; if it didn't work
	// then exit

	// check if the device in question is present; if it isn't then
	// exit and print a helpful end message

	// ensure that the device is a partition; e.g. /dev/sdb3
	// if not then exit

	// check if the device is already mounted; if it is, then exit and
	// print the current directory name

	// attempt to obtain the size of the device; e.g. 3.4Tb

	// since the device is a unmounted partition, then attempt to make
	// a directory that combines its size and a timestamp

	// if the directory creation failed, print a helpful message

	// otherwise attempt to run the POSIX mount command

	// if the mount command failed, throw an error and exit

	// otherwise the mount was successful, so print the directory path

	// everything worked fine, so return null
	return
}
