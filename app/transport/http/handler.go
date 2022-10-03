package http

// UserHandler is User HTTP handler
// which consist of embedded UserUsecase interface
type UserHandler struct {
	usecase UserUsecase
	log     Logger
}

// NewUserHandler will return a new instance
// of UserHandler struct accepting UserUsecase interface
func NewUserHandler(usecase UserUsecase, logger Logger) UserHandler {
	return UserHandler{
		usecase: usecase,
		log:     logger,
	}
}
