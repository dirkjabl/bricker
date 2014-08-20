bricker
=======

**bricker is an Go API for the Tinkerforge Hardware.**

*This is not a offical API.*

It implements connections over TCP/IP to a brick daemon (brickd).
More information about Tinkerforge Hardeware: http://www.tinkerforge.com/en/doc/index.html#/software-tcpip-open

**Important Hint**

This is actual a prealpha Version.
Please do not use this in production.
See the information in the version section.

# About

The Tinkerforge Hardware is a easy to use and has an open hardware and software approach.
You can get support for many programming languages, but not a API for Go (07/17/2014).

I want to program my controlling software for my hardware projects in Go, so i start to
implement a Go API for it.

## Supported Hardware

For the actual supportet bricks and bricklets please refer the HARDWARE.md file.

## Examples

You could find examples in the [examples repoitory](github.com/dirkjabl/examples/tree/master/bricker).

The typical structure for an application with the bricker starts with the
creation of a bricker object

    brick := bricker.New()
    defer brick.Done() 

Now you should add one or more connectors.
This connectors are the connections to a real hardware stack.
It could be a USB connection (with brickd), a WLAN or Ethernet master extension.
All this connectors work over TCP/IP.

    // USB connection, localhost, default port (needs a running brickd!)
    conn, err := buffered.NewUnbuffered("localhost:4223")
    if err != nil {
      fmt.Printf("No connection: %s\n", err.Error())
      return
    }
    defer conn.Done()

Attach the connection to the bricker with a name.

    err = brick.Attach(conn, "local")
	if err != nil {
      fmt.Printf("Could not attach connection to bricker: %s\n", err.Error())
      return
    }
    defer brick.Release("local") 

Now you could add subscriber to the bricker.
Depends on with bricklets you have.

## Makefile

The Makefile is only for an easy using, you do not need it.

	make clean     # -> makes in every subdirectory "go clean"

	make build     # -> make in every subdirectory "go build"

    make test      # -> makes in every subdirectory "go test"

    make deeptest  # -> makes in every subdirectory "go test -v"

	make install   # -> makes in every subdirectory "go install"

For comfort a *make all" is also implemented. It calls "build", "test" and "install".

# Version

prealpha -> alpha -> beta -> 0.0.1

The versions starts here with a prealpha version.
This means the implementation has not all features,
is not well tested and it has not all needed unit tests.

After the prealpha versions comes the alpha version.
This version will have all feature, but also it has not all unit tests and is not well tested.

And after the alpha comes the beta version. This version get all unit tests.
When all parts are ok tested, then the production versions started (with 0.0.1).

The first number is the major release number.
In a major release there are no API changes at existing code, only new or additional modifications are allowed.

The secound number is the minor release number.
All new futures ans addional extensions increase this number.

The last numer is the patch level.
Here all patches and bug fixes are counted.

## prealpha
This version inserts the code. All features should be added here.
The hardware and the brick daemon API also get new features (or new hardware),
this features will be added later.