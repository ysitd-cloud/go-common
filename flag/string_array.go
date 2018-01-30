package flag

import "strings"

type StringFlags []string

func (f *StringFlags) String() string {
	return strings.Join(*f, " ")
}

func (f *StringFlags) Set(value string) error {
	*f = append(*f, value)
	return nil
}
