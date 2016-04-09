package gtlws

import (
	"regexp"
)

type Routes []Route

type Route struct {
	String  string
	Handler func([]string) (string, []string)
}

type RegRoute struct {
	Regexp  *regexp.Regexp
	Handler func([]string) (string, []string)
}
