package cmn

type ErrorCode int

const (
	Ok ErrorCode = iota
	UnsupportedFilterByParameter
	IncorrectValueForSelectedFilter
)

type ErrorResponse struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}
