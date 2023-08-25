package sysexits_test

import (
	"errors"
	"fmt"
	"os/user"
	"testing"

	"github.com/andrewhowdencom/sysexits"
)

func ExampleSysexit_Error() {
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
}

func TestIsError(t *testing.T) {
	t.Parallel()

	var s interface{} = sysexits.Sysexit{
		Code:    99,
		Message: "example failure",
	}

	if _, ok := s.(error); !ok {
		t.Errorf("declared sysexit is not an error type")
	}
}

func TestAs(t *testing.T) {
	t.Parallel()

	var err error = sysexits.DataErr
	var exit sysexits.Sysexit

	if !errors.As(err, &exit) {
		t.Errorf("sysexit cannot be interpreted as error")
	}
}

func TestIs(t *testing.T) {
	t.Parallel()

	if errors.Is(sysexits.IOErr, sysexits.DataErr) {
		t.Error("expected IOErr and DataErr to be different")
	}

	if !errors.Is(sysexits.IOErr, sysexits.Sysexit{
		Code:    sysexits.IOErr.Code,
		Message: "random",
	}) {
		t.Error("expected exit code equality to mean error equality")
	}
}
