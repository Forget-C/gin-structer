package base

import "net/http"

type DefaultResp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func (c *DefaultResp) SetSuccess() {
	c.Code = 0
	c.Message = "Success"
}

func (c *DefaultResp) SetBadRequest(message string) {
	c.Code = http.StatusBadRequest
	c.Message = message
}

func (c *DefaultResp) SetNotFound(message string) {
	c.Code = http.StatusNotFound
	c.Message = message
}

func (c *DefaultResp) SetAccepted() {
	c.Code = http.StatusAccepted
	c.Message = "Success"
}

func (c *DefaultResp) SetServerError(message string) {
	c.Code = http.StatusInternalServerError
	c.Message = message
}

func (c *DefaultResp) SetCodeAndMsg(code int, message string) {
	c.Code = code
	c.Message = message
}

func (c *DefaultResp) IsNull() bool {
	if c == nil || c.Result == nil {
		return true
	}
	return false
}

type PaginationMeta struct {
	Page     uint32 `json:"page,omitempty"`
	PageSize uint32 `json:"page_size,omitempty"`
	Total    uint32 `json:"total,omitempty"`
}

type PaginationResult struct {
	Items      interface{}    `json:"items"`
	Pagination PaginationMeta `json:"pagination,omitempty"`
}
