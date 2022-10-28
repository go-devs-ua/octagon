package rest

const (
	MsgUserCreated        = "User created successfully"
	MsgEmailConflict      = "Email is already taken"
	MsgInternalSeverErr   = "Internal server error"
	MsgBadRequest         = "Bad request"
	MsgNotFound           = "Not found"
	MsgTimeOut            = "Connection timeout"
	handlerTimeoutSeconds = 30
	readTimeoutSeconds    = 2
	writeTimeoutSeconds   = 5
)

const (
	offset    = "offset"
	limit     = "limit"
	sort      = "sort"
	firstName = "first_name"
	lastName  = "last_name"
	createdAt = "created_at"
)
