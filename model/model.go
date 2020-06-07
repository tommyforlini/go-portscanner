package model

import "strconv"

// State result
type State struct {
	Port     string
	Protocol string
	Opened   bool
}

func (s *State) String() string {
	return "{" + s.Port + " " + s.Protocol + " " + strconv.FormatBool(s.Opened) + "}"
}
