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

	allowedSortArgs := map[string]bool{"first_name": true, "last_name": true, "created_at": true}

	for _, arg := range strings.Split(qp.Sort, ",") {
		if _, ok := allowedSortArgs[arg]; !ok {
			return fmt.Errorf("`%s` does not match tupple of allowed sort arguments: %+v", arg, allowedSortArgs)
		}
	}

	return nil
}
