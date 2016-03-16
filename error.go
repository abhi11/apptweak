package apptweak

import "errors"

var (
	AuthNotPresent     = errors.New("Auth is empty for the request")
	CategoryNotPresent = errors.New("Category is empty for the request")
	TermNotPresent     = errors.New("Term in paramters is empty")
)
