# sysexits

There is a convention on UNIX systems about what specific application codes mean. Originally from UNIX, these
conventions are [known as "sysexits.h"](https://manpages.ubuntu.com/manpages/lunar/man3/sysexits.h.3head.html).

Go, by convention, [passes errors around to indicate a failure](https://go.dev/blog/go1.13-errors). It uses an 
error type to indicate a class of errors and the error value to indicate a specific error condition within an 
error type.

This package combines these approaches by defining an error type (sysexit) and a series of sysexits conveying 
the termination of an application.

## Usage

See `sysexits_test.go` for up-to-date examples, but broadly:

```go
// Lets imagine we want to express a specific failure mode — For example, a user lookup:
_, err := user.Lookup("fish")

switch err {
case nil:
    err = sysexits.OK
case user.UnknownUserError("fish"):
    err = fmt.Errorf("%w: %s", sysexits.NoUser, err)
default:
    err = fmt.Errorf("%w: %s", sysexits.Software, err)
}

// Later, in the main function, we want to see whether the program has given us a more
// specific exit code:
var exit sysexits.Sysexit
if errors.As(err, &exit) {
    fmt.Println(exit.Code)
    // Output: 67
}
```