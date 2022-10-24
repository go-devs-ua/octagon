package rest

import "regexp"

var (
	queryParamsRegexp = regexp.MustCompile(maskParams)
	queryArgsRegexp   = regexp.MustCompile(maskArgs)
)
