//
// Package
//
package main

//
// Imports
//
import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
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

	// default mounting location
	defaultMount = "/run/media/"

	// default user
	user = "tmp"

	// whether or not to print debug output
	debug = false
)

// Initialize the argument input flags.
func init() {

	// Device flag
	flag.StringVar(&deviceArg, "device", "",
		"The device to be mounted; e.g. /dev/sdb3")

	// User flag
	flag.StringVar(&user, "user", "tmp",
		"The user who has permission to the drive; e.g. root")

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

	// obtain the last portion of the device
	pieces := strings.Split(deviceToMount, "/")
	if len(pieces) == 0 {
		fmt.Println("Error: the following is an invalid device...",
			deviceToMount)
		return
	}
	finalPiece := pieces[len(pieces)-1]
	if finalPiece == "" {
		fmt.Println("Error: the following is an invalid device...",
			deviceToMount)
		return
	}

	// if no user is specified, revert to the default
	if user == "" {
		user = "tmp"
	}

	// obtain a timestamp, this is used for drive directory name
	datetime := time.Now().UnixNano()

	// check if the device in question is a proper directory; if that
	// didn't work then exit
	file, err := os.Open(deviceToMount)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = file.Stat()
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	file.Close()

	// execute `lsblk`
	bytes, err := lsblk()
	if err != nil {
		fmt.Println(err)
		return
	}
	lsblkOutput := bytes.String()

	// if that worked, string.Split() it via "\n"
	lines := strings.Split(lsblkOutput, "\n")

	// for each line...
	device := ""
	majmin := ""
	rm := ""
	size := ""
	ro := ""
	Type := ""
	mountPoint := ""
	deviceWasNotFound := true
	for _, line := range lines {

		// combine multiple whitespace into a single one
		columns := regexp.MustCompile("\\s+").Split(line, -1)
		if len(columns) < 6 {
			continue
		}

		rawDevice := columns[0]
		majmin = columns[1]
		rm = columns[2]
		size = columns[3]
		ro = columns[4]
		Type = columns[5]
		mountPoint = columns[6]

		// ensure that the device is a partition; e.g. /dev/sdb3
		// if not then exit
		if Type != "part" {
			continue
		}

		// attempt to obtain the size of the device; e.g. 3.4Tb
		if size == "" {
			continue
		}

		reASCII := regexp.MustCompile("([a-zA-Z0-9]+)")
		matches := reASCII.FindAllString(rawDevice, -1)

		if len(matches) == 0 {
			continue
		}

		device = matches[0]

		// check if the device in question is present ...
		if device == finalPiece {
			deviceWasNotFound = false
			break
		}
	}

	// ... if it was not found, exit and print a helpful end message
	if deviceWasNotFound {
		fmt.Println("Error: the following is device is not found...",
			deviceToMount)
		return
	}

	// print the details of the selected device
	if debug {
		fmt.Println(deviceToMount, device, majmin, rm, size,
			ro, Type, mountPoint)
	}

	// check if the device is already mounted; if it is, then exit and
	// print the current directory name
	if mountPoint != "" {
		fmt.Println("Note: the device <", deviceToMount, "> is "+
			"already mounted in the following location:\n", mountPoint)
		return
	}

	// convert T/G/M/K to tb/gb/mb/kb
	sizeString := strings.Replace(size, "T", "tb", -1)
	sizeString = strings.Replace(sizeString, "G", "gb", -1)
	sizeString = strings.Replace(sizeString, "M", "mb", -1)
	sizeString = strings.Replace(sizeString, "K", "kb", -1)

	// swap the . with a _ for the size name
	sizeString = strings.Replace(sizeString, ".", "_", -1)

	// calculate the base32 label from the datetime value
	label := strconv.FormatInt(datetime, 32)

	// since the device is a unmounted partition, then attempt to make
	// a directory that combines its size and a timestamp
	directoryName := sizeString + defaultName + label

	// create the directory path
	path := filepath.Join(defaultMount, user, directoryName)

	// if the directory creation failed, print a helpful message
	err = os.MkdirAll(path, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}

	// otherwise attempt to run the POSIX mount command
	_, err = mount(deviceToMount, path)

	// if the mount command failed, throw an error and exit
	if err != nil {
		fmt.Println(err)
		return
	}

	// otherwise the mount was successful, so print the directory path
	fmt.Println(path)

	// everything worked fine, so return null
	return
}

//! Attempt to execute the lsblk command.
/*
 *  @param    ...string    list of arguments
 *
 *  @return   bytes[]      array of byte buffer data
 */
func lsblk() (bytes.Buffer, error) {

	// variable declaration
	var output bytes.Buffer

	// assemble the command from the list of string arguments
	cmd := exec.Command("lsblk", "-in")
	cmd.Stdout = &output
	cmd.Stderr = &output

	// attempt to execute the command
	err := cmd.Run()

	// if an error occurred, go ahead and pass it back
	if err != nil {
		return output, err
	}

	// having ran the command, pass back the result if no error has
	// occurred
	return output, nil
}

//! Attempt to execute the mount command.
/*
 *  @param    ...string    list of arguments
 *
 *  @return   bytes[]      array of byte buffer data
 */
func mount(args ...string) (bytes.Buffer, error) {

	// variable declaration
	var output bytes.Buffer

	// assemble the command from the list of string arguments
	cmd := exec.Command("mount", args...)
	cmd.Stdout = &output
	cmd.Stderr = &output

	// attempt to execute the command
	err := cmd.Run()

	// if an error occurred, go ahead and pass it back
	if err != nil {
		return output, err
	}

	// having ran the command, pass back the result if no error has
	// occurred
	return output, nil
}
