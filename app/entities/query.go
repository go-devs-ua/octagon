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
	if qp.Limit == "" {
		return fmt.Errorf("limit argument has to be set")
	}

	if _, err := strconv.Atoi(qp.Limit); err != nil {
		return fmt.Errorf("limit argument has to be a number")
	}

	if qp.Offset == "" {
		return fmt.Errorf("offset argument has to be set")
	}

	if _, err := strconv.Atoi(qp.Offset); err != nil {
		return fmt.Errorf("offset argument has to be a number")
	}

	allowedSortArgs := []string{"first_name", "last_name", "created_at"}

	const separator = ","

	for _, arg := range append(allowedSortArgs, separator) {
		qp.Sort = strings.ReplaceAll(qp.Sort, arg, "")
	}

	if len(qp.Sort) > 0 {
		return fmt.Errorf("`%s` does not match allowed sort arguments: %v", qp.Sort, strings.Join(allowedSortArgs, separator+" "))
	}

	return nil
}
