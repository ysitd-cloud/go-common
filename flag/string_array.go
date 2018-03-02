package flag

import "strings"

// StringFlags is for load multi from flags
type StringFlags []string

// String for display
func (f *StringFlags) String() string {
	return strings.Join(*f, " ")
}

// Set value by append string array
func (f *StringFlags) Set(value string) error {
	*f = append(*f, value)
	return nil
}
