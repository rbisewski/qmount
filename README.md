# qmount

Quickly mount a device in Linux, written in golang.

# Requirements

This program should be able to run on any Linux or Unix OS that utilizes
GNU mount and has golang. The minimum requirements for this are as follows:

* golang 1.6+
* GNU mount
* GNU lsblk

# Installation

Enter the following command to build the executable (if necessary as root):

```
make
```

Afterwards run the binary from the commandline, as you would any typical
golang program. Alternatively, consider installing it, like so:

```
make install
```

At this point qmount will become installed on the system. Later, if it is
no longer needed, it can be removed via the following command:

```
make uninstall
```

# Running qmount

Simply run the compiled file from the commandline and specify the device
that you wish to mount. For example:

```
mnt -device /dev/sda6
```

This will mount /dev/sda6 and will then return a directory of where the
device in question was mounted. The current timestamp will be appended to the
directory, along with the size. Specifically, if the device is successfully
mounted, it will display output like so:

```
/run/media/user/1_7tb_drive_bjgl80ssu07s
```

At this point the device is mounted and is ready to be read or written to.

# Additional Notes

If a given device is already mounted, the program will simply return the
current mount path of the device.
