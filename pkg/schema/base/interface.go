package base

type Response interface {
	SetSuccess()
	SetBadRequest(message string)
	SetNotFound(message string)
	SetAccepted()
	SetServerError(message string)
	SetCodeAndMsg(code int, message string)
	IsNull() bool
}

type ResponseImmediately struct {
}

func (r ResponseImmediately) SetSuccess() {
}

func (r ResponseImmediately) SetBadRequest(message string) {
}

func (r ResponseImmediately) SetNotFound(message string) {
}

func (r ResponseImmediately) SetAccepted() {
}

func (r ResponseImmediately) SetServerError(message string) {
}

func (r ResponseImmediately) SetCodeAndMsg(code int, message string) {
}

func (r ResponseImmediately) IsNull() bool {
	return false
}

type SortReqInterface interface {
	SortFilter() interface{}
}

type SearchReqInterface interface {
	SearchFilter(fuzzy bool) interface{}
}

type WordReqInterface interface {
	WordFilter(fuzzy bool, fields ...string) interface{}
}
