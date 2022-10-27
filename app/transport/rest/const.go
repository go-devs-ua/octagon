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

const (
	// Allowed and default query params and args
	// in query that are used to fetch all users from repo:
	offset    = "offset"
	limit     = "limit"
	sort      = "sort"
	firstName = "first_name"
	lastName  = "last_name"
	createdAt = "created_at"
)
