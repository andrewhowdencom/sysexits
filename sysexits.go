package sysexits

import (
	"errors"
)

// Sysexit is a type that implements the error interface, allowing the application to indicate the failure
// mode.
type Sysexit struct {
	// Code is the numeric code that the application is expecte to exit on.
	Code int

	// Message is a represnetation of this underlying failure condition.
	Message string
}

// Variables that describe specific failure modes.
var (
	// OK indicates nothing is wrong.
	OK = Sysexit{Code: 0, Message: ""}

	// Usage indicates command was used incorrectly, e.g., with the wrong number of arguments, a
	// bad flag, bad syntax in a parameter, or whatever.
	Usage = Sysexit{Code: 64, Message: "command line usage error"}

	// DataErr indicates the input data was incorrect in some way. This should only be used for
	// user's data and not system files.
	DataErr = Sysexit{Code: 65, Message: "data format error"}

	// NoInput indicates he input file (not a system file) did not exist or was not  readable. This
	// could also include errors like "No message" to a mailer (if it cared to catch it).
	NoInput = Sysexit{Code: 66, Message: "cannot open input"}

	// NoUser indicates the user specified does not exist. This might be used for mail addresses or
	// remote logins.
	NoUser = Sysexit{Code: 67, Message: "addressee unknown"}

	// NoHost indicates the host specified does not exist. This is used in mail addresses or network
	// requests.
	NoHost = Sysexit{Code: 68, Message: "host name unknown"}

	// Software indicates the service is uanvailable. This can occur if a support program or file
	// does not exist. This can also be used as a catch-all message when something you wanted to do
	// doesn't work, but you don't know why.
	Unavailable = Sysexit{Code: 69, Message: "service unavailable"}

	// An internal software error has been detected. This should be limited to non-operating system
	// related errors if possible.
	Software = Sysexit{Code: 70, Message: "internal software error"}

	// OSErr indicates an operating system has been detected. This is intended to be used for such
	// things like "cannot fork", "cannot create pipe" or the like. It includes things like getuid(2)
	// returning a user that does not exist in the passwd(5) file.
	OSErr = Sysexit{Code: 71, Message: "system error (e.g., can't fork)"}

	// OSFile indicates Some system file (e.g., /etc/passwd, /etc/utmp, etc.) does not exist, cannot
	// be opened or has some sort of error (e.g., syntax error)
	OSFile = Sysexit{Code: 72, Message: "critical OS file missing"}

	// A (user specified) output file cannot be created.
	CantCreat = Sysexit{Code: 73, Message: "can't create (user) output file"}

	// An error occurred while doing I/O on some file.
	IOErr = Sysexit{Code: 74, Message: "input/output error"}

	// TempFail (or Temporary  failure) indicates something that is not really an error. For
	// example that a mailer could not create a connection, and the request should be  reattempted
	// later.
	TempFail = Sysexit{Code: 75, Message: "temp failure; user is invited to retry"}

	// Protocol indicates the  remote  system  returned  something  that was "not possible" during a
	// protocol exchange.
	Protocol = Sysexit{Code: 76, Message: "remote error in protocol"}

	// NoPerm indicates that permission was denied.
	NoPerm = Sysexit{Code: 77, Message: "permission denied"}

	// Config indicates something was found in an unconfigured or misconfigured state.
	Config = Sysexit{Code: 78, Message: "configuration error"}
)

// Error allows the sysexits to be treated as a standard error type. This should be used
// with the error wrapping functionality. For example,
//
//	fmt.Errorf("%w: %s", sysexits.OSErr, err)
func (s Sysexit) Error() string {
	return s.Message
}

// Is allows custom sysexits to be compared, where those sysexits have unuual error messages
// or otherwise are unforseen.
func (s Sysexit) Is(err error) bool {
	var in Sysexit
	if !errors.As(err, &in) {
		return false
	}

	return s.Code == in.Code
}
