package rest

const (
	MsgUserCreated        = "User created successfully"
	MsgEmailConflict      = "Email is already taken"
	MsgInternalSeverErr   = "Internal server error"
	MsgBadRequest         = "Bad request"
	MsgUserNotFound       = "User not found"
	MsgTimeOut            = "Connection timeout"
	handlerTimeoutSeconds = 30
	readTimeoutSeconds    = 2
	writeTimeoutSeconds   = 5
)
