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

	// TODO: complete this pseudo code
	fmt.Println(lsblkOutput, datetime)

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
	cmd := exec.Command("lsblk")
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
