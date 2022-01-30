package assert

import (
	"errors"
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

// Equal fails the test if got results don't equal expected results.
func Equal(t *testing.T, want, got interface{}, msg string) {
	if reflect.DeepEqual(want, got) {
		return
	}

	_, file, line, _ := runtime.Caller(1)

	fmt.Printf(
		"[%s:%d]: %s. Expected: %#v, real: %#v\n",
		filepath.Base(file), line, msg, want, got,
	)
	t.FailNow()
}

// NotEqual fails the test if got results equal expected results.
func NotEqual(t *testing.T, want, got interface{}, msg string) {
	if !reflect.DeepEqual(want, got) {
		return
	}

	_, file, line, _ := runtime.Caller(1)

	fmt.Printf(
		"[%s:%d]: %s. Expected and real result is the same: %#v\n",
		filepath.Base(file), line, msg, want,
	)
	t.FailNow()
}

// True fails the test if the condition is false.
func True(t *testing.T, ok bool, msg string) {
	if ok {
		return
	}

	_, file, line, _ := runtime.Caller(1)

	fmt.Printf(
		"[%s:%d]: the condition is failed, %s\n",
		filepath.Base(file), line, msg,
	)
	t.FailNow()
}

// False fails the test if the condition is true.
func False(t *testing.T, ok bool, msg string) {
	if !ok {
		return
	}

	_, file, line, _ := runtime.Caller(1)

	fmt.Printf(
		"[%s:%d]: the condition is passed but it shouldn't been, %s\n",
		filepath.Base(file), line, msg,
	)
	t.FailNow()
}

// NotError fails the test if there is an error (not nil).
func NotError(t *testing.T, err error, msg string) {
	if err == nil {
		return
	}

	_, file, line, _ := runtime.Caller(1)

	fmt.Printf(
		"[%s:%d]: %s, error: %s\n",
		filepath.Base(file), line, msg, err,
	)
	t.FailNow()
}

// Error fails the test if there is no error (nil).
func Error(t *testing.T, err error, msg string) {
	if err != nil {
		return
	}

	_, file, line, _ := runtime.Caller(1)

	fmt.Printf(
		"[%s:%d]: %s, no error\n",
		filepath.Base(file), line, msg,
	)
	t.FailNow()
}

// Nil fails the test if got is not nil.
func Nil(t *testing.T, got interface{}, msg string) {
	if got == nil {
		return
	}

	_, file, line, _ := runtime.Caller(1)

	fmt.Printf(
		"[%s:%d]: %s. Expected nil equals %#v\n",
		filepath.Base(file), line, msg, got,
	)
	t.FailNow()
}

// NotNil fails the test if got is nil.
func NotNil(t *testing.T, got interface{}, msg string) {
	if got != nil {
		return
	}

	_, file, line, _ := runtime.Caller(1)

	fmt.Printf(
		"[%s:%d]: %s. Nil is unexpected\n",
		filepath.Base(file), line, msg,
	)
	t.FailNow()
}

// ErrorIs fails the test if errors is not equals.
func ErrorIs(t *testing.T, want, got error, msg string) {
	if errors.Is(want, got) {
		return
	}

	_, file, line, _ := runtime.Caller(1)

	fmt.Printf(
		"[%s:%d]: %s. Expected: %#v, real: %#v\n",
		filepath.Base(file), line, msg, want, got,
	)
	t.FailNow()
}
