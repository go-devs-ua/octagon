package entities

import (
	"fmt"
	"strconv"
	"strings"
)

// QueryParams represent request query params
// that will be used on transport and repository level.
type QueryParams struct {
	Offset string
	Limit  string
	Sort   string
}

// Validate checks if QueryParameter fields are valid.
func (qp QueryParams) Validate() error {
	limit, err := strconv.Atoi(qp.Limit)
	if err != nil {
		return fmt.Errorf("limit argument has to be a number")
	}

	if limit < 0 {
		return fmt.Errorf("limit argument has to be a positive number")
	}

	offset, err := strconv.Atoi(qp.Offset)
	if err != nil {
		return fmt.Errorf("offset argument has to be a number")
	}

	if offset < 0 {
		return fmt.Errorf("offset argument has to be a positive number")
	}

	allowedSortArgs := []string{"first_name", "last_name", "created_at", ","}

	for _, arg := range allowedSortArgs {
		qp.Sort = strings.ReplaceAll(qp.Sort, arg, "")
	}

	if len(qp.Sort) > 0 {
		return fmt.Errorf("sort argument '%v' does not fit list: %v", qp.Sort, allowedSortArgs)
	}

	return nil
}
