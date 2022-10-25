package rest

const (
	MsgUserCreated      = "User created successfully"
	MsgEmailConflict    = "Email is already taken"
	MsgInternalSeverErr = "Internal server error"
	MsgBadRequest       = "Bad request"
	MsgNoContent        = "Nothing to show"
	MsgUserNotFound     = "User not found"
	MsgTimeOut          = "Connection timeout"
	MsgParam            = "Parameter"
	MsgArg              = "Argument"
	MsgErr              = "Err"
)

const (
	handlerTimeoutSeconds = 30
	readTimeoutSeconds    = 2
	writeTimeoutSeconds   = 5
	queryTimeoutSeconds   = 3
)

const (
	// Allowed and default query params and args
	// in query that are used to fetch all users from repo:
	maskParams    = "^limit|offset|sort$"
	maskArgs      = "^(first_name|last_name|created_at)([,]*(first_name|last_name|created_at))*$|^[0-9]+$"
	offset        = "offset"
	limit         = "limit"
	sort          = "sort"
	defaultOffset = 0
	defaultLimit  = 5
	defaultSort   = "first_name,last_name"
)
