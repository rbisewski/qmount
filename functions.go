package main

import (
	"bytes"
	"os/exec"
)

//! Attempt to execute the lsblk command.
/*
 *  @param    ...string    list of arguments
 *
 *  @return   bytes[]      stdout data, as bytes
 *  @return   bytes[]      sterr data, as bytes
 *  @return   error        message of error, if any
 */
func lsblk() (bytes.Buffer, bytes.Buffer, error) {

	// variable declaration
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	// assemble the command from the list of string arguments
	cmd := exec.Command("lsblk", "-iln")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// attempt to execute the command
	err := cmd.Run()

	// if an error occurred, go ahead and pass it back
	if err != nil {
		return stdout, stderr, err
	}

	// having ran the command, pass back the result if no error has
	// occurred
	return stdout, stderr, nil
}

//! Attempt to execute the mount command.
/*
 *  @param    ...string    list of arguments
 *
 *  @return   bytes[]      array of byte buffer data
 */
func mount(device, path string) (bytes.Buffer, error) {

	// variable declaration
	var output bytes.Buffer

	// assemble the command from the list of string arguments
	cmd := exec.Command("mount", device, path)
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
