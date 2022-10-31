package entities

import (
	"fmt"
	"regexp"
)

// QueryParams represent request query params
// that will be used on transport and repository level.
type QueryParams struct {
	Offset string
	Limit  string
	Sort   string
}

const (
	numericArgsMask = `^[0-9]+$`
	sortArgsMask    = `^(first_name|last_name|created_at)([,]*(first_name|last_name|created_at))*$`
)

var (
	numericArgsRegex = regexp.MustCompile(numericArgsMask)
	sortArgsRegex    = regexp.MustCompile(sortArgsMask)
)

// Validate checks if QueryParameter fields are valid.
func (qp QueryParams) Validate() error {
	if !numericArgsRegex.MatchString(qp.Limit) && qp.Limit != "" {
		return fmt.Errorf("limit argument does not match with regex: `%s`", numericArgsMask)
	}

	if !numericArgsRegex.MatchString(qp.Offset) && qp.Offset != "" {
		return fmt.Errorf("offset argument does not match with regex: `%s`", numericArgsMask)
	}

	if !sortArgsRegex.MatchString(qp.Sort) && qp.Sort != "" {
		return fmt.Errorf("sort arguments does not match with regex: `%s`", sortArgsMask)
	}

	return nil
}
