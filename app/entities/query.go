package entities

// QueryParams represent request query params
// that will be used on transport and repository level.
type QueryParams struct {
	Offset int
	Limit  int
	Sort   string
}
