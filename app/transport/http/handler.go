package http

// UserHandler is User HTTP handler
// which consist of embedded UserUsecase interface
type UserHandler struct {
	usecase UserUsecase
}

// NewUserHandler will return a new instance
// of UserHandler struct accepting UserUsecase interface
func NewUserHandler(usecase UserUsecase) UserHandler {
	return UserHandler{
		usecase: usecase,
	}
}
