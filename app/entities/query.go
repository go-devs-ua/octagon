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

	allowedSortArgs := Set{"first_name": {}, "last_name": {}, "created_at": {}}

	for _, arg := range strings.Split(qp.Sort, ",") {
		if _, ok := allowedSortArgs[arg]; !ok {
			return fmt.Errorf("`%s` does not match allowed sort arguments: %v", arg, allowedSortArgs)
		}
	}

	return nil
}

type Set map[string]struct{}

func (set Set) String() string {
	var (
		end = len(set) - 1
		str strings.Builder
		i   int
	)

	str.WriteString("{")

	for k, _ := range set {
		str.WriteString(k)
		if i != end {
			str.WriteString(", ")
		}
		i++
	}

	str.WriteString("}")

	return str.String()
}
