package rest

const (
	MsgInternalSeverErr   = "Internal server error"
	MsgBadRequest         = "Bad request"
	MsgNotFound           = "Not found"
	MsgTimeOut            = "Connection timeout"
	handlerTimeoutSeconds = 30
	readTimeoutSeconds    = 2
	writeTimeoutSeconds   = 5
)

const (
	offset      string = "offset"
	limit       string = "limit"
	sort        string = "sort"
	firstName   string = "first_name"
	lastName    string = "last_name"
	createdAt   string = "created_at"
	defaultSort        = firstName + "," + lastName
)
