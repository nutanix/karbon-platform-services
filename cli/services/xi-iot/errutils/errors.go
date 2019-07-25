package errutils

import (
	"fmt"
	"os"
)

// ErrorType is an enum for type of error
type ErrorType int

const (
	IOErr ErrorType = iota
	SerializeErr
	DeserializeErr
	DependencyErr
	InvalidYamlErr
	GenericErr
)

// XiErr encapsulates error and error type.
// The intended use of this type is for any methods that are not directly called by cobra
type XiErr struct {
	msg string
	t   ErrorType
}

// Exit prints the given string and exits with errCode=1
func Exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

// Exitf prints the given formatted string and exits with errCode=1
func Exitf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
	os.Exit(1)
}

// NewDependencyErr creates a Dependency XiErr object
func NewDependencyErr(msg string) *XiErr {
	return &XiErr{msg: msg, t: DependencyErr}
}

// NewIOErr creates an error IO type
func NewIOErr(msg string) *XiErr {
	return &XiErr{msg: msg, t: IOErr}
}

// NewSerializeErr creates a serialization error object
func NewSerializeErr(entity, fromFormat, details string) *XiErr {
	return &XiErr{msg: fmt.Sprintf("Failed to serialize %s %s: %s", entity, fromFormat, details), t: SerializeErr}
}

// NewGenericErr creates  a generic error object
func NewGenericErr(msg string) *XiErr {
	return &XiErr{msg: msg, t: GenericErr}
}

// NewInvalidYamlErr creates an invalid yaml error
func NewInvalidYamlErr(msg string) *XiErr {
	return &XiErr{msg: fmt.Sprintf("Invalid yaml: %q", msg), t: InvalidYamlErr}
}

// NewCreateErr creates an error object of type "Create"
func NewCreateErr(entityType, entityName, msg string) *XiErr {
	return &XiErr{msg: fmt.Sprintf("Failed to create %s %s: %s", entityType, entityName, msg), t: GenericErr}
}

// CheckErr checks for error and handles it accordingly
func CheckErr(err *XiErr) {
	if err == nil {
		return
	}
	Exit(err.msg)
}
