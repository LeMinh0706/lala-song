package res

const (
	StatusOk      = 200
	CodeSuccess   = 201
	ErrBadRequest = 40000
	ErrFailed     = 40001
	ErrNotFound   = 40400
)

var msg = map[int]string{
	StatusOk:      "Ok",
	CodeSuccess:   "Success",
	ErrBadRequest: "Bad request",
	ErrFailed:     "Request Fail",
	ErrNotFound:   "Not found",
}
