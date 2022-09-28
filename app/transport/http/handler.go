package http

// UserHandler is User HTTP handler
type UserHandler struct {
	logic UserLogic
}

// NewApiHandler will return *UserHandler
// accepting UseCases interface
func NewApiHandler(logic UserLogic) *UserHandler {
	return &UserHandler{
		logic: logic,
	}
}
