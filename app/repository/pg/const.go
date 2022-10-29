package pg

const (
	ErrCodeUniqueViolation           = "unique_violation"
	ErrCodeInvalidTextRepresentation = "invalid_text_representation"
)

const (
	defaultParamSort   = "first_name,last_name"
	defaultParamLimit  = "NULL"
	defaultParamOffset = "0"
)
