package gtlws

import (
	"regexp"
)

type Routes []Route

type Route struct {
	String string
	Regexp regexp.Regexp
	Handler func([]string) (string, []string)
}
